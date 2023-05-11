package telegram

import (
	"database/sql"
	"errors"
	tgClient "github.com/vi350/vk-internship/internal/app/clients/telegram"
	"github.com/vi350/vk-internship/internal/app/localization"
	"github.com/vi350/vk-internship/internal/app/storage/user_storage"
	"log"
	"strings"
)

func (ep *EventProcessor) doCommand(text string, userFromMessage tgClient.User) (err error) {
	text = strings.TrimSpace(text)

	log.Printf("user %s:\n sent: %s", userFromMessage.Username, text)

	var userFromStore *user_storage.User
	userFromStore, err = ep.userStorage.Read(userFromMessage.ID)

	switch {
	case strings.HasPrefix(text, "/start"):
		if errors.Is(err, sql.ErrNoRows) {
			if err = ep.userStorage.SaveFromTg(&userFromMessage, text); err != nil {
				return err
			}
			// TODO: send language selection message
			if err = ep.tgcli.SendTextMessage(userFromMessage.ID,
				localization.GetLocalizedText(localization.StartMessage, userFromMessage.LanguageCode),
				localization.GetLocalizedInlineKeyboardMarkup(localization.StartMessage, userFromMessage.LanguageCode)); err != nil {
				return err
			}
			return nil
		} else if err != nil {
			if err = ep.userStorage.SetState(userFromMessage.ID, user_storage.MainMenu); err != nil {
				return err
			}
			if err = ep.tgcli.SendTextMessage(userFromMessage.ID,
				localization.GetLocalizedText(localization.StartMessage, userFromStore.Language),
				localization.GetLocalizedInlineKeyboardMarkup(localization.StartMessage, userFromStore.Language)); err != nil {
				return err
			}
			return nil
		}
		return err

	case strings.HasPrefix(text, "/help"):
		if err = ep.tgcli.SendTextMessage(userFromMessage.ID,
			localization.GetLocalizedText(localization.HelpMessage, userFromStore.Language),
			localization.GetLocalizedInlineKeyboardMarkup(localization.HelpMessage, userFromStore.Language)); err != nil {
			return err
		}

	default:
		if err = ep.tgcli.SendTextMessage(userFromMessage.ID,
			localization.GetLocalizedText(localization.UnknownCommandMessage, userFromStore.Language),
			localization.GetLocalizedInlineKeyboardMarkup(localization.UnknownCommandMessage, userFromStore.Language)); err != nil {
			return err
		}
	}

	return nil
}
