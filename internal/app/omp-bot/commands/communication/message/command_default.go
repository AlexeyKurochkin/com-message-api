package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

//Default bot command
func (c *MessageCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("MessageCommander.Help: error sending reply message to chat - %v", err)
	}
}
