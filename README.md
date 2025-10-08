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
# Использование Makefile
make docker-up

# Или напрямую через docker-compose
docker-compose up -d
```

Все сервисы запустятся автоматически:
- **PostgreSQL**: порт **5436** (чтобы избежать конфликтов с другими сервисами)
- **Миграции**: применяются автоматически через контейнер `migrate`
- **Приложение**: доступно на http://localhost:8081
- **Метрики**: доступны на http://localhost:8081/metrics

### Вариант 2: Локальный запуск с Docker БД

```bash
# Запустить только базу данных
make dev

# Запустить приложение локально
make run
```

### Вариант 3: Полностью локальный запуск

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

### Makefile команды

Проект использует Makefile для упрощения работы. Все доступные команды:

```bash
# Справка
make help

# Базовые команды
make build          # Собрать бинарник в bin/smk-sellerservice
make run            # Запустить приложение локально
make test           # Запустить тесты
make clean          # Очистить артефакты сборки и логи
make install        # Установить Go зависимости

# Docker команды
make docker-build   # Собрать Docker образ
make docker-up      # Запустить все сервисы (postgres + migrate + app)
make docker-down    # Остановить все сервисы
make docker-restart # Перезапустить сервисы
make docker-logs    # Показать логи всех контейнеров
make docker-logs-app # Показать логи только приложения
make docker-clean   # Остановить и удалить volumes
make docker-prune   # Удалить Docker образы проекта

# База данных
make migrate-up     # Применить миграции
make migrate-down   # Откатить миграции
make db-reset       # Сбросить БД (удалить volumes + поднять заново)

# Разработка
make dev            # Запустить только БД для локальной разработки
```

### Типичные сценарии разработки

```bash
# Первый запуск проекта
make install
make docker-up

# Разработка с локальным запуском приложения
make dev          # Поднять только БД
make run          # Запустить приложение локально

# Полный Docker запуск (приложение + БД)
make docker-up

# Сброс базы данных
make db-reset

# Полная очистка проекта
make clean-all
```

### Основные команды без Makefile

```bash
# Сборка
go build -o bin/smk-sellerservice cmd/main.go

# Запуск
go run cmd/main.go

# Тесты
go test ./...

# Установка зависимостей
go mod download
go mod tidy
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

## 📈 Мониторинг и метрики

### Prometheus метрики

Сервис экспортирует метрики в формате Prometheus на endpoint `/metrics`:

```bash
curl http://localhost:8081/metrics
```

### Доступные метрики

**HTTP метрики:**
- `http_requests_total` - общее количество HTTP запросов (с лейблами: method, endpoint, status_code, service)
- `http_request_duration_seconds` - гистограмма длительности HTTP запросов

**Database метрики:**
- `db_queries_total` - общее количество SQL запросов (с лейблами: operation, table, status, service)
- `db_query_duration_seconds` - гистограмма длительности SQL запросов
- `db_connections_active` - количество активных соединений с БД
- `db_connections_idle` - количество простаивающих соединений в пуле
- `db_connections_max` - максимальное количество соединений

### Централизованный мониторинг

Метрики автоматически собираются централизованным сервисом **[SMK-Monitoring](https://github.com/m04kA/SMK-Monitoring)**:

- **Prometheus** - сбор и хранение метрик (http://localhost:9090)
- **Grafana** - визуализация метрик (http://localhost:3000, admin/admin)
- **PostgreSQL Exporter** - метрики базы данных

### Дашборд в Grafana

Готовый дашборд "SMK-SellerService Metrics" включает:
- График HTTP запросов по endpoint и status code
- График длительности HTTP запросов (p50, p95, p99)
- График SQL запросов по операциям (SELECT, INSERT, UPDATE, DELETE, транзакции)
- График длительности SQL запросов
- Состояние connection pool (active, idle, max connections)
- Использование памяти и CPU (через PostgreSQL Exporter)

### Конфигурация метрик

Метрики настраиваются через `config.toml` или переменные окружения:

```toml
[metrics]
enabled = true
path = "/metrics"
service_name = "sellerservice"
```

Переменные окружения:
```bash
METRICS_ENABLED=true
METRICS_PATH=/metrics
METRICS_SERVICE_NAME=sellerservice
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
