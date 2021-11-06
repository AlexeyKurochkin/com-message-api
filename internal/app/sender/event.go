package sender

import "github.com/ozonmp/com-message-api/internal/model"

//EventSender interface
type EventSender interface {
	Send(messageEvent *model.MessageEvent) error
}
