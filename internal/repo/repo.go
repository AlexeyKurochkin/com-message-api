package repo

import (
	"context"

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
		Columns("from", "to", "text", "datetime").
		Values(message.From, message.To, message.Text, message.Datetime).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var id uint64
	err = r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error) {
	query, args, err := psql.Select("id", "from", "to", "text", "datetime", "removed", "created_at", "updated_at").
		From("messages").
		Where(sq.Eq{"id": messageID}).ToSql()

	if err != nil {
		return nil, err
	}

	message := model.Message{}
	err = r.db.GetContext(ctx, &message, query, args...)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *repo) ListMessage(ctx context.Context) ([]model.Message, error) {
	query, args, err := psql.Select("id", "from", "to", "text", "datetime", "removed", "created_at", "updated_at").
		From("messages").OrderBy("id ASC").
		ToSql()

	if err != nil {
		return nil, err
	}

	var messages []model.Message
	err = r.db.SelectContext(ctx, &messages, query, args...)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
func (r *repo) RemoveMessage(ctx context.Context, messageID uint64) (bool, error) {
	query, args, err := psql.Delete("messages").
		Where(sq.Eq{"id": messageID}).
		ToSql()

	if err != nil {
		return false, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return false, err
	}

	return true, nil
}
