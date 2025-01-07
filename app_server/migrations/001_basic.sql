-- Таблица пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    base_id UUID NOT NULL DEFAULT gen_random_uuid(),  
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_base_id UNIQUE(base_id)   
);

-- Предзаполнение таблицы пользователей
INSERT INTO users (base_id, name) VALUES
-- (gen_random_uuid(), 'Default'),
(gen_random_uuid(), 'John Doe'),
(gen_random_uuid(), 'Jane Smith'),
(gen_random_uuid(), 'Alice Johnson'),
(gen_random_uuid(), 'Bob Brown'),
(gen_random_uuid(), 'Charlie Davis');


-- Статусы
CREATE TYPE status AS ENUM ('draft', 'moderation', 'published');  

CREATE TABLE statuses (
    id SERIAL PRIMARY KEY,
    status status NOT NULL  
);

INSERT INTO statuses (status)
VALUES 
    ('draft'),
    ('moderation'),
    ('published');


-- Таблица категорий
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    creator UUID REFERENCES users(base_id) ON DELETE SET NULL 
);

-- Временная таблица для предзаполнения(чтоб автоинкремент не съехал)
CREATE TEMP TABLE temp_categories (
id INT,
name VARCHAR(255) NOT NULL,
parent_id INT
);

-- Данные предзаполнения категорий 
INSERT INTO temp_categories (id, name, parent_id) VALUES
(1, 'Электроника', 0),
(2, 'Мобильные телефоны', 1),
(3, 'Ноутбуки', 1),
(4, 'Одежда', 0),
(5, 'Мужская одежда', 4),
(6, 'Женская одежда', 4),
(7, 'Футболки', 5),
(8, 'Джинсы', 5),
(9, 'Обувь', 5),
(10, 'Услуги', 0),
(11, 'Услуги по уборке', 10),
(12, 'Услуги по саду', 10),
(13, 'Репетиторство', 10),
(14, 'Курсы английского языка', 13),
(15, 'Подготовка к ОГЭ', 13),
(16, 'Подготовка к ЕГЭ', 13),
(17, 'Сантехнические работы', 0),
(18, 'Математика ЕГЭ', 16);

-- Переносим данные по категориям из временной таблицы
INSERT INTO categories (name, parent_id)
SELECT name, parent_id
FROM temp_categories;

-- Удаляем временную таблицу
DROP TABLE temp_categories;



-- Функция для получения дефолтного статуса объявлений
CREATE OR REPLACE FUNCTION get_default_status() RETURNS INTEGER AS $$
BEGIN
    RETURN (SELECT id FROM statuses WHERE status = 'draft');
END;
$$ LANGUAGE plpgsql;


-- Объявления
CREATE TABLE advertisements (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(base_id) ON DELETE SET NULL, 
    status_id INTEGER NOT NULL DEFAULT get_default_status(),  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    last_upd TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    moderator UUID,  
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,  
    title VARCHAR(255) NOT NULL, 
    description TEXT NOT NULL,  
    price NUMERIC NOT NULL,  
    contact_info TEXT NOT NULL  
);

-- Временная таблица объявлений для предзаполнения
CREATE TEMP TABLE temp_advertisements (
    id SERIAL,
    user_id UUID,
    status_id INTEGER,
    created_at TIMESTAMP,
    last_upd TIMESTAMP,
    moderator UUID,
    category_id INTEGER,
    title VARCHAR(255),
    description TEXT,
    price NUMERIC,
    contact_info TEXT
);

-- Заполнение временной таблицы объявлений
INSERT INTO advertisements (user_id, status_id, category_id, title, description, price, contact_info) VALUES
((SELECT base_id FROM users WHERE name = 'John Doe'), 1, 1, 'Установка унитаза', 'Услуги сантехников для установки унитаза в квартире', 150.00, 'contact1@example.com'),
((SELECT base_id FROM users WHERE name = 'Jane Smith'), 2, 2, 'Ремонт мобильного телефона', 'Ремонт мобильных телефонов различных брендов', 80.00, 'contact2@example.com'),
((SELECT base_id FROM users WHERE name = 'Alice Johnson'), 3, 3, 'Продажа ноутбука', 'Ноутбук в хорошем состоянии, без дефектов', 500.00, 'contact3@example.com'),
((SELECT base_id FROM users WHERE name = 'Bob Brown'), 1, 4, 'Продажа мужской одежды', 'Продам мужскую куртку в отличном состоянии', 75.00, 'contact4@example.com'),
((SELECT base_id FROM users WHERE name = 'Charlie Davis'), 3, 5, 'Футболка', 'Продам футболку, новая, с бирками', 20.00, 'contact5@example.com'),
((SELECT base_id FROM users WHERE name = 'John Doe'), 2, 6, 'Джинсы', 'Продам джинсы, практически не носили, в отличном состоянии', 40.00, 'contact6@example.com'),
((SELECT base_id FROM users WHERE name = 'Jane Smith'), 3, 10, 'Услуги по уборке', 'Предлагаю услуги по уборке квартир и офисов', 100.00, 'contact7@example.com'),
((SELECT base_id FROM users WHERE name = 'Alice Johnson'), 1, 11, 'Репетиторство по математике', 'Предлагаю услуги репетитора по математике для школьников', 30.00, 'contact8@example.com'),
((SELECT base_id FROM users WHERE name = 'Bob Brown'), 2, 12, 'Услуги по саду', 'Предлагаю услуги по уходу за садом и огородом', 50.00, 'contact9@example.com'),
((SELECT base_id FROM users WHERE name = 'Charlie Davis'), 3, 13, 'Курсы английского языка', 'Предлагаю курсы английского языка для начинающих и продолжающих', 200.00, 'contact10@example.com');

-- Перенос данных объявлений из временной таблицы
INSERT INTO advertisements (user_id, status_id, created_at, last_upd, moderator, category_id, title, description, price, contact_info)
SELECT user_id, status_id, created_at, last_upd, moderator, category_id, title, description, price, contact_info
FROM temp_advertisements;

-- Удаляем временную таблицу
DROP TABLE temp_advertisements;


