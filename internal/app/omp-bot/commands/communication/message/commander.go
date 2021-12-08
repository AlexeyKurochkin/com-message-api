package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
	"github.com/ozonmp/com-message-api/internal/config"
	"github.com/ozonmp/com-message-api/internal/service/communication/message"
	"log"
)

//IMessageCommander interface for bot commands
type IMessageCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

//MessageCommander type for handling bot commands and callbacks
type MessageCommander struct {
	bot            *tgbotapi.BotAPI
	messageService IMessageService
}

//NewMessageCommander constructor for MessageCommander
func NewMessageCommander(bot *tgbotapi.BotAPI, cfg *config.Config) *MessageCommander {
	messageService := message.NewMessageService(cfg)
	return &MessageCommander{
		bot:            bot,
		messageService: messageService,
	}
}

//HandleCallback handle bot callback
func (m MessageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case list:
		m.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

//HandleCommand handles bot command
func (m MessageCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case help:
		m.Help(message)
	case list:
		m.List(message)
	case get:
		m.Get(message)
	case delete:
		m.Delete(message)
	case new:
		m.New(message)
	case edit:
		m.Edit(message)
	default:
		m.Default(message)
	}
}
