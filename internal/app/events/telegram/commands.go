package telegram

import (
	"github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/storage"
	"log"
	"strings"
	"time"
)

func (ep *EventProcessor) doCommand(text string, user telegram.User) (err error) {
	text = strings.TrimSpace(text)

	log.Printf("user %s:\n sent: %s", user.Username, text)

	switch {
	case strings.HasPrefix(text, "/start"):
		var exists bool
		if exists, err = ep.storage.IsExistUser(user.ID); err != nil {
			return err
		} else if exists {
			if err = ep.storage.SetStateUser(user.ID, storage.MainMenu); err != nil {
				return err
			}
			if err = ep.tgcli.SendTextMessage(user.ID, startMessage); err != nil {
				return err
			}
			return nil
		} else {
			if user.LanguageCode == "" {
				user.LanguageCode = "en"
			}
			if len(text) > 6 {
				text = text[7:]
			}
			userToStore := storage.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				Username:  user.Username,
				StartDate: time.Now().Unix(),
				Language:  user.LanguageCode,
				State:     storage.MainMenu,
				Refer:     text,
			}
			if err = ep.storage.SaveUser(&userToStore); err != nil {
				return err
			}
			if err = ep.tgcli.SendTextMessage(user.ID, startMessage); err != nil {
				return err
			}
		}

	case strings.HasPrefix(text, "/help"):
		if err = ep.tgcli.SendTextMessage(user.ID, helpMessage); err != nil {
			return err
		}

	default:
		if err = ep.tgcli.SendTextMessage(user.ID, unknownCommandMessage); err != nil {
			return err
		}
	}

	return nil
}
