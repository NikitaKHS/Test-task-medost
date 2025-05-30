# Auth Service (Тестовое задание для medost)

Небольшой сервис для авторизации пользователей, реализованный на Go.

## Что сделано:

Генерация пары токенов (access и refresh).

Access-токен: JWT-токен с SHA512-подписью, не хранится в БД.

Refresh-токен: случайно сгенерированная строка в формате base64, хранится в БД в виде bcrypt-хеша.

Оба токена связаны через поле jti (JWT ID), что позволяет проверять их валидность и предотвращать повторное использование одного refresh-токена.

Привязка токенов к IP-адресу клиента. Если при запросе нового токена IP изменился, система уведомляет об этом (лог)

Данные refresh-токенов сохраняются в PostgreSQL.
  
## Запуск

1. Установите переменные окружения:

export DATABASE_URL=postgres://auth:secret@localhost:5432/auth?sslmode=disable
export JWT_SECRET=supersecret

2. Затем соберите и запустите приложение:
   
go build -o authsvc
./authsvc
