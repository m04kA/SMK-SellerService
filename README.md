# SMK-SellerService

Микросервис для управления компаниями (автомойками) и их услугами в платформе онлайн-записи на автомойку.

## 🏗️ Архитектура

Проект построен на **Clean Architecture** с четким разделением слоёв:
- **Domain** - доменные модели (Company, Address, WorkingHours, Service)
- **Service** - бизнес-логика с авторизацией и DTO моделями
- **Repository** - работа с БД (PostgreSQL + lib/pq + squirrel)
- **Handlers** - HTTP API (handler per endpoint паттерн)
- **Middleware** - упрощённая аутентификация (MVP)
- **Logging** - многоуровневое логирование (INFO, WARN, ERROR)

## 🚀 Быстрый старт

### Вариант 1: Запуск в Docker (рекомендуется)

```bash
docker-compose up -d
```

Все сервисы запустятся автоматически:
- **PostgreSQL**: порт **5436** (чтобы избежать конфликтов с другими сервисами)
- **Миграции**: применяются автоматически через контейнер `migrate`

Для запуска приложения:
```bash
# В корне проекта
go run cmd/main.go
```

Приложение будет доступно на http://localhost:8081

### Вариант 2: Локальный запуск

1. Запустить PostgreSQL:
```bash
docker-compose up -d postgres
docker-compose up migrate
```

2. Запустить приложение:
```bash
go run cmd/main.go
```

Логи записываются в консоль и `logs/app.log` (WARN и ERROR)

### Тестирование API

#### Создание компании (требует X-User-ID и X-User-Role: superuser)
```bash
curl -X POST http://localhost:8081/api/v1/companies \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -H "X-User-Role: superuser" \
  -d '{
    "name": "Автомойка Премиум",
    "description": "Лучшая мойка в городе",
    "tags": ["#мойка", "#премиум", "#москва"],
    "logo_url": "https://example.com/logo.png",
    "manager_ids": [123456789],
    "addresses": [
      {
        "city": "Москва",
        "street": "Ленинский проспект",
        "building": "10",
        "latitude": 55.123456,
        "longitude": 37.123456
      }
    ],
    "working_hours": {
      "monday": "09:00-21:00",
      "tuesday": "09:00-21:00",
      "wednesday": "09:00-21:00",
      "thursday": "09:00-21:00",
      "friday": "09:00-21:00",
      "saturday": "10:00-20:00",
      "sunday": "10:00-20:00"
    }
  }'
```

#### Получение списка компаний (публичный endpoint)
```bash
curl -X GET 'http://localhost:8081/api/v1/companies?tags=#мойка,#москва&page=1&limit=10'
```

#### Получение компании по ID (публичный endpoint)
```bash
curl -X GET http://localhost:8081/api/v1/companies/1
```

#### Обновление компании (требует X-User-ID и X-User-Role)
```bash
curl -X PUT http://localhost:8081/api/v1/companies/1 \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 123456789" \
  -H "X-User-Role: user" \
  -d '{
    "name": "Автомойка Премиум Плюс",
    "description": "Обновленное описание"
  }'
```

#### Создание услуги (требует X-User-ID и X-User-Role)
```bash
curl -X POST http://localhost:8081/api/v1/companies/1/services \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 123456789" \
  -H "X-User-Role: user" \
  -d '{
    "name": "Комплексная мойка",
    "description": "Мойка кузова + салона",
    "duration_minutes": 60,
    "address_ids": [1]
  }'
```

#### Получение списка услуг компании (публичный endpoint)
```bash
curl -X GET http://localhost:8081/api/v1/companies/1/services
```

## 📋 API Endpoints

### Companies (Компании)

#### Public
- `GET /api/v1/companies` - список компаний с фильтрами (tags, page, limit)
- `GET /api/v1/companies/{id}` - получение компании по ID

#### Protected (требуют X-User-ID и X-User-Role)
- `POST /api/v1/companies` - создание компании (только superuser)
- `PUT /api/v1/companies/{id}` - обновление компании (superuser или manager компании)
- `DELETE /api/v1/companies/{id}` - удаление компании (только superuser)

### Services (Услуги)

#### Public
- `GET /api/v1/companies/{company_id}/services` - список услуг компании
- `GET /api/v1/companies/{company_id}/services/{service_id}` - получение услуги по ID

#### Protected (требуют X-User-ID и X-User-Role)
- `POST /api/v1/companies/{company_id}/services` - создание услуги (superuser или manager компании)
- `PUT /api/v1/companies/{company_id}/services/{service_id}` - обновление услуги (superuser или manager)
- `DELETE /api/v1/companies/{company_id}/services/{service_id}` - удаление услуги (superuser или manager)

## 🔧 Разработка

### Основные команды

```bash
# Сборка
go build -o bin/sellerservice cmd/main.go

# Запуск
go run cmd/main.go

# Тесты
go test ./...

# Установка зависимостей
go mod download
go mod tidy
```

### Docker команды

```bash
# Запуск всех сервисов
cd schemas && docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановка
docker-compose down

# Полная очистка (volumes + data)
docker-compose down -v
```

