package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/localization"
	userStorage "github.com/vi350/vk-internship/internal/app/storage/user"
	"io"
	"io/fs"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
)

func (c *Client) EditMediaMessageByUser(userFromRegistry *userStorage.User, message Message, mType localization.MessageType) (err error) {
	defer func() { err = e.WrapIfErr("error editing message media", err) }()

	err = c.EditMessageMedia(userFromRegistry.ID, message.MessageID,
		localization.GetLocalizedImagePath(mType, userFromRegistry.Language, c.ir),
		nil, // supposed we update reply markup with caption update
	)
	err = c.EditMessageCaption(userFromRegistry.ID, message.MessageID,
		localization.GetLocalizedText(mType, userFromRegistry.Language),
		localization.GetLocalizedInlineKeyboardMarkup(mType, userFromRegistry.Language),
	)

	return
}

func (c *Client) EditMessageCaption(chatID int64, messageID int64, caption string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error editing message caption", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(chatID, 10))
	values.Add("message_id", strconv.FormatInt(messageID, 10))
	values.Add("caption", caption)
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return
	}

	var data []byte
	if data, err = c.doRequest(editMessageCaptionMethod, values, nil, nil); err != nil {
		return
	}
	var mes Message
	err = json.Unmarshal(data, &mes) // if not unmarshalled -> can't edit message
	return
}

func (c *Client) EditMessageMedia(chatID int64, messageID int64, image string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error editing message media", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(chatID, 10))
	values.Add("message_id", strconv.FormatInt(messageID, 10))
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return
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

		var data []byte
		if data, err = c.doRequest(editMessageMediaMethod, values, &buf,
			map[string][]string{
				"Content-Type": {writer.FormDataContentType()},
			}); err != nil {
			return
		}

		var mes Message
		if err = json.Unmarshal(data, &mes); err != nil { // if not unmarshalled -> can't edit message
			return
		}
		c.ir.Save(image, mes.Photo[len(mes.Photo)-1].FileID)

	} else if errors.Is(err, fs.ErrNotExist) {
		err = nil
		values.Add("photo", image)
		_, err = c.doRequest(editMessageMediaMethod, values, nil, nil)
	} else {
		return
	}

	return
}
