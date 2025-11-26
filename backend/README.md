# Ticketing System - Backend API

REST API for live events ticketing, built with Go and clean architecture principles.

## Project Structure

```
cmd/api/          - Main application entry point
internal/
  ├── config/     - Environment configuration
  ├── dto/        - Request/response structures
  ├── handlers/   - HTTP handlers
  ├── middleware/ - Auth, logging, CORS
  ├── models/     - Domain entities
  ├── repositories/ - Database layer
  ├── router/     - Route definitions (per-domain)
  └── services/   - Business logic
pkg/              - Shared utilities
migrations/       - Database migrations
```

## Architecture Decisions

### Clean Architecture (Layered Approach)

We follow a **layered architecture** pattern to keep concerns separated:

```
Request → Router → Middleware → Handler → Service → Repository → Database
```

**Why this approach?**
- **Testability**: Each layer can be tested independently
- **Maintainability**: Changes in one layer don't cascade to others
- **Team scalability**: Different developers can work on different layers without conflicts

The flow is simple: handlers stay thin (just parse requests), services contain business logic, and repositories handle data access. This makes the codebase easier to understand and modify.

### Modular Router Structure

Routes are organized **per-domain** in separate files:

```
internal/router/
  ├── router.go         # Main setup
  ├── auth_routes.go    # Authentication endpoints
  ├── event_routes.go   # Event management
  ├── ticket_routes.go  # Ticket operations
  └── user_routes.go    # User management
```

**Why split routes by domain?**
- **Concurrent development**: Multiple developers can work on different route files without merge conflicts
- **Easier navigation**: Finding event-related routes? Check `event_routes.go`
- **Scale-friendly**: As the API grows, each domain file stays manageable
- **Clear ownership**: Team members can own specific domains

This becomes critical when you have 5-10 developers working on the codebase simultaneously.

### Dependency Injection Pattern

Dependencies are wired in `main.go` using a simple struct-based approach:

```go
repos → services → handlers → router
```

This makes it easy to swap implementations (e.g., switch from Postgres to MySQL) or inject mocks for testing.

## Tech Stack

- **Go 1.21+** - Fast compilation, great concurrency
- **Gin** - Lightweight HTTP framework
- **PostgreSQL** - ACID compliance for ticket transactions
- **Redis** - Caching layer (upcoming)
- **sqlx** - Thin wrapper over database/sql
- **zerolog** - Structured logging
- **JWT** - Stateless authentication

Full stack rationale in [`docs/ARCHITECTURE.md`](../docs/ARCHITECTURE.md)

## Quick Start

### With Docker Compose (recommended)

```bash
docker-compose up -d
```

Everything (PostgreSQL, Redis, MinIO, Backend API) starts with one command.

### Local Development

```bash
# Copy environment template
cp .env.example .env

# Install dependencies
make deps

# Run server
make run
```

Make sure PostgreSQL, Redis, and MinIO are running locally (use docker-compose or run them separately).

## API Endpoints

**Public routes:**
- `POST /api/auth/login`
- `POST /api/auth/register`
- `GET /api/events`
- `GET /api/events/:id`

**Protected routes (requires JWT):**
- `POST /api/tickets/purchase`
- `GET /api/tickets/my-orders`
- `GET /api/users/me`

**Admin-only routes:**
- `POST /api/events` - Create event
- `DELETE /api/users/:id` - Delete user

See [`docs/ARCHITECTURE.md`](../docs/ARCHITECTURE.md) for full API specification and flow diagrams.

## Development Commands

```bash
make help          # List all commands
make run           # Run the server
make build         # Build binary
make test          # Run tests
make migrate-up    # Apply database migrations
make mocks         # Generate mocks for testing
```

## Testing

### Unit Tests

Tests follow table-driven approach with generated mocks:

```bash
# Run all tests
make test

# Run specific package
go test -v ./internal/services/

# Generate mocks (mockery required)
make mocks
```

Test files are located alongside source code (e.g., `auth_service_test.go` next to `auth_service.go`). Mocks are in `internal/repositories/mocks/`.

### HTTP/API Testing

Use `.http` files in `api/` directory for manual endpoint testing. Works with:
- **VSCode**: Install [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension
- **JetBrains IDEs**: Built-in HTTP Client (IntelliJ, GoLand, etc.)

Example workflow:
1. Start the server (`make run`)
2. Open `api/auth.http`
3. Click "Send Request" to execute
4. Copy the JWT token from response
5. Use token in protected endpoints (already configured in files)

Files:
- `api/auth.http` - Login, register
- `api/events.http` - List events
- `api/tickets.http` - Purchase tickets, view orders
- `api/users.http` - User management

## Configuration

Environment variables are loaded from `.env` file. Required settings:

- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Token signing key
- `PORT` - Server port (default: 8080)
- `MINIO_ENDPOINT` - MinIO server endpoint (e.g., minio:9000)
- `MINIO_ACCESS_KEY` - MinIO access credentials
- `MINIO_SECRET_KEY` - MinIO secret credentials
- `S3_BUCKET` - Bucket name for storing assets (default: ticketing-assets)

See `.env.example` for all available options.

### Object Storage (MinIO)

The backend integrates with MinIO for S3-compatible object storage:
- **Local**: MinIO container (http://localhost:9001 console)
- **Production**: AWS S3 or compatible cloud storage
- **Bucket**: Automatically created via `minio-init` service
- **Use case**: Ticket QR codes, event images, receipts

MinIO configuration is handled via environment variables, making it easy to swap between local MinIO and cloud S3 without code changes.

## Documentation

For detailed architecture, design decisions, and deployment strategy:

**[`docs/ARCHITECTURE.md`](../docs/ARCHITECTURE.md)**