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

const (
	deferredString  = "deferred"
	processedString = "processed"
	newString       = "new"
	lockString      = "lock"
	unlockString    = "unlock"
	unknownString   = "unknown"
)

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
	MessageID uint64       `db:"message_id"`
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
		return deferredString
	case Processed:
		return processedString
	case New:
		return newString
	case Lock:
		return lockString
	case Unlock:
		return unlockString
	default:
		return unknownString
	}
}

//Scan implementation of scan interface
func (e *EventStatus) Scan(value interface{}) error {
	var result EventStatus
	switch value {
	case newString:
		result = New
	case deferredString:
		result = Deferred
	case processedString:
		result = Processed
	case lockString:
		result = Lock
	case unlockString:
		result = Unlock
	}

	*e = result
	return nil
}
