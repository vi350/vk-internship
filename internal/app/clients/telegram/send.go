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

func (c *Client) SendTextMessageByUser(userFromRegistry *models.User, mType localization.MessageType) (err error) {
	defer func() { err = e.WrapIfErr("error sending text message", err) }()

	err = c.SendTextMessage(userFromRegistry.ID,
		localization.GetLocalizedText(mType, userFromRegistry.Language),
		GetLocalizedInlineKeyboardMarkup(mType, userFromRegistry.Language),
	)

	return
}

func (c *Client) SendTextMessage(ID int64, text string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error sending text message", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(ID, 10))
	values.Add("text", text)
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return err
	}

	_, err = c.doRequest(sendMessageMethod, values, nil, nil)
	return err
}

func (c *Client) SendImageByUser(userFromRegistry *models.User, mType localization.MessageType) (err error) {
	defer func() { err = e.WrapIfErr("error sending image message", err) }()

	err = c.SendImage(userFromRegistry.ID,
		localization.GetLocalizedText(mType, userFromRegistry.Language),
		localization.GetLocalizedImagePath(mType, userFromRegistry.Language, c.ir),
		GetLocalizedInlineKeyboardMarkup(mType, userFromRegistry.Language),
	)

	return
}

func (c *Client) SendImage(ID int64, text string, image string, replyMarkup ReplyMarkup) (err error) {
	defer func() { err = e.WrapIfErr("error sending image message", err) }()

	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(ID, 10))
	values.Add("caption", text)
	if err = addIfReplyMarkup(&values, replyMarkup); err != nil {
		return err
	}

	if _, err = os.Stat(image); err == nil {
		var file *os.File
		if file, err = os.Open(image); err != nil {
			return e.WrapIfErr("error opening file: ", err)
		}
		defer func() { _ = file.Close() }()

		buf := &bytes.Buffer{}
		writer := multipart.NewWriter(buf)

		var imagePart io.Writer
		imagePart, err = writer.CreateFormFile("photo", image)
		if err != nil {
			return e.WrapIfErr("Failed to create form field:", err)
		}

		_, err = io.Copy(imagePart, file)
		if err != nil {
			return e.WrapIfErr("Failed to copy file data:", err)
		}

		if err = writer.Close(); err != nil {
			return e.WrapIfErr("error closing writer: ", err)
		}

		var data []byte
		if data, err = c.doRequest(sendPhotoMethod, values, buf,
			map[string][]string{
				"Content-Type": {writer.FormDataContentType()},
			}); err != nil {
			return
		}

		var mes MessageResponse
		if err = json.Unmarshal(data, &mes); err != nil {
			return
		}
		c.ir.Save(image, mes.Result.Photo[len(mes.Result.Photo)-1].FileID)

	} else if errors.Is(err, fs.ErrNotExist) {
		if strings.HasPrefix(image, "AgACAg") {
			err = nil
			values.Add("photo", image)
			_, err = c.doRequest(sendPhotoMethod, values, nil, nil)
		}
	}

	return
}
