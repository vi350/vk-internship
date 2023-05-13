package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/localization"
	"github.com/vi350/vk-internship/internal/app/storage/user_storage"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const (
	getMeMethod       = "getMe"
	updatesMethod     = "getUpdates"
	sendMessageMethod = "sendMessage"
	sendPhotoMethod   = "sendPhoto"
)

type Client struct {
	scheme   string
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) (c *Client, err error) {
	defer func() { err = e.WrapIfErr("error getting me: ", err) }()

	c = &Client{
		scheme:   "https",
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}

	values := url.Values{}
	data, err := c.doRequest(getMeMethod, values, nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetMeResponse
	err = json.Unmarshal(data, &res)
	if err != nil || res.Ok != true || res.Result.IsBot != true {
		return nil, err
	}

	return c, nil
}

func newBasePath(token string) string {
	return "bot" + token
}

func addIfReplyMarkup(values *url.Values, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error marshalling reply markup", err) }()

	if replyMarkup != nil {
		var j []byte
		j, err = json.Marshal(replyMarkup)
		if err != nil {
			return err
		}
		values.Add("reply_markup", string(j))
	}

	return nil
}

func (c *Client) doRequest(method string, values url.Values, reqBody *bytes.Buffer, header http.Header) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("error performing request: ", err) }()

	u := url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), reqBody)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	req.Header = header
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) SendTextMessageByUser(userFromRegistry *user_storage.User, mType localization.MessageType) (err error) {
	defer func() { err = e.WrapIfErr("error sending text message: ", err) }()

	err = c.SendTextMessage(userFromRegistry.ID,
		localization.GetLocalizedText(mType, userFromRegistry.Language),
		localization.GetLocalizedInlineKeyboardMarkup(mType, userFromRegistry.Language))

	return
}

func (c *Client) SendTextMessage(ID int64, text string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error sending text message: ", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(ID, 10))
	values.Add("text", text)
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return err
	}

	_, err = c.doRequest(sendMessageMethod, values, nil, nil)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) SendImage(ID int64, text string, replyMarkup ReplyMarkup, image string) (err error) {
	defer func() { err = e.WrapIfErr("error sending image message: ", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(ID, 10))
	values.Add("caption", text)
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return err
	}

	if _, err = os.Stat(image); err == nil {
		var file *os.File
		if file, err = os.Open(image); err == nil {
			return e.WrapIfErr("error opening file: ", err)
		}
		defer func() { _ = file.Close() }()

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		var fileField io.Writer
		fileField, err = writer.CreateFormFile("file", image)
		if err != nil {
			return e.WrapIfErr("Failed to create form field:", err)
		}

		_, err = io.Copy(fileField, file)
		if err != nil {
			return e.WrapIfErr("Failed to copy file data:", err)
		}

		if err = writer.Close(); err != nil {
			return e.WrapIfErr("error closing writer: ", err)
		}

		_, err = c.doRequest(sendMessageMethod, values, &buf,
			map[string][]string{
				"Content-Type": {writer.FormDataContentType()},
			})
		if err != nil {
			return err
		}

	} else if errors.Is(err, fs.ErrNotExist) {
		err = nil
		values.Add("photo", image)
		_, err = c.doRequest(sendPhotoMethod, values, nil, nil)
	} else {
		return err
	}

	return nil
}

func (c *Client) Updates(offset int64, limit int) (updates []Update, err error) {
	defer func() { err = e.WrapIfErr("error getting updates: ", err) }()

	values := url.Values{}
	values.Add("offset", strconv.FormatInt(offset, 10))
	values.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(updatesMethod, values, nil, nil)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}
