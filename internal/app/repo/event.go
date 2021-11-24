package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/com-message-api/internal/metrics"
	"github.com/ozonmp/com-message-api/internal/model"
	pb "github.com/ozonmp/com-message-api/pkg/com-message-api"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	lockedStatus   = "lock"
	unlockedStatus = "unlock"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

//EventRepo interface
type EventRepo interface {
	Lock(n uint64) ([]model.MessageEvent, error)
	Unlock(eventIDs []uint64) error
	Remove(eventIDs []uint64) error
	Add(event model.MessageEvent) error
}

//NewEventRepo constructor for EventRepo
func NewEventRepo(db *sqlx.DB) EventRepo {
	return &repo{db: db}
}

type repo struct {
	db *sqlx.DB
}

func (r *repo) Lock(n uint64) ([]model.MessageEvent, error) {
	query, args, err := psql.Select("id", "message_id", "type", "status", "payload", "updated").
		From("messages_events").Where(sq.NotEq{"status": lockedStatus}).
		OrderBy("id ASC").ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "Error on creating lock sql query")
	}

	var messagesEvents []model.MessageEvent
	err = r.db.Select(&messagesEvents, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Error on executing lock query")
	}

	ids := getEventIds(messagesEvents)
	updateQuery, updateArgs, updateErr := psql.Update("messages_events").
		Set("status", lockedStatus).
		Where(sq.Eq{"id": ids}).ToSql()
	if updateErr != nil {
		return nil, errors.Wrap(updateErr, "Error on creating sql update query for locked items statuses")
	}

	_, updateErr = r.db.Exec(updateQuery, updateArgs...)
	if updateErr != nil {
		return nil, errors.Wrap(updateErr, "Error on executing update sql query")
	}

	metrics.EventsTotal.WithLabelValues("lock").Inc()

	return messagesEvents, nil
}

func (r *repo) Unlock(eventIDs []uint64) error {
	query, args, err := psql.Update("messages_events").
		Set("status", unlockedStatus).
		Where(sq.Eq{"id": eventIDs}).ToSql()
	if err != nil {
		return errors.Wrap(err, "Error on creating unlock sql query")
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "Error on executing unlock sql query")
	}

	metrics.EventsTotal.WithLabelValues("unlock").Inc()

	return nil
}

func (r *repo) Remove(eventIDs []uint64) error {
	query, args, err := psql.Delete("messages_events").
		Where(sq.Eq{"id": eventIDs}).ToSql()
	if err != nil {
		return errors.Wrap(err, "Error on creating remove query")
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "Error on executing remove query")
	}

	metrics.EventsTotal.WithLabelValues("remove").Inc()

	return nil
}

func getEventIds(events []model.MessageEvent) []uint64 {
	ids := make([]uint64, len(events))
	for i := 0; i < len(events); i++ {
		ids[i] = events[i].ID
	}

	return ids
}

func (r *repo) Add(event model.MessageEvent) error {
	pbMessage := &pb.Message{
		Id:       event.MessageId,
		From:     event.Entity.From,
		To:       event.Entity.To,
		Text:     event.Entity.Text,
		Datetime: timestamppb.New(event.Entity.Datetime),
	}

	var payload, err = protojson.Marshal(pbMessage)
	if err != nil {
		return errors.Wrap(err, "Error on marshalling proto message into json")
	}

	query, args, insertErr := psql.Insert("messages_events").
		Columns("message_id", "type", "status", "payload", "updated").
		Values(event.MessageId, event.TypeDb, event.Status.String(), payload, time.Now()).ToSql()
	if insertErr != nil {
		return errors.Wrap(insertErr, "Error on creating sql query for adding messages")
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "Error on executing query for adding messages")
	}

	metrics.EventsTotal.WithLabelValues("add").Inc()

	return nil
}
