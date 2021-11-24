package repo

import (
	"context"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
	"github.com/ozonmp/com-message-api/internal/model"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// Repo is DAO for Message
type Repo interface {
	CreateMessage(ctx context.Context, message *model.Message) (uint64, error)
	DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error)
	ListMessage(ctx context.Context) ([]model.Message, error)
	RemoveMessage(ctx context.Context, messageID uint64) (bool, error)
	UpdateMessage(ctx context.Context, message *model.Message) (*model.Message, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) CreateMessage(ctx context.Context, message *model.Message) (uint64, error) {
	query, args, err := psql.Insert("messages").
		Columns("\"from\"", "\"to\"", "text", "datetime").
		Values(message.From, message.To, message.Text, message.Datetime).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return 0, errors.Wrap(err, "Error on creating CreateMessage query")
	}

	var id uint64
	err = r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		return 0, errors.Wrap(err, "Error on executing CreateMessage query")
	}

	return id, nil
}

func (r *repo) DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error) {
	query, args, err := psql.Select("id", "\"from\"", "\"to\"", "text", "datetime", "removed", "created", "updated").
		From("messages").
		Where(sq.Eq{"id": messageID}).ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "Error on creating DescribeMessage query")
	}

	message := model.Message{}
	err = r.db.GetContext(ctx, &message, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "Error on executing DescribeMessage query")
	}

	return &message, nil
}

func (r *repo) ListMessage(ctx context.Context) ([]model.Message, error) {
	query, args, err := psql.Select("id", "\"from\"", "\"to\"", "text", "datetime", "removed", "created", "updated").
		From("messages").OrderBy("id ASC").
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "Error on creating ListMessage query")
	}

	var messages []model.Message
	err = r.db.SelectContext(ctx, &messages, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "Error on executing ListMessage query")
	}

	return messages, nil
}
func (r *repo) RemoveMessage(ctx context.Context, messageID uint64) (bool, error) {
	query, args, err := psql.Delete("messages").
		Where(sq.Eq{"id": messageID}).
		ToSql()

	if err != nil {
		return false, errors.Wrap(err, "Error on creating RemoveMessage query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "Error on executing RemoveMessage query")
	}

	return true, nil
}

func (r *repo) UpdateMessage(ctx context.Context, message *model.Message) (*model.Message, error) {
	query, args, err := psql.Update("messages").
		Set("\"from\"", message.From).
		Set("\"to\"", message.To).
		Set("text", message.Text).
		Set("datetime", message.Datetime).
		Where(sq.Eq{"id": message.ID}).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "Error on creating UpdateMessage query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Error on executing UpdateMessage query")
	}

	return message, nil
}
