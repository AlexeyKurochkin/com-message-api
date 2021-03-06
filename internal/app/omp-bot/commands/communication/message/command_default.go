package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

//Default bot command
func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}
