package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

//Get handles bot Get command
func (m MessageCommander) Get(inputMsg *tgbotapi.Message) {
	arguments := inputMsg.CommandArguments()
	index, error := strconv.ParseUint(arguments, 0, 64)
	text := ""
	if error != nil {
		text = "Wrong id provided"
	} else {
		message, serviceError := m.messageService.Describe(index)
		if serviceError != nil {
			text = "Message with such id does not exist"
		} else {
			text = message.String()
		}

	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	_, error = m.bot.Send(msg)
	if error != nil {
		log.Printf("Error sending message to chat %v", error)
	}
}
