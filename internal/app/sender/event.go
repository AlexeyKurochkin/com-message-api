package sender

import "github.com/ozonmp/com-message-api/internal/model"

type EventSender interface {
	Send(messageEvent *model.MessageEvent) error
}
