# Auth Service (Тестовое задание для mdeost)

Простой сервис аутентификации на Go.

## Что реализовано

- Выдача access и refresh токенов
- Access токен — JWT, подписан алгоритмом HS512 (SHA512), (Не сохраняется в Базе данных)
- Refresh токен — случайная строка, передаётся в base64, хранится в виде bcrypt-хеша
- Refresh токен связан с access токеном через jti
- При попытке использовать refresh токен второй раз — показывается ошибка
- Токены привязаны к IP клиента, при изменении IP можно логировать или отправлять уведомление
- Все данные refresh токенов хранятся в PostgreSQL
  
## Запуск

1. Установите переменные окружения:

export DATABASE_URL=postgres://auth:secret@localhost:5432/auth?sslmode=disable
export JWT_SECRET=supersecret

2. Запустите скрипт
   
go build -o authsvc
./authsvc

