package repo

import (
	"context"
	"database/sql"
	"food-recipes-backend/internal/queries"
)

type IKeyRepository interface {
	UpsertKey(ctx context.Context, upsertParams queries.UpsertRefreshTokenParams) (queries.Key, error)
	RemoveRefreshToken(ctx context.Context, userID int) error
}

type keyRepository struct {
	queries *queries.Queries
}

func NewKeyRepository(db *sql.DB) IKeyRepository {
	return &keyRepository{
		queries: queries.New(db),
	}
}

func (kr *keyRepository) RemoveRefreshToken(ctx context.Context, userID int) error {
	err := kr.queries.RemoveRefreshToken(ctx, int32(userID))
	if err != nil {
		return err
	}
	return nil
}

func (kr *keyRepository) UpsertKey(ctx context.Context, upsertParams queries.UpsertRefreshTokenParams) (queries.Key, error) {
	key, err := kr.queries.UpsertRefreshToken(ctx, upsertParams)
	if err != nil {
		return queries.Key{}, err
	}
	return key, nil
}