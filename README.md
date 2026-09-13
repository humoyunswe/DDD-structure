# DDD Structure Template

Шаблон Go-сервиса по той же DDD-слоевой архитектуре, что и рабочий проект `szpt_new`.

Полный пример модуля — **`users` (v1)**. Остальные папки модулей (organizations, documents, …) уже созданы — копируй паттерн `users`.

## Слои

```
cmd/server/main.go                 # composition root (wire)
configs/                           # env → Config + DSN
package/response/                  # единый HTTP response envelope

internal/
  domain/v{N}/{module}/            # entity, repository port, errors
  application/v{N}/{module}_use_case/  # usecase + ports alias
  infrastructure/persistence/
    models/v{N}/{module}/          # GORM models (DB tags)
    pg/v{N}/{module}/              # repository implementations
  interface/
    dto/v{N}/{module}/             # request/response DTO
    http/handlers/v{N}/{module}/   # Gin handlers
    http/routers/v{N}/{module}/    # route registration
    http/routers/dependencies.go   # DI bag
    http/routers/engineRouter/     # root router
    http/middleware/               # logging, auth
```

### Зависимости слоёв (строго)

```
handler → use case → domain ← repository (pg)
                ↑
              dto (только в handler)
```

- **domain** не знает про Gin / GORM / HTTP
- **application** оркестрирует бизнес-правила, зависит только от domain ports
- **infrastructure** реализует domain.Repository
- **interface** = транспорт (HTTP) + DTO

## Как добавить новый модуль

Пример: `organizations` v1.

1. Domain: `internal/domain/v1/organizations/{entity,repository,errors}.go`
2. Application: `internal/application/v1/organizations_use_case/{ports,usecase}.go`
3. Model: `internal/infrastructure/persistence/models/v1/organizations/`
4. Repo: `internal/infrastructure/persistence/pg/v1/organizations/`
5. DTO + Handler + Router в `internal/interface/...`
6. В `cmd/server/main.go` сварить: repo → service → handler
7. Добавить handler в `routers.Dependencies`
8. Зарегистрировать router в `engineRouter.NewRouter`

Смотри эталон: `internal/**/v1/users/**`.

## Запуск

```bash
cp .env.example .env
docker compose up -d postgres
# или весь стек:
docker compose up --build
```

Локально без Docker:

```bash
go mod tidy
go run ./cmd/server
```

Health: `GET /health`

API (нужен заголовок `Authorization: Bearer demo`):

```bash
curl -H "Authorization: Bearer demo" http://localhost:8080/api/private/v1/users
curl -H "Authorization: Bearer demo" -H "Content-Type: application/json" \
  -d '{"email":"carol@example.com","first_name":"Carol","last_name":"Example"}' \
  http://localhost:8080/api/private/v1/users
```

## Config

`configs/config.go` читает env так же, как `szpt_new`. Поменяй дефолты/поля под свой проект — имена переменных в `.env.example`.

В шаблоне auth — `DemoAuth` (любой Bearer). В проде замени на IdP userinfo, как `ValidationToken` в `szpt_new`.

## Модули (имена как в szpt_new)

| Version | Modules |
|---------|---------|
| v1 | users *(полный пример)*, organizations, documents, roles, references |
| v2 | users, organizations, documents, roles, references, approvals, contracts, balance, system |
| v3 | documents, references, warehouse |
| — | auth |

Пустые папки модулей содержат `README.md` со ссылкой на эталон.
