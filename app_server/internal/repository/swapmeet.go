package repository

import (
	"app_server/internal/models"
	"app_server/pkg/db/postgres"
	"app_server/pkg/logger"
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DBSwapmeetRepo struct {
	db     *postgres.DB
	cache  *redis.Client
	logger logger.Logger
}

func NewSwapmeetRepo(ctx context.Context, db *postgres.DB, cache *redis.Client) *DBSwapmeetRepo {
	return &DBSwapmeetRepo{
		db:     db,
		cache:  cache,
		logger: logger.GetLoggerFromCtx(ctx),
	}
}

func (r *DBSwapmeetRepo) WarmUpCache(ctx context.Context) error {

	categories, err := r.getCategoriesFromDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch categories from DB: %w", err)
	}

	err = r.setCategoriesToCache(ctx, categories)
	if err != nil {
		return fmt.Errorf("failed to set categories in Redis: %w", err)
	}

	publishedAds, err := r.getPublishedAdvertisementsFromDB(ctx, nil)
	if err != nil {
		if errors.Is(err, models.ErrAdvertisementsNotFound) {
			return nil
		}
		return fmt.Errorf("failed to fetch advertisements from DB: %w", err)
	}

	err = r.setPublishedAdvertisementsToCache(ctx, publishedAds)
	if err != nil {
		return fmt.Errorf("failed to set advertisements in Redis: %w", err)
	}

	return nil
}

func (r *DBSwapmeetRepo) EnsureUserExists(ctx context.Context, userID, userName string) error {
	query := `
		INSERT INTO users (base_id, name) 
		VALUES ($1, $2) 
		ON CONFLICT (base_id) DO NOTHING
	`
	_, err := r.db.Db.ExecContext(ctx, query, userID, userName)
	return err
}
