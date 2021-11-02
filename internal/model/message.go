package model

import "fmt"

type Message struct {
	ID  uint64 `db:"id"`
	Foo uint64 `db:"foo"`
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type MessageEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Message
}

func (m MessageEvent) String() string {
	return fmt.Sprintf("EventId: %v\nEntityId: %v", m.ID, m.Entity.ID)
}



