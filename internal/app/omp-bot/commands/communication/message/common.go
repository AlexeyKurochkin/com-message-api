package message

import (
	"fmt"
	"github.com/ozonmp/com-message-api/internal/model"
	"strings"
)

const newArgumentRowsCount = 3
const editArgumentRowsCount = 4
const messagesPerPage uint64 = 5
const list = "list"
const help = "help"
const get = "get"
const delete = "delete"
const new = "new"
const edit = "edit"

//IMessageService interface for working with messages
type IMessageService interface {
	Describe(messageID uint64) (*model.Message, error)
	List(cursor uint64, limit uint64) ([]*model.Message, error)
	Create(*model.Message) (uint64, error)
	Update(messageID uint64, message *model.Message) error
	Remove(messageID uint64) (bool, error)
}

func checkMessageInput(commandData string, argumentsRowsCount int) ([]string, error) {
	messageData := strings.Split(commandData, "\n")
	if len(messageData) != argumentsRowsCount {
		return nil, fmt.Errorf("Less then %v rows of values were provided", argumentsRowsCount)
	}

	return messageData, nil
}
