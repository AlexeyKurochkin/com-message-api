package message

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
	"github.com/ozonmp/com-message-api/internal/model"
	"log"
)

type callbackListData struct {
	Offset uint64 `json:"offset"`
}

//CallbackList handles bot list callback
func (c *Commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData, err := parseCallbackData(callbackPath)
	if err != nil {
		log.Printf("Commander.CallbackList: "+
			"error reading json data for type callbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	msg := buildMessage(c, parsedData, callback)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat - %v", err)
	}
}

func buildMessage(c *Commander, parsedData callbackListData, callback *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	values, boundsError := c.messageService.List(uint64(parsedData.Offset), messagesPerPage)
	isMessageListEnded := boundsError != nil
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, getMessageText(isMessageListEnded, values))
	if !isMessageListEnded {
		numericKeyboard := getNumericKeyboard(parsedData)
		msg.ReplyMarkup = numericKeyboard
	}

	return msg
}

func getNumericKeyboard(parsedData callbackListData) tgbotapi.InlineKeyboardMarkup {
	serializedData, _ := json.Marshal(callbackListData{parsedData.Offset + messagesPerPage})
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Load more", fmt.Sprintf("communication__message__list__%v", string(serializedData))),
		),
	)

	return numericKeyboard
}

func getMessageText(isNoMoreMessages bool, values []*model.Message) string {
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

func parseCallbackData(callbackPath path.CallbackPath) (callbackListData, error) {
	parsedData := callbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	return parsedData, err
}
