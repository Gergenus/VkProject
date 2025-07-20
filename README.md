# Установка 
1. Клонируйте репозиторий
    ```bash
    git clone https://github.com/Gergenus/VkProject.git
    ```
    
2. Создайте внутри проекта .env файл, готовая конфигурация:
  - **POSTGRES_URL=postgresql://admin:12345@app_db:5432/app_db?sslmode=disable**
  - **LOG_LEVEL=debug**
  - **HTTP_PORT=8080**
  - **JWT_SECRET=a-string-secret-at-least-256-bits-long**
  - **LOG_TYPE=text**
  - **TOKEN_TTL=24h**

3. Поднимите проект
    ```bash
    docker compose up --build
    ```
4. Примените миграции
   ```bash
   goose -dir ./internal/migrations/ postgres "postgresql://admin:12345@localhost:5433/app_db?sslmode=disable" up
   ```
# Взаимодйествие с API
### 1. Регистрация
   `POST http://localhost:8080/auth/signUp`
#### Разумные ограничения:
  - password 8->50 символов
  - login 3->50 симоволов
##### Тело запроса:
```
{
    "login": "login",
    "password": "password"
}
```
##### Тело ответа:
```
{
    "login": "login",
    "uuid": "9c9e1eeb-5fb3-49a8-abb6-6af0ee4f2e66"
}
```
### 2. Авторизация
   `POST http://localhost:8080/auth/signIn`
##### Тело запроса:
```
{
    "login": "login",
    "password": "password"
}
```
##### Тело ответа:
```
{
    "token": "jwt_token"
}
```
### 3. Создание объявления
`POST http://localhost:8080/post/create`
#### Разумные ограничения:
  - subject 0->100 символов
  - post_text 0->2500 симоволов
  - image_address size <= 5mb
  - allowedPhotoTypes: jpg, png
  ##### Тело запроса:
```
{
    "subject": "subject",
    "post_text": "post_text",
    "image_address": "https://www.example.ru/upload/photo.jpg",
    "price": price
}
```
##### Тело ответа:
```
{
    "id": id,
    "subject": "subject"
}
```
### 4. Отображение ленты сообщений
`GET http://localhost:8080/posts`
#### Query параметы:
- page - int (страница) (обязательный)
- page_size - int (количество объявлений на странице) (обязательный)
- sort_by - (price, created_at) (тип сортировки, доступны сортировки по времени добавления и цене) (опциональный)
- sort_dir - (asc, desc) (тип сортировки, доступны сортировки по возрастанию и убыванию) (опциональный)
- min_price - float (минимальная цена) (опциональный)
- max_price - float (максимальная цена) (опциональный)
#### При отстуствии опциональных параметров выводятся самые свежие объявления
##### Тело ответа:
```
[
    {
        "id": id,
        "login": "login",
        "subject": "subject",
        "post_text": "post_text",
        "image_address": "https://www.example.ru/upload/photo.jpg",
        "price": price,
        "created_at": "created_at",
        "is_owner": is_owner
    },
]
```
