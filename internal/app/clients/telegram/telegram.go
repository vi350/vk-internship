package telegram

import (
	"encoding/json"
	"github.com/vi350/vk-internship/internal/app/e"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getMeMethod       = "getMe"
	updatesMethod     = "getUpdates"
	sendMessageMethod = "sendMessage"
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
	data, err := c.doRequest(getMeMethod, values)
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

func (c *Client) doRequest(method string, values url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("error performing request: ", err) }()

	u := url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) SendTextMessage(ID int64, message string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error sending message: ", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(ID, 10))
	values.Add("text", message)
	if replyMarkup != nil {
		j, err := json.Marshal(replyMarkup)
		values.Add("reply_markup", string(j))
		if err != nil {
			return err
		}
	}

	_, err = c.doRequest(sendMessageMethod, values)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) Updates(offset int64, limit int) (updates []Update, err error) {
	defer func() { err = e.WrapIfErr("error getting updates: ", err) }()

	values := url.Values{}
	values.Add("offset", strconv.FormatInt(offset, 10))
	values.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(updatesMethod, values)
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
