package telegram

import (
	"database/sql"
	"errors"
	"github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/storage/user_storage"
	"log"
	"strings"
	"time"
)

func (ep *EventProcessor) doCommand(text string, userFromMessage telegram.User) (err error) {
	text = strings.TrimSpace(text)

	log.Printf("user %s:\n sent: %s", userFromMessage.Username, text)

	var userFromStore *user_storage.User
	userFromStore, err = ep.userStorage.Read(userFromMessage.ID)

	switch {
	case strings.HasPrefix(text, "/start"):
		if errors.Is(err, sql.ErrNoRows) {
			if userFromMessage.LanguageCode == "" {
				userFromMessage.LanguageCode = "en"
			}
			if len(text) < 6 {
				text = ""
			}
			if err = ep.userStorage.Save(&user_storage.User{
				ID:        userFromMessage.ID,
				FirstName: userFromMessage.FirstName,
				Username:  userFromMessage.Username,
				StartDate: time.Now().Unix(),
				Language:  userFromMessage.LanguageCode,
				State:     user_storage.MainMenu,
				Refer:     text[7:],
			}); err != nil {
				return err
			}
			// TODO: send language selection message
			if err = ep.tgcli.SendTextMessage(userFromMessage.ID, GetLocalizedText(startMessage, userFromStore.Language), nil); err != nil {
				return err
			}
			return nil
		} else if err != nil {
			if err = ep.userStorage.SetState(userFromMessage.ID, user_storage.MainMenu); err != nil {
				return err
			}
			if err = ep.tgcli.SendTextMessage(userFromMessage.ID, GetLocalizedText(startMessage, userFromStore.Language), nil); err != nil {
				return err
			}
			return nil
		}
		return err

	case strings.HasPrefix(text, "/help"):
		if err = ep.tgcli.SendTextMessage(userFromMessage.ID, GetLocalizedText(helpMessage, userFromStore.Language), nil); err != nil {
			return err
		}

	default:
		if err = ep.tgcli.SendTextMessage(userFromMessage.ID, GetLocalizedText(helpMessage, userFromStore.Language), nil); err != nil {
			return err
		}
	}

	return nil
}
