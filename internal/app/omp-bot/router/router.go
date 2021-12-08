package router

import (
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/commands/communication"
	"github.com/ozonmp/com-message-api/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
)

const communicationString = "communication"

//Commander interface for handling callbacks and commands
type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

//Router for bot subdomains requests
type Router struct {
	// bot
	bot                    *tgbotapi.BotAPI
	communicationCommander Commander
}

//NewRouter constructor for Router
func NewRouter(
	bot *tgbotapi.BotAPI,
	cfg *config.Config,
) *Router {
	return &Router{
		// bot
		bot:                    bot,
		communicationCommander: communication.NewCommunicationCommander(bot, cfg),
	}
}

//HandleUpdate handles bot update
func (c *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case communicationString:
		c.communicationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Domain {
	case communicationString:
		c.communicationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
