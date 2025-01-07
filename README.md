# SwapMeet - доска объявлений.

*Сервис является более детальной проработкой моего pet-проекта: https://github.com/idalgo-2021/ya_lms_swapmeet, выполненного в рамках ограниченного времени. Это своего рода работа над ошибками.*

## Описание изменений

* Во всех компонентах(микросервисах) изменена работа с логером(теперь он создается на уровне конструктора соответствующего модуля), упрощены запросы к БД, частично отрефакторен код, устранены найденные ошибки.
* Добавлен функционал отправки объявлений на модерацию и собственно функционал модерации. 
* При изменении пользователем объявления, оно (вне зависимости от текущего статуса) переводится в черновики.
* Упрощен и расширен контракт взаимодействия с Swapmeet-сервером(теперь токены доступа передаются в метаданных запросов).
* Обновлена и расширена документация Swagger(на шлюзе).

## Краткое описание

### Сервис предусматривает три авторизованных роли пользователей
*    `user` - 'Обычный пользователь' (пока при регистрации все создаются с такой ролью - в БД значение по умолчанию)
*    `moderator` - 'Модератор контента'
*    `admin` - 'Администратор системы'

### Сервис предусматривает три статуса объявлений
* `draft` - Черновик(новое, возвращенное с модерации, снятое с публикации)
* `moderation` На модерации
* `published` - Опубликовано

### Описание функциональности
* Неавторизованные пользователи
    - Просмотр всех опубликованных объявлений
    - Просмотр категорий объявлений
    - Просмотр объявления по их идентификатору(любых)

* Обычный пользователь(авторизованный)
    - Просмотр всех опубликованных объявлений
    - Просмотр категорий объявлений
    - Просмотр объявления по их идентификатору(любых)
    - Создание новых объявлений
    - Редактирование собственных объявлений
    - Просмотр всех своих объявлений

* Модератор и администратор
     * у пользователя во вх.параметрах HTTP запроса валидный токен доступа с ролью moderator или admin(проверили в auth-service)
Доступно
    - Просмотр всех опубликованных объявлений
    - Просмотр категорий объявлений
    - Просмотр объявления по их идентификатору(любых)
    - Создание новых объявлений
    - Редактирование собственных объявлений
    - Просмотр всех своих объявлений
    - Создание новых категорий

---

Проект реализует сервис доски объявлений и состоит из трёх микросервисов: 

## 1.  Сервис авторизации 

Предоставляет функционал авторизации пользователей на основе JWT-токенов(использует отдельную БД PostgreSQL). 

* [Каталог сервиса авторизации](auth_service)
* [Описание сервиса авторизации](auth_service/Readme.md)


## 2. Шлюз(API Gateway) 

Предоставляет HTTP API потребителям и обеспечивает маршрутизацию пользовательских запросов между сервисами.

* [Каталог сервиса API Gateway](api_gateway)
* [Описание сервиса API Gateway](api_gateway/Readme.md)


## 3. Сервер Swapmeet

Реализует базовую логику работы с оъявлениями(использует отдельные БД PostgreSQL и Redis).

* [Каталог сервиса Swapmeet](app_server)
* [Описание сервиса Swapmeet](app_server/Readme.md)

---


Взаимодействие микросервисов основано на gRPC вызовах, контракты которых(протофайлы) находятся в папках `proto`(каждого из сервисов). 

## Архитектура SwapMeet


* Архитектурный скетч реализованного сервиса(прототипа):

![Архитектура прототипа](docs/Scetch_MvpArch.jpg)



* Архитектурный скетч целевого сервиса:

![Целевая архитектура](docs/Scetch_PurposeArch.jpg)


## Подготовка и запуск проекта
* Скачайте локально репозиторий проекта(сервиса)
```
git clone https://gitlab.crja72.ru/gospec/go12/swapmeet.git
```

* Установите локально Golang(последней версии)
* Установите локально Docker

### Покомпонентный запуск
* В каталоге каждого из компонентов(микросервисов) нужно выполнить команды:
    * Установка зависимостей: 
    ```
    go mod tidy
    ``` 
    * Сборка контейнеров с БД: 
    ```
    sudo docker compose up --build
    ```
    * Запуск сикросервиса: 
    ```
    go run main.go
    ```