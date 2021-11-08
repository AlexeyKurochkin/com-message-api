package repo

import "github.com/ozonmp/com-message-api/internal/model"

//EventRepo interface
type EventRepo interface {
	Lock(n uint64) ([]model.MessageEvent, error)
	Unlock(eventIDs []uint64) error
	Add(event []model.MessageEvent) error
	Remove(eventIDs []uint64) error
}
