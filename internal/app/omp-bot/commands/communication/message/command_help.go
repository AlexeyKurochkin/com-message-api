package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (m MessageCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__communication__message - help\n"+
			"/get__communication__message {messageId} - get message\n"+
			"/list__communication__message - list messages\n"+
			"/delete__communication__message {mmessageId} - delete message\n"+
			"/new__communication__message - create message\n"+
			"{From}\n"+
			"{To}\n"+
			"{Text}\n"+
			"/edit__communication__message - edit message\n"+
			"{messageId}\n"+
			"{From}\n"+
			"{To}\n"+
			"{Text}\n",
	)

	_, err := m.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
