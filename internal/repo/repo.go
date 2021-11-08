package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/com-message-api/internal/model"
)

var ids uint64

// Repo is DAO for Message
type Repo interface {
	CreateMessage(ctx context.Context, message *model.Message) (uint64, error)
	DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error)
	ListMessage(ctx context.Context) ([]*model.Message, error)
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
	ids++
	return ids, nil
}

func (r *repo) DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error) {
	return nil, nil
}

func (r *repo) ListMessage(ctx context.Context) ([]*model.Message, error) {
	return nil, nil
}
func (r *repo) RemoveMessage(ctx context.Context, messageID uint64) (bool, error) {
	return true, nil
}
