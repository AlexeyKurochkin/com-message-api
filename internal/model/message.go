package model

import (
	"fmt"
	"time"
)

//Message type
type Message struct {
	ID       uint64    `db:"id"`
	From     string    `db:"from"`
	To       string    `db:"to"`
	Text     string    `db:"text"`
	Datetime time.Time `db:"datetime"`
}

//EventType type of message event
type EventType uint8

//EventStatus status of message event
type EventStatus uint8

//EventType enum
const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

//MessageEvent type
type MessageEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Message
}

func (m MessageEvent) String() string {
	return fmt.Sprintf("EventId: %v\nEntityId: %v", m.ID, m.Entity.ID)
}
