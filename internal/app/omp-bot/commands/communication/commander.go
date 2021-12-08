package communication

import (
	"github.com/ozonmp/com-message-api/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/commands/communication/message"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
)

const messageString = "message"

//ICommander for handling bot messages
type ICommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

//Commander subdomain commander
type Commander struct {
	bot              *tgbotapi.BotAPI
	messageCommander ICommander
}

//NewCommunicationCommander constructor for Commander
func NewCommunicationCommander(
	bot *tgbotapi.BotAPI,
	cfg *config.Config,
) *Commander {
	return &Commander{
		bot: bot,
		// subdomainCommander
		messageCommander: message.NewMessageCommander(bot, cfg),
	}
}

//HandleCallback handles bot callback
func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case messageString:
		c.messageCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Commander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

//HandleCommand handles bot command
func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case messageString:
		c.messageCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Commander.HandleCommand: unknown subdomain - #{commandPath.Subdomain}")
	}
}
