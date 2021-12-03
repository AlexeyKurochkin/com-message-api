package message

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
	"github.com/ozonmp/com-message-api/internal/model"
	"log"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *MessageCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData, err := ParseCallbackData(callbackPath)
	if err != nil {
		log.Printf("MessageCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	msg := BuildMessage(c, parsedData, callback)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat - %v", err)
	}
}

func BuildMessage(c *MessageCommander, parsedData CallbackListData, callback *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	values, boundsError := c.messageService.List(uint64(parsedData.Offset), messagesPerPage)
	isMessageListEnded := boundsError != nil
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, GetMessageText(isMessageListEnded, values))
	if !isMessageListEnded {
		numericKeyboard := GetNumericKeyboard(parsedData)
		msg.ReplyMarkup = numericKeyboard
	}

	return msg
}

func GetNumericKeyboard(parsedData CallbackListData) tgbotapi.InlineKeyboardMarkup {
	serializedData, _ := json.Marshal(CallbackListData{parsedData.Offset + messagesPerPage})
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Load more", fmt.Sprintf("communication__message__list__%v", string(serializedData))),
		),
	)

	return numericKeyboard
}

func GetMessageText(isNoMoreMessages bool, values []*model.Message) string {
	text := ""
	if isNoMoreMessages {
		text = "There are no more messages"
	} else {
		for i := range values {
			text += values[i].String() + "\n"
		}
	}

	return text
}

func ParseCallbackData(callbackPath path.CallbackPath) (CallbackListData, error) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	return parsedData, err
}
