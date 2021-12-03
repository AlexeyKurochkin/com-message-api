package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/model"
	"strconv"
)

func (m MessageCommander) Edit(inputMsg *tgbotapi.Message) {
	messageData, error := CheckMessageInput(inputMsg.CommandArguments(), editArgumentRowsCount)
	text := ""
	if error != nil {
		text = fmt.Sprintf("Less then %v rows of values were provided", editArgumentRowsCount)
	} else {
		message := CreateMessage(messageData[1:])
		lookupMessageIndex, error := strconv.ParseUint(messageData[0], 0, 64)
		if error != nil {
			text = "Incorrect id provided"
		} else {
			error := UpdateMessage(lookupMessageIndex, message, m.messageService)
			if error != nil {
				text = fmt.Sprintf("Message for update was not found")
			} else {
				text = fmt.Sprintf("Message was successfully updated")
			}
		}
	}

	newBotMessage := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	m.bot.Send(newBotMessage)
}

func UpdateMessage(id uint64, updateMessageData model.Message, messageService IMessageService) error {
	return messageService.Update(id, &updateMessageData)
}
