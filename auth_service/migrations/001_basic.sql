-- Активируем расширение pgcrypto, если еще не подключено
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Создание таблицы ролей
CREATE TABLE roles (
    id SERIAL PRIMARY KEY, -- INT с автоинкрементом
    name VARCHAR(50) UNIQUE NOT NULL, -- Название роли
    description TEXT -- Описание роли
);

-- Предзаполнение таблицы ролей
INSERT INTO roles (id, name, description) VALUES
    (1, 'user', 'Обычный пользователь'),
    (2, 'moderator', 'Модератор контента'),
    (3, 'admin', 'Администратор системы');

-- Таблица пользователей
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),   -- Уникальный идентификатор пользователя
    username VARCHAR(50) UNIQUE NOT NULL,           -- Имя пользователя
    email VARCHAR(255) UNIQUE,                      -- Электронная почта
    password_hash VARCHAR(255) NOT NULL,            -- Хэш пароля
    role_id INT NOT NULL DEFAULT 1 REFERENCES roles(id) ON DELETE SET NULL, -- Роль пользователя по умолчанию
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата регистрации
);

-- Предзаполнение таблицы пользователей
INSERT INTO users (id, username, email, password_hash, role_id, created_at) VALUES
    (gen_random_uuid(), 'Pasha', 'paul@example.com', '$2a$10$fUXpHWTa5OSy8aMDWBs2A.uxoZfVKUkNXD2bFk2KJ.i9OR2QHbDuu', 1, NOW()), --//pass=123
    (gen_random_uuid(), 'Serg', 'serg_moderator@example.com', '$2a$10$m0heb7ILz5mVpxlQZI1ZPO2pBsYlGXHnwui8aOuidByxa3a0vxfGG', 2, NOW()), --//pass=1234
    (gen_random_uuid(), 'Boris', 'boris_admin@example.com', '$2a$10$l1rgthxmp3Hs/CD1p7wxSu2CuF7IIZxGcJbtWjaeYI1tWpnXP8vhy', 3, NOW()); --//pass=12345


-- -- Создание таблицы отозванных токенов
-- CREATE TABLE revoked_tokens (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Уникальный идентификатор записи
--     token TEXT UNIQUE NOT NULL,                    -- Хэш отозванного токена
--     revoked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата и время отзыва
--     expires_at TIMESTAMP NOT NULL                  -- Дата истечения токена
-- );

-- -- Добавление индекса для ускорения поиска истекших токенов
-- CREATE INDEX idx_revoked_tokens_expires_at ON revoked_tokens (expires_at);



-- -- Очистка устаревших записей в таблице `revoked_tokens`
-- CREATE OR REPLACE FUNCTION clean_expired_tokens()
-- RETURNS VOID AS $$
-- BEGIN
--     DELETE FROM revoked_tokens WHERE expires_at < CURRENT_TIMESTAMP;
-- END;
-- $$ LANGUAGE plpgsql;
-- 
-- или для какого-то фонового процесса
-- DELETE FROM revoked_tokens WHERE expires_at < NOW();

