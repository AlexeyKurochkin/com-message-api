package model

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"time"
)

//Message type
type Message struct {
	ID       uint64      `db:"id"`
	From     string      `db:"from"`
	To       string      `db:"to"`
	Text     string      `db:"text"`
	Datetime time.Time   `db:"datetime"`
	Removed  bool        `db:"removed"`
	Created  time.Time   `db:"created"`
	Updated  pq.NullTime `db:"updated"`
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
	Unknown

	Deferred EventStatus = iota
	Processed
	New
	Lock
	Unlock
)

//MessageEvent type
type MessageEvent struct {
	ID        uint64       `db:"id"`
	MessageId uint64       `db:"message_id"`
	TypeDb    string       `db:"type"`
	Status    EventStatus  `db:"status"`
	Payload   string       `db:"payload"`
	Updated   sql.NullTime `db:"updated"`
	Type      EventType
	Entity    *Message
}

func (m MessageEvent) String() string {
	return fmt.Sprintf("EventId: %v\nEntityId: %v", m.ID, m.Entity.ID)
}

func (e EventType) String() string {
	switch e {
	case Created:
		return "created"
	case Removed:
		return "removed"
	case Updated:
		return "updated"
	default:
		return "unknown"
	}
}

func (e EventStatus) String() string {
	switch e {
	case Deferred:
		return "deferred"
	case Processed:
		return "processed"
	case New:
		return "new"
	case Lock:
		return "lock"
	case Unlock:
		return "unlock"
	default:
		return "unknown"
	}
}

func (e *EventStatus) Scan(value interface{}) error {
	var result EventStatus
	switch value {
	case "new":
		result = New
	case "deferred":
		result = Deferred
	case "processed":
		result = Processed
	case "lock":
		result = Lock
	case "Unlock":
		result = Unlock
	}

	*e = result
	return nil
}
