package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/com-message-api/internal/model"
)

// Repo is DAO for Message
type Repo interface {
	DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeMessage(ctx context.Context, messageID uint64) (*model.Message, error) {
	return nil, nil
}
