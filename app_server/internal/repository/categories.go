package repository

import (
	"app_server/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *DBSwapmeetRepo) GetCategories(ctx context.Context) ([]models.Category, error) {

	categories, err := r.getCategoriesFromCache(ctx)
	if err == nil {
		return categories, nil
	}

	categories, err = r.getCategoriesFromDB(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.setCategoriesToCache(ctx, categories); err != nil {
		r.logger.Info(ctx, fmt.Sprintf("failed to set categories to cache: %v", err))
	}

	return categories, nil
}

func (r *DBSwapmeetRepo) getCategoriesFromDB(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := `
		SELECT 
   			id, 
    		name, 
    		COALESCE(parent_id, 0) AS parent_id 
		FROM categories
		ORDER BY id, parent_id
	`
	err := r.db.Db.SelectContext(ctx, &categories, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCategoriesNotFound
		}
		return nil, models.ErrDBQuery
	}
	return categories, nil
}

func (r *DBSwapmeetRepo) getCategoriesFromCache(ctx context.Context) ([]models.Category, error) {
	cachedCategories, err := r.cache.Get(ctx, "categories").Result()
	if err != nil {
		return nil, err
	}
	return unmarshalCategories(cachedCategories)
}

func (r *DBSwapmeetRepo) setCategoriesToCache(ctx context.Context, categories []models.Category) error {
	data, err := marshalCategories(categories)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, "categories", data, 0).Err()
}

func (r *DBSwapmeetRepo) CreateCategory(ctx context.Context, userID string, name string, parentID int) (*models.Category, error) {
	category, err := r.createCategoryInDB(ctx, userID, name, parentID)
	if err != nil {
		return nil, err
	}
	err = r.appendCategoryToCache(ctx, *category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *DBSwapmeetRepo) createCategoryInDB(ctx context.Context, userID string, name string, parentID int) (*models.Category, error) {
	var category models.Category
	query := `
		INSERT INTO categories (creator, name, parent_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, 
                 COALESCE(CAST(parent_id AS TEXT), '0') AS parent_id
		`
	err := r.db.Db.QueryRowContext(ctx, query, userID, name, parentID).Scan(
		&category.ID,
		&category.Name,
		&category.ParentID)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *DBSwapmeetRepo) appendCategoryToCache(ctx context.Context, category models.Category) error {
	data, err := marshalCategory(category)
	if err != nil {
		return err
	}
	return r.cache.Append(ctx, "categories", data).Err()
}
