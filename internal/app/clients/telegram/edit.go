package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/vi350/vk-internship/internal/app/e"
	"github.com/vi350/vk-internship/internal/app/localization"
	"github.com/vi350/vk-internship/internal/app/models"
	"io"
	"io/fs"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func (c *Client) EditMediaMessageByUser(userFromRegistry *models.User, message *Message, mType localization.MessageType) (err error) {
	defer func() { err = e.WrapIfErr("error editing message media", err) }()

	err = c.EditMediaMessage(userFromRegistry.ID, message.MessageID,
		localization.GetLocalizedText(mType, userFromRegistry.Language),
		localization.GetLocalizedImagePath(mType, userFromRegistry.Language, c.ir),
		GetLocalizedInlineKeyboardMarkup(mType, userFromRegistry.Language),
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
	var mes MessageResponse
	err = json.Unmarshal(data, &mes) // if not unmarshalled -> can't edit message
	return
}

type inputMediaPhoto struct {
	Type    string `json:"type"`
	Media   string `json:"media"`
	Caption string `json:"caption"`
}

func (c *Client) EditMediaMessage(chatID int64, messageID int64, caption string, image string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error editing message media", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(chatID, 10))
	values.Add("message_id", strconv.FormatInt(messageID, 10))
	inputMediaPhoto := inputMediaPhoto{
		Type:    "photo",
		Media:   "attach://photo",
		Caption: caption,
	}
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return
	}

	if _, err = os.Stat(image); err == nil {
		var file *os.File
		if file, err = os.Open(image); err != nil {
			return e.WrapIfErr("error opening file: ", err)
		}
		defer func() { _ = file.Close() }()

		buf := &bytes.Buffer{}
		writer := multipart.NewWriter(buf)

		var fileField io.Writer
		fileField, err = writer.CreateFormFile("photo", image)
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

		var jsonBytes []byte
		jsonBytes, err = json.Marshal(inputMediaPhoto)
		values.Add("media", string(jsonBytes))
		var data []byte
		if data, err = c.doRequest(editMessageMediaMethod, values, buf,
			map[string][]string{
				"Content-Type": {writer.FormDataContentType()},
			}); err != nil {
			return
		}

		var mes MessageResponse
		if err = json.Unmarshal(data, &mes); err != nil { // if not unmarshalled -> can't edit message
			return
		}
		c.ir.Save(image, mes.Result.Photo[len(mes.Result.Photo)-1].FileID)

	} else if errors.Is(err, fs.ErrNotExist) {
		if strings.HasPrefix(image, "AgACAg") {
			err = nil
			inputMediaPhoto.Media = image
			var jsonBytes []byte
			jsonBytes, err = json.Marshal(inputMediaPhoto)
			values.Add("media", string(jsonBytes))
			_, err = c.doRequest(editMessageMediaMethod, values, nil, nil)
		}
	}

	return
}