### Команды для работы с БД

```bash
# Применить миграции (автоматически через docker-compose)
cd schemas
docker-compose up -d postgres
docker-compose up migrate

# Откатить миграцию
docker-compose run --rm migrate -path /migrations \
  -database "postgres://postgres:postgres@postgres:5432/smk_sellerservice?sslmode=disable" \
  down

# Пересоздать БД
docker-compose down -v
docker-compose up -d postgres
sleep 5
docker-compose up migrate
```

## 📁 Структура проекта

```
SMK-SellerService/
├── cmd/
│   └── main.go                          # Entry point с DI и routing
├── internal/
│   ├── config/                          # Конфигурация (config.toml loader)
│   ├── domain/                          # Доменные модели (Company, Address, Service)
│   ├── service/                         # Бизнес-логика + DTOs + авторизация
│   │   ├── constants.go                # RoleSuperuser, RoleUser
│   │   ├── companies/                  # Сервис для компаний
│   │   └── services/                   # Сервис для услуг
│   ├── infra/storage/                   # Репозитории (PostgreSQL)
│   │   ├── company/                    # CRUD для компаний + связанные сущности
│   │   └── service/                    # CRUD для услуг
│   └── api/
│       ├── handlers/                    # HTTP handlers (handler per endpoint)
│       │   ├── utils.go                # RespondJSON, RespondError, DecodeJSON
│       │   ├── create_company/
│       │   ├── get_company/
│       │   ├── list_companies/
│       │   ├── update_company/
│       │   ├── delete_company/
│       │   ├── create_service/
│       │   ├── get_service/
│       │   ├── list_services/
│       │   ├── update_service/
│       │   └── delete_service/
│       └── middleware/
│           └── auth.go                 # UserIDAuth middleware
├── pkg/
│   ├── logger/                          # Injectable logger
│   └── psqlbuilder/                     # SQL query builder (squirrel wrapper)
├── migrations/                          # SQL миграции (golang-migrate)
├── schemas/
│   ├── schema.yaml                      # OpenAPI 3.0 спецификация
│   └── docker-compose.yml              # Docker окружение
├── config.toml                          # Конфигурация приложения
├── go.mod
├── CLAUDE.md                            # Детальная документация для AI
└── README.md
```

## ⚙️ Конфигурация

Файл `config.toml`:
```toml
[logs]
level = "info"

[server]
http_port = 8081

[database]
host = "localhost"
port = 5436
user = "postgres"
password = "postgres"
dbname = "smk_sellerservice"
sslmode = "disable"
```

### Особенности конфигурации

- **Порт PostgreSQL**: 5436 (не стандартный 5432) - чтобы избежать конфликтов с другими сервисами
- **Порт приложения**: 8081 (SMK-UserService использует 8080)
- Настройки БД из `config.toml` используются при локальном запуске
- При запуске в Docker можно переопределить через переменные окружения

## 🔐 Аутентификация и Авторизация

### Упрощенная аутентификация (MVP)

Для доступа к защищённым endpoints требуется передавать два заголовка:
```
X-User-ID: <user_id>
X-User-Role: <superuser|user>
```

⚠️ **Важно**: Это временное решение для MVP. В продакшене будет использоваться полноценная JWT аутентификация через отдельный SMK-AuthService.

### Роли и права доступа

Система поддерживает 2 роли:

#### 1. **superuser** (администратор системы)
- **Полный доступ** ко всем операциям
- Может создавать/удалять компании
- Может изменять любые компании и услуги
- Не требуется проверка на принадлежность к компании

#### 2. **user** (менеджер автомойки)
- Может изменять **только компании, где он указан в `manager_ids`**
- Может создавать/изменять/удалять услуги **только для своих компаний**
- При попытке доступа к чужой компании получает 403 Forbidden

### Примеры запросов с ролями

**Superuser создает компанию:**
```bash
curl -X POST http://localhost:8081/api/v1/companies \
  -H "X-User-ID: 1" \
  -H "X-User-Role: superuser" \
  -H "Content-Type: application/json" \
  -d '{"name": "Новая Автомойка", "manager_ids": [123456789]}'
```

**User (менеджер) обновляет свою компанию:**
```bash
curl -X PUT http://localhost:8081/api/v1/companies/1 \
  -H "X-User-ID: 123456789" \
  -H "X-User-Role: user" \
  -H "Content-Type: application/json" \
  -d '{"name": "Обновленное название"}'
```

**User пытается изменить чужую компанию (403 Forbidden):**
```bash
curl -X PUT http://localhost:8081/api/v1/companies/1 \
  -H "X-User-ID: 999999999" \
  -H "X-User-Role: user" \
  -H "Content-Type: application/json" \
  -d '{"name": "Попытка изменить"}'
# Ответ: {"code":403,"message":"access denied"}
```

**Superuser может удалить любую компанию:**
```bash
curl -X DELETE http://localhost:8081/api/v1/companies/1 \
  -H "X-User-ID: 1" \
  -H "X-User-Role: superuser"
```

