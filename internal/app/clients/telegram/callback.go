package telegram

import (
	"github.com/vi350/vk-internship/internal/app/e"
	"net/url"
)

func (c *Client) AnswerCallbackQuery(queryID string, notification string) (err error) {
	defer func() { err = e.WrapIfErr("error answering callback", err) }()

	values := url.Values{}
	values.Add("callback_query_id", queryID)
	if notification != "" {
		values.Add("text", notification)
	}

	_, err = c.doRequest(answerCallbackMethod, values, nil, nil)
	return
}
