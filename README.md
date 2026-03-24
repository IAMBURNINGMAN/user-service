# user-service

> gRPC microservice for user management. Part of a microservices system alongside [task-service](https://github.com/IAMBURNINGMAN/task-service).

## Stack

### Transport
- **gRPC / `google.golang.org/grpc`** — binary RPC over HTTP/2 with generated stubs from the shared [`proto`](https://github.com/IAMBURNINGMAN/proto) module. Contract versioning is delegated to the proto module; services consume it as a pinned dependency.
- **Protocol Buffers** — wire format ensuring type-safe, compact inter-service payloads without hand-written serialization.

### Persistence
- **[GORM v2](https://gorm.io)** — ORM with the `pgx/v5` PostgreSQL driver. Native binary protocol and connection pooling via `pgxpool`.
- **PostgreSQL** — primary datastore.

### Architecture

```
internal/user/
  model.go        # domain entity
  repository.go   # persistence interface + GORM implementation
  service.go      # business logic
internal/transport/grpc/
  handler.go      # gRPC request/response mapping
  server.go       # server registration and lifecycle
internal/database/
  db.go           # connection initialization
cmd/server/
  main.go         # dependency wiring
```

### Shared Contracts
- **[`github.com/IAMBURNINGMAN/proto`](https://github.com/IAMBURNINGMAN/proto)** — pinned at `v1.0.0`. Exposes `UserService` RPC methods: `CreateUser`, `GetUser`, `UpdateUser`, `DeleteUser`, `ListUsers`.

## Quick Start

```bash
go run cmd/server/main.go
```
