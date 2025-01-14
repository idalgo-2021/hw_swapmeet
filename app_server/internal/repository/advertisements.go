package repository

import (
	"app_server/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func (r *DBSwapmeetRepo) GetPublishedAdvertisements(ctx context.Context, categoryIDs []string) ([]models.UserAdvertisement, error) {

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

func (r *DBSwapmeetRepo) getPublishedAdvertisementsFromCache(ctx context.Context) ([]models.UserAdvertisement, error) {
	cachedAds, err := r.cache.Get(ctx, "publishedAds").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, models.ErrAdvertisementsNotFound
		}
		return nil, err
	}
	return unmarshalAdvertisements(cachedAds)
}

func (r *DBSwapmeetRepo) getPublishedAdvertisementsFromDB(ctx context.Context, categoryIDs []string) ([]models.UserAdvertisement, error) {

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
		WHERE s.status = 'published'
	`
	var args []interface{}

	if len(categoryIDs) > 0 {
		query += " AND a.category_id = ANY($1)"
		args = append(args, pq.Array(categoryIDs))
	}

	query += " ORDER BY a.id"

	err := r.db.Db.SelectContext(ctx, &advertisements, query, args...)
	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementsNotFound
		}
		return nil, models.ErrDBQuery
	}

	return advertisements, nil
}

////

func (r *DBSwapmeetRepo) setPublishedAdvertisementsToCache(ctx context.Context, advertisements []models.UserAdvertisement) error {
	data, err := marshalAdvertisements(advertisements)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, "publishedAds", data, 0).Err()
}

////

func (r *DBSwapmeetRepo) GetPublishedAdvertisementByID(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {

	var advertisement models.UserAdvertisement

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
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

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
		ORDER BY a.id
	`
	err := r.db.Db.SelectContext(ctx, &advertisements, query, userID)
	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

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
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		return nil, err
	}
	return &advertisement, nil
}

func (r *DBSwapmeetRepo) UpdateAdvertisement(ctx context.Context, userID, advertisementID, title, description, price, contactInfo string) (*models.UserAdvertisement, error) {

	// FIXME: Правильно сбрасывать статус в DRAFT только:
	// 	- если реквизиты действительно были изменены
	//	- если были изменены определенные реквизиты(например, при изменении цены или контактов не нужно возвращать обяъаление в черновики)

	var advertisement models.UserAdvertisement

	query := `
        UPDATE advertisements 
        SET title = $1, description = $2, price = $3, contact_info = $4, last_upd = NOW(), status_id = (
            SELECT id FROM statuses WHERE status = 'draft'
        )
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
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, models.ErrDBQuery
	}

	return &advertisement, nil
}

func (r *DBSwapmeetRepo) SetModerationStatusForAdvertisement(ctx context.Context, userID, advertisementID string) (*models.UserAdvertisement, error) {

	var advertisement models.UserAdvertisement

	query := `
        UPDATE advertisements 
        SET last_upd = NOW(), status_id = (
            SELECT id FROM statuses WHERE status = 'moderation'
        )
        WHERE id = $1 AND user_id = $2 
        RETURNING id, user_id, category_id, title, description, price, contact_info, status_id
    `
	err := r.db.Db.QueryRowContext(ctx, query, advertisementID, userID).Scan(
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
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, models.ErrDBQuery
	}

	return &advertisement, nil
}

////

func (r *DBSwapmeetRepo) GetAdvertisements(ctx context.Context, statuses, categoryIDs []string) ([]models.UserAdvertisement, error) {

	// FIXME: Использование кэша объявлений

	advertisements, err := r.getAdvertisementsFromDB(ctx, statuses, categoryIDs)
	if err != nil {
		return nil, err
	}

	// FIXME: Возможное обогащение кэша данными из БД

	return advertisements, nil
}

func (r *DBSwapmeetRepo) getAdvertisementsFromDB(ctx context.Context, statuses, categoryIDs []string) ([]models.UserAdvertisement, error) {

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
		WHERE 1 = 1
	`
	var args []interface{}

	argIndex := 1

	if len(statuses) > 0 {
		query += fmt.Sprintf(" AND s.status = ANY($%d)", argIndex)
		args = append(args, pq.Array(statuses))
		argIndex++
	}

	if len(categoryIDs) > 0 {
		query += fmt.Sprintf(" AND a.category_id = ANY($%d)", argIndex)
		args = append(args, pq.Array(categoryIDs))
		argIndex++
	}

	query += " ORDER BY a.id"

	err := r.db.Db.SelectContext(ctx, &advertisements, query, args...)
	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementsNotFound
		}
		return nil, models.ErrDBQuery
	}

	return advertisements, nil
}

