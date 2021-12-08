package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

//Delete handles bot Delete command
func (m Commander) Delete(inputMsg *tgbotapi.Message) {
	arguments := inputMsg.CommandArguments()
	messageID, error := strconv.ParseUint(arguments, 0, 64)
	text := ""
	if error != nil {
		text = fmt.Sprintf("Incorrect message number")
	} else {
		successfullyDeleted, serviceError := m.messageService.Remove(messageID)
		if serviceError != nil {
			text = fmt.Sprintf("%v", serviceError)
		}

		if successfullyDeleted {
			text = fmt.Sprintf("Successfully deleted message with messageID %v", messageID)
		}
	}

	message := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	_, error = m.bot.Send(message)
	if error != nil {
		log.Printf("Error sending message to chat %v", error)
	}
}
