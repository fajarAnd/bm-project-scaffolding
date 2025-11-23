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

Everything (Postgres, Redis, API) starts with one command.

### Local Development

```bash
# Copy environment template
cp .env.example .env

# Install dependencies
make deps

# Run server
make run
```

Make sure Postgres and Redis are running locally.

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
```

## Configuration

Environment variables are loaded from `.env` file. Required settings:

- `DATABASE_URL` - Postgres connection string
- `JWT_SECRET` - Token signing key
- `PORT` - Server port (default: 8080)

See `.env.example` for all available options.

## Documentation

For detailed architecture, design decisions, and deployment strategy:

**[`docs/ARCHITECTURE.md`](../docs/ARCHITECTURE.md)**