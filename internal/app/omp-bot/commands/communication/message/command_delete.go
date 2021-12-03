package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (m MessageCommander) Delete(inputMsg *tgbotapi.Message) {
	arguments := inputMsg.CommandArguments()
	messageId, error := strconv.ParseUint(arguments, 0, 64)
	text := ""
	if error != nil {
		text = fmt.Sprintf("Incorrect message number")
	} else {
		successfullyDeleted, error := m.messageService.Remove(messageId)
		if error != nil {
			text = fmt.Sprintf("%v", error)
		}

		if successfullyDeleted {
			text = fmt.Sprintf("Successfully deleted message with messageId %v", messageId)
		}
	}

	message := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	_, error = m.bot.Send(message)
	if error != nil {
		log.Printf("Error sending message to chat %v", error)
	}
}
