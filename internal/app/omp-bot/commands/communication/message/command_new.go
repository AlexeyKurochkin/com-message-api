package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/model"
	"github.com/rs/zerolog/log"
	"time"
)

//New handles bot New command
func (m MessageCommander) New(inputMsg *tgbotapi.Message) {
	messageData, error := checkMessageInput(inputMsg.CommandArguments(), newArgumentRowsCount)
	text := ""
	if error != nil {
		text = fmt.Sprintf("Less then %v rows of values were provided", newArgumentRowsCount)
	} else {
		newMessageID := addNewMessage(messageData, m)
		text = fmt.Sprintf("New message Id is: %v", newMessageID)
	}

	newBotMessage := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	_, err := m.bot.Send(newBotMessage)
	if err != nil {
		log.Error().Err(err).Msg("Cannot send message.Command_new")
	}
}

func addNewMessage(messageData []string, m MessageCommander) uint64 {
	message := createMessage(messageData)
	newIndex, _ := m.messageService.Create(&message)
	return newIndex
}

func createMessage(messageData []string) model.Message {
	message := model.Message{
		From:     messageData[0],
		To:       messageData[1],
		Text:     messageData[2],
		Datetime: time.Now(),
	}

	return message
}