////

func (r *DBSwapmeetRepo) PublishAdvertisement(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {

	tx, err := r.db.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Шаг 1: Блокировка записи и получение её текущего статуса
	var currentStatus string
	err = tx.QueryRowContext(ctx, `
        SELECT s.status 
        FROM advertisements a
        JOIN statuses s ON a.status_id = s.id
        WHERE a.id = $1 FOR UPDATE
    `, advertisementID).Scan(&currentStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, err
	}

	if currentStatus != "moderation" {
		return nil, models.ErrAdvertisementNotPublishable
	}

	// Шаг 2: Получение ID статуса "published"
	var publishedStatusID string
	err = tx.QueryRowContext(ctx, `
        SELECT id 
        FROM statuses 
        WHERE status = 'published'
    `).Scan(&publishedStatusID)
	if err != nil {
		return nil, err
	}

	// Шаг 3: Обновление статуса
	_, err = tx.ExecContext(ctx, `
        UPDATE advertisements 
        SET last_upd = NOW(), status_id = $1
        WHERE id = $2
    `, publishedStatusID, advertisementID)
	if err != nil {
		return nil, err
	}

	// Шаг 4: Получение информации о объявлении
	var advertisement models.UserAdvertisement
	err = tx.QueryRowContext(ctx, `
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
		WHERE a.id = $1;
	`, advertisementID).Scan(
		&advertisement.ID,
		&advertisement.UserID,
		&advertisement.UserName,
		&advertisement.StatusID,
		&advertisement.StatusName,
		&advertisement.CategoryID,
		&advertisement.CategoryName,
		&advertisement.CreatedAt,
		&advertisement.LastUpd,
		&advertisement.Title,
		&advertisement.Description,
		&advertisement.Price,
		&advertisement.ContactInfo,
	)

	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, models.ErrDBQuery
	}

	return &advertisement, nil
}

func (r *DBSwapmeetRepo) ReturnAdvertisementToDraft(ctx context.Context, advertisementID string) (*models.UserAdvertisement, error) {

	var advertisement models.UserAdvertisement

	tx, err := r.db.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Шаг 1: Блокировка записи и получение её текущего статуса
	var currentStatus string
	err = tx.QueryRowContext(ctx, `
        SELECT s.status 
        FROM advertisements a
        JOIN statuses s ON a.status_id = s.id
        WHERE a.id = $1
        FOR UPDATE
    `, advertisementID).Scan(&currentStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, err
	}

	// Если объявление уже в статусе draft, не обновляем его
	if currentStatus == "draft" {
		return nil, models.ErrAdvertisementAlreadyDraft
	}

	// Шаг 2: Получение ID статуса "draft"
	var draftStatusID string
	err = tx.QueryRowContext(ctx, `
        SELECT id 
        FROM statuses 
        WHERE status = 'draft'
    `).Scan(&draftStatusID)
	if err != nil {
		return nil, err
	}

	// Шаг 3: Обновление статуса на "draft"
	_, err = tx.ExecContext(ctx, `
        UPDATE advertisements 
        SET last_upd = NOW(), status_id = $1
        WHERE id = $2
    `, draftStatusID, advertisementID)
	if err != nil {
		return nil, err
	}

	// Шаг 4: Получение информации о объявлении после обновления
	err = tx.QueryRowContext(ctx, `
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
		WHERE a.id = $1;
	`, advertisementID).Scan(
		&advertisement.ID,
		&advertisement.UserID,
		&advertisement.UserName,
		&advertisement.StatusID,
		&advertisement.StatusName,
		&advertisement.CategoryID,
		&advertisement.CategoryName,
		&advertisement.CreatedAt,
		&advertisement.LastUpd,
		&advertisement.Title,
		&advertisement.Description,
		&advertisement.Price,
		&advertisement.ContactInfo,
	)

	if err != nil {
		r.logger.Info(ctx, fmt.Sprintf("DB query error: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrAdvertisementNotFound
		}
		return nil, models.ErrDBQuery
	}

	return &advertisement, nil
}
