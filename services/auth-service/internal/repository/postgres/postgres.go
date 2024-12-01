package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Storage struct {
	connectionPool *pgxpool.Pool
}

func New(ctx context.Context, url string) (*Storage, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to postgres")
	}

	return &Storage{pool}, nil
}

func (s *Storage)
