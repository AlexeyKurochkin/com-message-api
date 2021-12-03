package communication

import (
	"github.com/ozonmp/com-message-api/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/commands/communication/message"
	"github.com/ozonmp/com-message-api/internal/app/omp-bot/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CommunicationCommander struct {
	bot              *tgbotapi.BotAPI
	messageCommander Commander
}

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

func (c *CommunicationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "message":
		c.messageCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CommunicationCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CommunicationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "message":
		c.messageCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CommunicationCommander.HandleCommand: unknown subdomain - #{commandPath.Subdomain}")
	}
}
