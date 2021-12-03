package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/model"
	"time"
)

func (m MessageCommander) New(inputMsg *tgbotapi.Message) {
	messageData, error := CheckMessageInput(inputMsg.CommandArguments(), newArgumentRowsCount)
	text := ""
	if error != nil {
		text = fmt.Sprintf("Less then %v rows of values were provided", newArgumentRowsCount)
	} else {
		newMessageId := AddNewMessage(messageData, m)
		text = fmt.Sprintf("New message Id is: %v", newMessageId)
	}

	newBotMessage := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	m.bot.Send(newBotMessage)
}

func AddNewMessage(messageData []string, m MessageCommander) uint64 {
	message := CreateMessage(messageData)
	newIndex, _ := m.messageService.Create(&message)
	return newIndex
}

func CreateMessage(messageData []string) model.Message {
	message := model.Message{
		From:     messageData[0],
		To:       messageData[1],
		Text:     messageData[2],
		Datetime: time.Now(),
	}

	return message
}
