package communication

import (
	"github.com/ozonmp/com-message-api/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/commands/communication/message"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
)

const messageString = "message"

//Commander for handling bot messages
type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

//CommunicationCommander subdomain commander
type CommunicationCommander struct {
	bot              *tgbotapi.BotAPI
	messageCommander Commander
}

//NewCommunicationCommander constructor for CommunicationCommander
func NewCommunicationCommander(
	bot *tgbotapi.BotAPI,
	cfg *config.Config,
) *CommunicationCommander {
	return &CommunicationCommander{
		bot: bot,
		// subdomainCommander
		messageCommander: message.NewMessageCommander(bot, cfg),
	}
}

//HandleCallback handles bot callback
func (c *CommunicationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case messageString:
		c.messageCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CommunicationCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

//HandleCommand handles bot command
func (c *CommunicationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case messageString:
		c.messageCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CommunicationCommander.HandleCommand: unknown subdomain - #{commandPath.Subdomain}")
	}
}
