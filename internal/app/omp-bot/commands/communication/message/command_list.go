package message

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (m MessageCommander) List(inputMsg *tgbotapi.Message) {
	values, _ := m.messageService.List(0, messagesPerPage)
	text := ""
	for i := range values {
		text += values[i].String() + "\n\n"
	}

	serializedData, _ := json.Marshal(CallbackListData{messagesPerPage})
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Load more", fmt.Sprintf("communication__message__list__%v", string(serializedData))),
		),
	)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	msg.ReplyMarkup = numericKeyboard
	_, err := m.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat - %v", err)
	}
}
