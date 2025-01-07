package repository

import (
	"app_server/internal/models"
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func (r *DBSwapmeetRepo) GetPublishedAdvertisements(ctx context.Context, categoryIDs []string) ([]models.PublishedAdvertisement, error) {

	// FIXME: Переработать механизмы работы с кэшем объявлений
	if len(categoryIDs) <= 0 {
		publishedAds, err := r.getPublishedAdvertisementsFromCache(ctx)
		if err == nil {
			return publishedAds, nil
		}
	}

	publishedAds, err := r.getPublishedAdvertisementsFromDB(ctx, categoryIDs)
	if err != nil {
		return nil, err
	}

	// FIXME: Возможное обогащение кэша данными из БД

	return publishedAds, nil
}

func (r *DBSwapmeetRepo) getPublishedAdvertisementsFromCache(ctx context.Context) ([]models.PublishedAdvertisement, error) {
	cachedAds, err := r.cache.Get(ctx, "publishedAds").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, models.ErrAdvertisementsNotFound
		}
		return nil, err
	}
	return unmarshalAdvertisements(cachedAds)
}

func (r *DBSwapmeetRepo) getPublishedAdvertisementsFromDB(ctx context.Context, categoryIDs []string) ([]models.PublishedAdvertisement, error) {

	var advertisements []models.PublishedAdvertisement

	query := `
		SELECT 
   			a.id,  
			a.user_id,
   			u.name AS user_name, 
			a.status_id, 
   			s.status AS status_name,
			a.category_id,
			c.name AS category_name,  
   			TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
   			TO_CHAR(a.last_upd, 'YYYY-MM-DD HH24:MI:SS') AS last_upd,
   			a.title, 
   			a.description, 
   			TO_CHAR(a.price, 'FM9999999.99') AS price,
   			a.contact_info 
		FROM advertisements a
		JOIN users u ON a.user_id = u.base_id
		JOIN statuses s ON a.status_id = s.id
		JOIN categories c ON a.category_id = c.id
		WHERE s.status = 'published'
	`
	var args []interface{}
	if len(categoryIDs) > 0 {
		query += " AND a.category_id = ANY($1)"
		args = append(args, pq.Array(categoryIDs))
	}

	err := r.db.Db.SelectContext(ctx, &advertisements, query, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementsNotFound
		}
		return nil, models.ErrDBQuery
	}
	return advertisements, nil
}

////

func (r *DBSwapmeetRepo) setPublishedAdvertisementsToCache(ctx context.Context, advertisements []models.PublishedAdvertisement) error {
	data, err := marshalAdvertisements(advertisements)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, "publishedAds", data, 0).Err()
}

////

func (r *DBSwapmeetRepo) GetPublishedAdvertisementByID(ctx context.Context, advertisementID string) (*models.PublishedAdvertisement, error) {

	var advertisement models.PublishedAdvertisement

	query := `
		SELECT 
   			a.id,  
			a.user_id,
   			u.name AS user_name, 
			a.status_id, 
   			s.status AS status_name,
			a.category_id, 
   			c.name AS category_name,  
   			TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
   			TO_CHAR(a.last_upd, 'YYYY-MM-DD HH24:MI:SS') AS last_upd,
   			a.title, 
   			a.description, 
   			TO_CHAR(a.price, 'FM9999999.99') AS price,
   			a.contact_info 
		FROM advertisements a
		JOIN users u ON a.user_id = u.base_id
		JOIN statuses s ON a.status_id = s.id
		JOIN categories c ON a.category_id = c.id
		WHERE s.status = 'published' AND a.id = $1
	`
	err := r.db.Db.QueryRowContext(ctx, query, advertisementID).Scan(
		&advertisement.ID, &advertisement.UserID, &advertisement.UserName,
		&advertisement.StatusID, &advertisement.StatusName,
		&advertisement.CategoryID, &advertisement.CategoryName,
		&advertisement.CreatedAt, &advertisement.LastUpd, &advertisement.Title,
		&advertisement.Description, &advertisement.Price, &advertisement.ContactInfo,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, err
	}
	return &advertisement, nil
}

////

func (r *DBSwapmeetRepo) GetUserAdvertisements(ctx context.Context, userID string) ([]models.UserAdvertisement, error) {

	// FIXME: Добавить механизм использования кэша

	userAds, err := r.getUserAdvertisementsFromDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	// FIXME: Возможное обогащение кэша данными из БД

	return userAds, nil
}

func (r *DBSwapmeetRepo) getUserAdvertisementsFromDB(ctx context.Context, userID string) ([]models.UserAdvertisement, error) {

	var advertisements []models.UserAdvertisement

	query := `
		SELECT 
   			a.id,  
			a.user_id,
   			u.name AS user_name, 
			a.status_id, 
   			s.status AS status_name,
			a.category_id, 
   			c.name AS category_name,  
   			TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
   			TO_CHAR(a.last_upd, 'YYYY-MM-DD HH24:MI:SS') AS last_upd,
   			a.title, 
   			a.description, 
   			TO_CHAR(a.price, 'FM9999999.99') AS price,
   			a.contact_info 
		FROM advertisements a
		JOIN users u ON a.user_id = u.base_id
		JOIN statuses s ON a.status_id = s.id
		JOIN categories c ON a.category_id = c.id
		WHERE a.user_id = $1
	`
	err := r.db.Db.SelectContext(ctx, &advertisements, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementsNotFound
		}
		return nil, models.ErrDBQuery
	}
	return advertisements, nil
}

////

func (r *DBSwapmeetRepo) CreateAdvertisement(ctx context.Context, userID, categoryID, title, description, price, contactInfo string) (*models.UserAdvertisement, error) {

	var advertisement models.UserAdvertisement

	query := `
		INSERT INTO advertisements (user_id, category_id, title, description, price, contact_info, status_id) 
		VALUES ($1, $2, $3, $4, $5, $6, (SELECT id FROM statuses WHERE status = 'draft')) 
		RETURNING id, user_id, category_id, title, description, price, contact_info, status_id
    `
	err := r.db.Db.QueryRowContext(ctx, query, userID, categoryID, title, description, price, contactInfo).Scan(
		&advertisement.ID,
		&advertisement.UserID,
		&advertisement.CategoryID,
		&advertisement.Title,
		&advertisement.Description,
		&advertisement.Price,
		&advertisement.ContactInfo,
		&advertisement.StatusID,
	)
	if err != nil {
		return nil, err
	}
	return &advertisement, nil
}

func (r *DBSwapmeetRepo) UpdateAdvertisement(ctx context.Context, userID, advertisementID, title, description, price, contactInfo string) (*models.UserAdvertisement, error) {

	var advertisement models.UserAdvertisement

	query := `
        UPDATE advertisements 
        SET title = $1, description = $2, price = $3, contact_info = $4, last_upd = NOW() 
        WHERE id = $5 AND user_id = $6 
        RETURNING id, user_id, category_id, title, description, price, contact_info, status_id
    `
	err := r.db.Db.QueryRowContext(ctx, query, title, description, price, contactInfo, advertisementID, userID).Scan(
		&advertisement.ID,
		&advertisement.UserID,
		&advertisement.CategoryID,
		&advertisement.Title,
		&advertisement.Description,
		&advertisement.Price,
		&advertisement.ContactInfo,
		&advertisement.StatusID,
	)
	if err != nil {
		return nil, err
	}
	return &advertisement, nil
}
