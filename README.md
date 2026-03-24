# ABCs

> REST API service built with a contract-first approach.

## Stack

### Transport
- **[Echo v4](https://github.com/labstack/echo)** — high-performance HTTP router with middleware chaining, request binding, and custom error handling. Interface-based handler design integrates cleanly with generated server code.

### Contract-First Code Generation
- **[oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)** — generates strongly-typed Echo server interfaces and request/response models directly from the OpenAPI 3.0 spec. The spec is the single source of truth; implementation drift is a compile-time error.
- **OpenAPI 3.0** — contract defined in `openapi/openapi.yaml`, split by resource tags (`tasks`, `users`) and compiled into separate packages via `make gen` / `make genuser`.

### Persistence
- **[GORM v2](https://gorm.io)** — ORM with chainable query builder, hooks, and association handling.
- **pgx/v5** — native PostgreSQL driver with binary protocol support and connection pooling via `pgxpool`. Used as the GORM driver for better performance over `lib/pq`.
- **[golang-migrate](https://github.com/golang-migrate/migrate)** — versioned SQL migrations with up/down scripts. Every DDL change is tracked in `migrations/` and fully reproducible.

### Testing
- **[testify](https://github.com/stretchr/testify)** — assertion library paired with generated repository mocks. Service layer is tested in isolation without a live database.

### Tooling
- **golangci-lint** — static analysis aggregator covering ~50 linters in a single pass.
- **Makefile** — unified task runner: `run`, `test`, `lint`, `migrate`, `gen*`.

## Project Layout

```
cmd/              # entrypoint
internal/
  TaskService/    # domain: model, repository, service, mock, tests
  UsersService/   # domain: model, repository, service
  handlers/       # Echo handler implementations
  web/            # oapi-codegen output (server interfaces, types)
  db/             # database initialization
migrations/       # versioned SQL files (up + down)
openapi/          # OpenAPI 3.0 spec
```

## Quick Start

```bash
make migrate      # apply DB migrations
make run          # start the server
make test         # run test suite
make lint         # static analysis
```