**User не может удалять компании (только superuser):**
```bash
curl -X DELETE http://localhost:8081/api/v1/companies/1 \
  -H "X-User-ID: 123456789" \
  -H "X-User-Role: user"
# Ответ: {"code":403,"message":"access denied"}
```

## 🗄️ База данных

### Схема

Основные таблицы:
- **companies** - компании (автомойки) с массивами tags и manager_ids
- **addresses** - адреса компаний с геолокацией (many-to-one)
- **working_hours** - рабочие часы (one-to-one с companies)
- **services** - услуги компаний
- **service_addresses** - связь услуг с адресами (many-to-many)

### Ключевые особенности

- Все ID используют **BIGINT** (не UUID)
- Массивы хранятся как PostgreSQL массивы: `TEXT[]`, `BIGINT[]`
- Каскадное удаление через `ON DELETE CASCADE`
- Поддержка геолокации через `latitude` и `longitude` (NUMERIC)

### Миграции

Миграции находятся в `migrations/`:
- `000001_init_schema.up.sql` - создание всех таблиц
- `000001_init_schema.down.sql` - откат миграции

Применяются автоматически при запуске `docker-compose up`

## 📊 Логирование

Логирование через injectable logger (`pkg/logger`):

Все логи пишутся в консоль. Логи уровня **WARN** и **ERROR** дополнительно сохраняются в `logs/app.log`.

### Уровни логирования

- **INFO** - успешные операции (создание, обновление, получение)
- **WARN** - ошибки валидации и авторизации (4xx)
- **ERROR** - внутренние ошибки сервера (5xx)

### Примеры логов

```
[INFO] POST /companies - Created: company_id=1, user_id=123456789
[WARN] POST /companies - Access denied: user_id=999999999, role=user
[ERROR] POST /companies - Failed: user_id=123456789, error=...
```

## 🔗 Интеграция с другими сервисами

SMK-SellerService является частью микросервисной архитектуры **SMK (Smart Mobile Karwash)**:

### SMK-UserService (порт 8080)
- Управление пользователями и их автомобилями
- Предоставляет информацию о пользователях через endpoint `/internal/users/{tg_user_id}`
- Использует аналогичную архитектуру и паттерны

### Соответствие ролей

| SellerService | UserService | Описание |
|---------------|-------------|----------|
| `user` | `manager` | Менеджер автомойки |
| `superuser` | `superuser` | Администратор системы |
| - | `client` | Клиент (только в UserService) |

### SMK-PriceService (планируется)
- Управление ценами на услуги в зависимости от размера автомобиля
- Интеграция с SellerService для получения списка услуг

### Будущая архитектура

```
Telegram Bot
    ↓
SMK-AuthService (JWT генерация, валидация InitData)
    ↓
API Gateway (валидация JWT, маршрутизация)
    ↓
├── SMK-UserService (пользователи, автомобили) - :8080
├── SMK-SellerService (компании, услуги) - :8081
└── SMK-PriceService (цены) - :8082
```

## 🧪 Тестирование

### Запуск тестов

```bash
go test ./...
```

### Тестирование API через curl

Примеры запросов смотрите в секции [Тестирование API](#тестирование-api)

## 🛠️ Технический стек

- **Go**: 1.24.2
- **PostgreSQL**: 16-alpine
- **HTTP Router**: Gorilla Mux
- **Query Builder**: Squirrel (через psqlbuilder wrapper)
- **PostgreSQL Driver**: lib/pq
- **Миграции**: golang-migrate
- **Контейнеризация**: Docker + Docker Compose

## 📚 Документация

### Для разработчиков

Полная документация архитектуры и паттернов проектирования находится в [CLAUDE.md](CLAUDE.md):
- Детальное описание архитектурных паттернов
- Примеры кода для всех слоёв
- Best Practices и антипаттерны
- Обработка ошибок (трёхуровневая система)
- Dependency Injection паттерн
- Troubleshooting

### API спецификация

API Contract описан в [schemas/schema.yaml](schemas/schema.yaml) (OpenAPI 3.0).

## 🤝 Вклад в проект

При разработке новых функций придерживайтесь паттернов, описанных в [CLAUDE.md](CLAUDE.md):
- Clean Architecture с разделением слоёв
- Handler per Endpoint паттерн
- Трёхуровневая система обработки ошибок (Repository → Service → Handler)
- Injectable Logger через интерфейсы
- Явный rollback в транзакциях (не через defer)
- Использование squirrel.Eq для WHERE условий

## 📝 Лицензия

Проект разработан для платформы онлайн-записи на автомойку SMK (Smart Mobile Karwash).

## 🆘 Поддержка

При возникновении проблем:
1. Проверьте секцию [Troubleshooting в CLAUDE.md](CLAUDE.md#troubleshooting)
2. Убедитесь, что используете правильные порты (8081 для app, 5436 для PostgreSQL)
3. Проверьте логи в `logs/app.log`
4. Проверьте логи Docker: `docker-compose logs -f`
