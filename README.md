# Event Go Microservices

Микросервисная архитектура для управления событиями с аутентификацией и пользователями.

## Архитектура

### Сервисы

- **Auth Service** (порт 50052) - сервис аутентификации
  - JWT токены
  - Регистрация/логин
  - Обновление токенов
  - Redis для хранения refresh токенов

- **Users Service** (порт 50051) - сервис управления пользователями
  - CRUD операции с пользователями
  - PostgreSQL для хранения данных
  - GRPC API

### Инфраструктура

- **PostgreSQL** - база данных пользователей
- **Redis** - хранилище сессий и refresh токенов

## Запуск

### Требования

- Docker
- Docker Compose

### Быстрый старт

```bash
# Клонировать репозиторий
git clone <repository-url>
cd event-go

# Запустить все сервисы
docker-compose up --build

# Остановить сервисы
docker-compose down
```

### Порты

- Auth Service: `localhost:50052`
- Users Service: `localhost:50051`
- PostgreSQL: `localhost:5432`
- Redis: `localhost:6379`

## Разработка

### Структура проекта

```
event-go/
├── services/
│   ├── auth/          # Сервис аутентификации
│   └── users/         # Сервис пользователей
├── contracts/         # GRPC контракты
├── docker-compose.yml # Docker конфигурация
└── README.md
```

### Запуск в режиме разработки

```bash
# Запустить только базы данных
docker-compose up postgres redis -d

# Запустить auth сервис
cd services/auth
go run cmd/main.go

# Запустить users сервис
cd services/users
go run cmd/main.go
```

## API

### Auth Service

GRPC методы:
- `Login` - аутентификация пользователя
- `Refresh` - обновление токена
- `Logout` - выход из системы

### Users Service

GRPC методы:
- `CreateUser` - создание пользователя
- `GetUser` - получение пользователя
- `UpdateUser` - обновление пользователя
- `DeleteUser` - удаление пользователя

## Конфигурация

### Переменные окружения

#### Auth Service
- `REDIS_HOST` - хост Redis
- `REDIS_PORT` - порт Redis
- `USERS_SERVICE_ADDRESS` - адрес users сервиса
- `JWT_SECRET` - секрет для JWT
- `JWT_ACCESS_TOKEN_EXPIRATION` - время жизни access токена
- `JWT_REFRESH_TOKEN_EXPIRATION` - время жизни refresh токена

#### Users Service
- `POSTGRES_HOST` - хост PostgreSQL
- `POSTGRES_PORT` - порт PostgreSQL
- `POSTGRES_USER` - пользователь PostgreSQL
- `POSTGRES_PASSWORD` - пароль PostgreSQL
- `POSTGRES_DB` - имя базы данных
- `POSTGRES_SSLMODE` - режим SSL
