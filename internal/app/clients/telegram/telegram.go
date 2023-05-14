package telegram

import (
	"bytes"
	"encoding/json"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/registry/image"
	"io"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
)

const (
	getMeMethod              = "getMe"
	updatesMethod            = "getUpdates"
	sendMessageMethod        = "sendMessage"
	sendPhotoMethod          = "sendPhoto"
	editMessageCaptionMethod = "editMessageCaption"
	editMessageMediaMethod   = "editMessageMedia"
	answerCallbackMethod     = "answerCallbackQuery"
)

type Client struct {
	scheme   string
	host     string
	basePath string
	client   http.Client
	ir       *image.ImageRegistry
}

func New(host string, token string, ir *image.ImageRegistry) (c *Client, err error) {
	defer func() { err = e.WrapIfErr("error getting me: ", err) }()

	c = &Client{
		scheme:   "https",
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
		ir:       ir,
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

	var b io.Reader = http.NoBody
	if reqBody != nil && !reflect.ValueOf(reqBody).IsNil() {
		b = reqBody
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), b)
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
