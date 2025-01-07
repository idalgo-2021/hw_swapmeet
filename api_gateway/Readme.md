# API Gateway Service

API Gateway служит шлюзом между клиентами и микросервисами, обеспечивая маршрутизацию, аутентификацию и обработку запросов.

## Основные маршруты

### **Auth-service**

Обрабатывает запросы, связанные с аутентификацией и управлением токенами.

- **POST /auth/token**  
  Генерация access и refresh токенов. 
  
  **Тело запроса:**
  ```json
   {
      "username": "string",
      "password": "string"
   }
   ```
   **Ответ:**
   ```json
   {
      "access_token": "string",
      "refresh_token": "string"
   }
   ```

- **POST /auth/validate**  
  Валидация access токена. 

  **Заголовки:**
  ```
  Authorization: Bearer <access_token>
   ```
   **Ответ:**
   ```json
   {
      "user_id": "string"
   }
   ```


- **POST /auth/refresh**  
  Обновление токенов. 

  **Тело запроса:**
  ```json
   {
      "refresh_token": "string"
   }
   ```
   **Ответ:**
   ```json
   {
      "access_token": "string",
      "refresh_token": "string"
   }
   ```



- **POST /auth/register**  
  Регистрация нового пользователя.

  **Тело запроса:**
  ```json
   {
      "username": "string",
      "password": "string",
      "email": "string"
   }  
   ```
   **Ответ:**
   ```json
   {
      "userId": "8e5fdaad-9720-4f15-b265-801f062abd85"
   } 
   ``` 


### **Swapmeet-service**
Обрабатывает запросы, связанные с категориями и объявлениями(к сервису SwapMeet).


- **GET /categories**

   Получение списка категорий объявлений.

   **Ответ:**
   ```json
   [
      {
         "id": "string",
         "name": "string",
         "parent_id": "string"
      }
   ]
   ```

- **POST /categories**

   Создание новой категории (требуется аутентификация).
   
   **Заголовки:**
   ```json
   Authorization: Bearer <access_token>
   ```

   **Тело запроса:**
   ```json
   {
      "name": "string",
      "parent_id": "string"
   }
   ```
  **Ответ:**
   ```json
   {
      "category": {
         "id": "string",
         "name": "string",
         "parent_id": "string"
      }
   }   
   ```


- **GET /advertisements**

   Получение опубликованных объявлений.

    **Заголовки:**
   * Токен доступа (опционально)

   ```json
   Authorization: Bearer <access_token>
   ```

   **Параметры запроса:**

   * cat (опционально) — ID категорий
   ```
   /advertisements?cat=1&cat=2....
   ```

   **Ответ:**
   ```json
   [
      {
         "category_id": "string",
         "category_name": "string",
         "contact_info": "string",
         "created_at": "string",
         "description": "string",
         "id": "string",
         "last_upd": "string",
         "price": "string",
         "status_id": "string",
         "status_name": "string",
         "title": "string",
         "user_id": "string",
         "user_name": "string"
      }
   ]
   ```


- **GET /advertisements/user**

   Получение объявлений, созданных аутентифицированным пользователем(определяется из данных JWT-токена).
   
   **Заголовки:**
   * Токен доступа

   ```json
   Authorization: Bearer <access_token>
   ```
   **Ответ:**
   ```json
   [
      {
         "category_id": "string",
         "category_name": "string",
         "contact_info": "string",
         "created_at": "string",
         "description": "string",
         "id": "string",
         "last_upd": "string",
         "price": "string",
         "status_id": "string",
         "status_name": "string",
         "title": "string",
         "user_id": "string",
         "user_name": "string"
      }
   ]
   ```

- **GET /advertisement/{id}**

   Получение опубликованного объявления по его ID.
   
   **Параметры запроса:**

   * id (обязательно) — ID объявления.

   **Ответ:**
   ```json
   {
      "category_id": "string",
      "category_name": "string",
      "contact_info": "string",
      "created_at": "string",
      "description": "string",
      "id": "string",
      "last_upd": "string",
      "price": "string",
      "status_id": "string",
      "status_name": "string",
      "title": "string",
      "user_id": "string",
      "user_name": "string"
   }
   ```

- **POST /advertisements**

   Создание нового объявления (требуется аутентификация).

   **Заголовки:** 

   ```json
   Authorization: Bearer <access_token>
   ```

   **Тело запроса:**
   ```json
   {
      "category_id": "string",
      "title": "string",
      "description": "string",
      "price": "number",
      "contact_info": "string"
   }
   ```

   **Ответ:**
   ```json
   {
      "id": "string", 
      "category_id": "string",
      "category_name": "string",
      "contact_info": "string",
      "created_at": "string",
      "description": "string",
      "last_upd": "string",
      "price": "string",
      "status_id": "string",
      "status_name": "string",
      "title": "string",
      "user_id": "string",
      "user_name": "string"
   }
   ```


- **PUT /advertisements**

   Изменение существующего объявления (требуется аутентификация). При любом изменении пользователем объявления, оно переводится в черновики. 
  
  **Заголовки:** 

   ```json
   Authorization: Bearer <access_token>
   ```
   **Тело запроса:**
    ```json
   {
      "advertisement_id": "string",
      "title": "string",
      "description": "string",
      "price": "number",
      "contact_info": "string"
   }
   ```
   **Ответ:**
   ```json
   {
      "category_id": "string",
      "category_name": "string",
      "contact_info": "string",
      "created_at": "string",
      "description": "string",
      "id": "string",
      "last_upd": "string",
      "price": "string",
      "status_id": "string",
      "status_name": "string",
      "title": "string",
      "user_id": "string",
      "user_name": "string"
   }
   ```


- **PUT /advertisement/moderation/{id}**

   Отправка пользователем объявления на модерацию(требуется аутентификация). 
  
  **Заголовки:** 

   ```json
   Authorization: Bearer <access_token>
   ```
  
   **Ответ:**
   ```json
   {
      "category_id": "string",
      "category_name": "string",
      "contact_info": "string",
      "created_at": "string",
      "description": "string",
      "id": "string",
      "last_upd": "string",
      "price": "string",
      "status_id": "string",
      "status_name": "string",
      "title": "string",
      "user_id": "string",
      "user_name": "string"
   }
   ```


## Основные маршруты

Для упрощения разработки и тестирования API используется Swagger-документация, которая автоматически генерируется с помощью утилиты ```swag```

### **Генерация документации**

Для генерации Swagger-файлов выполните следующую команду в корневой директории проекта:

swag init --parseDependency --dir ./cmd/main,./internal/handlers --output ./docs

### **Расположение файлов**

Сгенерированная документация будет размещена в директории ./docs. После генерации документации API можно просматривать через Swagger UI.

### **Доступ к Swagger UI**

Swagger UI доступен по маршруту /swagger/index.html через Gateway сервис, что предоставляет удобный интерфейс для тестирования и изучения API.