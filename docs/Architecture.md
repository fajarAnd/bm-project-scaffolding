# Ticketing System - Architecture Documentation

> **Project Goal:** Create a scalable, maintainable architecture scaffold for a live events ticketing system that demonstrates production-ready patterns


---

## Table of Contents

1. [System Overview](#system-overview)
2. [Architecture Diagrams](#architecture-diagrams)
3. [Technology Stack](#technology-stack)
4. [Trade-offs](#trade-offs)
5. [Deployment Strategy](#deployment-strategy)
6. [Security Considerations](#security-considerations)
7. [Scalability Considerations](#scalability-considerations)


---

## System Overview

The ticketing system is designed as a **modern web application** with clear separation between frontend, backend, and infrastructure layers. The architecture supports both **local development** (Docker Compose) and **cloud deployment** (AWS via Terraform).

### Core Features

* **Authentication & Authorization** - JWT-based with role-based access control (user/admin)
* **Event Management** - Public event listings with detailed views
* **Ticket Purchasing** - Protected ticket purchase flow with inventory management
* **User Management** - CRUD operations with role-based permissions

### Key Architectural Principles

1. **Layered Architecture** - Clear separation: Handlers → Services → Repositories
2. **Stateless Backend** - Horizontal scaling capability via JWT (no session storage)
3. **Infrastructure as Code** - Repeatable deployments via Terraform
4. **Containerization** - Consistent environments from dev to production
5. **Security by Design** - Authentication, authorization, and secure defaults


---

## Architecture Diagrams

> Using **4+1 Architectural View Model** to describe the system from different perspectives.

### 1. Logical View - System Context

```mermaid

graph TB
    User[Customer]
    Admin[Admin]

    subgraph "Ticketing System"
        FE[React Frontend<br/>Port 5173]
        API[Go Backend API<br/>Port 8080]
        DB[(PostgreSQL<br/>Database)]
        Cache[(Redis<br/>Cache)]
        Storage[Object Storage<br/>MinIO/S3]
    end

    Payment[Payment Gateway<br/>Stripe/PayPal<br/><i>STUB</i>]
    Email[Email Service<br/>SendGrid<br/><i>STUB</i>]

    User -->|HTTPS| FE
    Admin -->|HTTPS| FE
    FE -->|REST API| API
    API --> DB
    API --> Cache
    API --> Storage
    API -.->|TODO| Payment
    API -.->|TODO| Email
```

**Legend:** Solid lines = Implemented, Dashed lines = Stubbed/TODO

---

### 2. Logical View - Component Architecture

```mermaid

graph LR
    subgraph "Client Layer"
        Client[HTTP Client<br/>Frontend/Mobile]
    end

    subgraph "Presentation Layer"
        Router[Router]
        MW[Middleware Chain<br/>Auth, CORS]
        Handler[Handlers<br/>HTTP Req/Resp]
    end

    subgraph "Business Logic Layer"
        Svc[Service]
    end

    subgraph "Data Access Layer"
        Repo[Repository]
    end

    subgraph "Infrastructure Layer"
        DB[(PostgreSQL)]
        Cache[(Redis)]
        Storage[Object Storage]
    end

    Client --> Router
    Router --> MW
    MW --> Handler
    Handler --> Svc

    Svc --> Repo

    Repo --> DB
    Svc -.-> Cache
```

**Patterns used:**
- Dependency Injection for services
- Repository pattern for data access
- Middleware chain for cross-cutting concerns

---

### 3. Process View - Authentication Flow

```mermaid

sequenceDiagram
    actor User
    participant FE as Frontend
    participant API as Backend API
    participant Auth as Auth Middleware
    participant Service as Auth Service
    participant Repo as User Repository
    participant DB as PostgreSQL

    User->>FE: Enter email/password
    FE->>API: POST /api/auth/login<br/>{email, password}
    API->>Service: Login(email, password)
    Service->>Repo: FindByEmail(email)
    Repo->>DB: Get User
    DB-->>Repo: User record
    Repo-->>Service: User
    Service->>Service: Validate password
    Service->>Service: Generate JWT
    Service-->>API: {token, user}
    API-->>FE: 200 OK {token, user}
    FE->>FE: Store token in localStorage

    Note over User,DB: Subsequent Authenticated Requests

    User->>FE: Request protected resource
    FE->>API: GET /api/users<br/>Authorization: Bearer {token}
    API->>Auth: Validate JWT
    Auth->>Auth: Parse token
    Auth-->>API: User context
    API->>Service: GetCurrentUser(userID)
    Service->>Repo: FindByID(userID)
    Repo->>DB: Get User
    DB-->>Repo: User record
    Repo-->>Service: User
    Service-->>API: User data
    API-->>FE: 200 OK {user}
    FE-->>User: Display user info
```


---

### 4. Process View - Ticket Purchase Flow

```mermaid

sequenceDiagram
    actor User
    participant FE as Frontend
    participant API as Backend API
    participant TicketSvc as Ticket Service
    participant Payment as Payment Service (Stub)
    participant DB as PostgreSQL

    User ->> FE: Click "Purchase Ticket"
    FE ->> API: POST /tickets/purchase {eventID, qty, token}

    API ->> API: Validate JWT

    API ->> TicketSvc: PurchaseTicket(userID, eventID, qty)

    TicketSvc ->> DB: SELECT event WITH LOCK
    DB -->> TicketSvc: Event data

    TicketSvc ->> TicketSvc: Check availability

    alt Tickets Available
        TicketSvc ->> Payment: Charge()
        Payment -->> TicketSvc: Success

        TicketSvc ->> DB: Decrement tickets + Create ticket
        DB -->> TicketSvc: OK

        TicketSvc -->> API: success {ticketID}
        API -->> FE: 200 OK
        FE -->> User: Show confirmation
    else Sold Out
        TicketSvc -->> API: Error "Sold out"
        API -->> FE: 409 Conflict
        FE -->> User: Show sold out message
    end
```


---

### 5. Development View - Deployment

#### Local Development Environment

```mermaid

graph TB
    subgraph "Developer Machine"
        Dev[Developer]
        Docker[Docker Engine]

        subgraph "Docker Compose Network"
            FE[Frontend Container<br/>React + Vite<br/>Port 5173]
            BE[Backend Container<br/>Go + Gin<br/>Port 8080]
            PG[(PostgreSQL<br/>Port 5432)]
            RD[(Redis<br/>Port 6379)]
            MN[MinIO<br/>Ports 9000, 9001]
        end
    end

    Dev -->|docker-compose up| Docker
    Docker --> FE
    Docker --> BE
    Docker --> PG
    Docker --> RD
    Docker --> MN

    FE -->|API Calls| BE
    BE --> PG
    BE --> RD
    BE --> MN
```

**Command:**

```bash
docker-compose up
```

**Access:**
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- MinIO Console: http://localhost:9001

---

#### Cloud Production Environment (AWS)

```mermaid

graph TB
    subgraph "AWS Cloud"
        subgraph "Public Subnet"
            ALB[Application Load Balancer<br/>Port 443 HTTPS]
            CF[Cloudflare CDN<br/>Frontend Assets]
        end

        subgraph "Private Subnet"
            ECS[ECS Fargate<br/>Go Backend API<br/>Auto-scaling 2-10 tasks]
            RDS[(RDS PostgreSQL<br/>Multi-AZ)]
            EC[(ElastiCache Redis<br/>Cluster Mode)]
        end

        S3[S3 Bucket<br/>Object Storage<br/>+ Frontend Static Assets]

        Cognito[AWS Cognito<br/>Auth Provider<br/><i>TODO</i>]
    end

    User[Users]

    User -->|HTTPS| CF
    User -->|HTTPS| ALB
    CF --> S3
    ALB --> ECS
    ECS --> RDS
    ECS --> EC
    ECS --> S3
    ECS -.->|TODO| Cognito
```

**AWS Services:**
- ECS Fargate - Serverless containers with auto-scaling
- RDS PostgreSQL - Managed database
- ElastiCache Redis - Managed cache
- S3 - Object storage + static assets
- Cloudflare CDN - Global distribution with DDoS protection
- ALB - Load balancer with health checks
- VPC - Network isolation

---

### 6. Data View - Entity Relationship Diagram

```mermaid 

erDiagram
    users ||--o{ ticket_orders : places
    events ||--o{ ticket_orders : has
    ticket_orders ||--o{ tickets : contains

    users {
        uuid id PK
        varchar email UK
        varchar password_hash
        varchar role
        timestamp created_at
        timestamp updated_at
    }

    events {
        uuid id PK
        varchar title
        text description
        timestamp event_date
        varchar venue
        decimal ticket_price
        int total_tickets
        int available_tickets
        timestamp created_at
        timestamp updated_at
    }

    ticket_orders {
        uuid id PK
        uuid event_id FK
        uuid user_id FK
        int quantity
        decimal total_price
        varchar status
        varchar payment_id
        timestamp created_at
        timestamp updated_at
    }

    tickets {
        uuid id PK
        uuid order_id FK
        varchar ticket_code UK
        varchar status
        timestamp used_at
        timestamp created_at
    }
```

**Ticket purchase mechanism:**
1. **Availability tracking:**
   - `events.total_tickets` - Original capacity
   - `events.available_tickets` - Current inventory (decremented on order, not ticket generation)
   - Inventory management decoupled from ticket generation

2. **Order and ticket states:**
   - `ticket_orders.status`: `pending`, `paid`, `confirmed`, `cancelled`, `refunded`
   - `tickets.status`: `valid`, `used`, `cancelled`, `transferred`
   - Payment processed at order level, not individual tickets

**Constraints:**
- `users.email` - Unique constraint
- `ticket_orders.event_id` - Foreign key with ON DELETE RESTRICT
- `ticket_orders.user_id` - Foreign key with ON DELETE CASCADE
- `tickets.order_id` - Foreign key with ON DELETE CASCADE
- `tickets.ticket_code` - Unique constraint (generated when ticket created)
- Check constraint: `events.available_tickets >= 0`
- Check constraint: `ticket_orders.quantity > 0`

---

## Technology Stack

### Backend Stack

| Component | Technology | Rationale |
|----|----|----|
| **Language** | Go | Strong concurrency (goroutines), small binaries (\~20MB), fast compilation, excellent for high-throughput APIs |
| **Web Framework** | Gin | Fast HTTP router , built-in validation, middleware ecosystem, minimal overhead |
| **Database** | PostgreSQL | ACID compliance for ticket transactions, row-level locking prevents double-booking, rich query capabilities |
| **Cache** | Redis | In-memory cache for event listings, session data (future), rate limiting capability |
| **Database Client** | sqlx | Light wrapper over database/sql with struct scanning, not a full ORM (keeps SQL visible) |
| **Migrations** | golang-migrate | Database-agnostic, CLI tool, supports up/down migrations |
| **Auth** | golang-jwt/jwt | JWT generation and validation, industry standard |
| **Environment Config** | godotenv | Load .env files for local development |

### Frontend Stack

| Component | Technology | Rationale |
|----|----|----|
| **Framework** | React | Industry standard, large ecosystem, well-understood by most developers |
| **Build Tool** | Vite | Fast HMR (Hot Module Replacement), modern ESM-based, smaller bundles than Webpack |
| **Language** | TypeScript | Type safety reduces runtime errors, pairs well with Go's strong typing |
| **Router** | React Router | Standard routing solution, supports protected routes |
| **HTTP Client** | Axios | Interceptors for auth token injection, better error handling than fetch |
| **State Management** | Context API | Sufficient for auth state, avoids Redux complexity for scaffold |

### Infrastructure Stack

| Component | Local (Docker) | Cloud (AWS) | Rationale |
|----|----|----|----|
| **Compute** | Docker Container | ECS Fargate | Serverless containers, auto-scaling, no EC2 management |
| **Database** | PostgreSQL Container | RDS PostgreSQL | Managed service, automated backups, Multi-AZ for HA |
| **Cache** | Redis Container | ElastiCache Redis | Managed Redis, cluster mode for scalability |
| **Object Storage** | MinIO | S3 | S3-compatible locally, production-ready globally-distributed storage |
| **CDN** | - | Cloudflare | Global CDN with DDoS protection, HTTPS termination |
| **Load Balancer** | - | ALB | Health checks, SSL termination, and auto-scaling integration |
| **IaC** | Docker Compose | Terraform | Infrastructure as Code, repeatable deployments |


---

## Trade-offs

### Decisions Made

| Decision | Benefit | Cost | Mitigation |
|----|----|----|----|
| **JWT (Stateless)** | Horizontal scaling, no session store | Cannot revoke tokens | Short TTL (24h) + refresh tokens (future) |
| **Docker Compose** | Easy local dev, environment parity | Requires Docker install (\~1GB) | Docker is industry standard, acceptable |
| **Gin Framework** | Fast, built-in features | Slightly opinionated | Still lightweight, easy to swap if needed |
| **Postgres (not NoSQL)** | Strong consistency, ACID | Complex setup vs. MongoDB | Necessary for ticket transactions |


---

## Deployment Strategy

### Local Development

Prerequisites: Docker Desktop + Git

**Setup:**
```bash
git clone <repository-url>
cd ticketing-system
docker-compose up
```

**Teardown:**
```bash
docker-compose down -v
```

---

### Cloud Deployment (AWS)

Prerequisites: AWS account + AWS CLI + Terraform

**Setup:**
```bash
cd infra/terraform
terraform init
terraform plan
terraform apply
```

**Provisioned resources:**
- VPC with public/private subnets
- ECS Fargate cluster with auto-scaling
- RDS PostgreSQL (Multi-AZ)
- ElastiCache Redis
- S3 bucket + Cloudflare CDN
- Application Load Balancer
- Security groups

**Teardown:**
```bash
terraform destroy
```


---

## Security Considerations

### Current Implementation

1. JWT-based authentication (token in localStorage)
2. Role-based access control (user/admin)
3. Environment variables for secrets (`.env`)
4. CORS configured for frontend origins

### Production TODOs

**Network:**
- Deploy backend in private subnets only
- Use VPC security groups (least-privilege)
- Enable Cloudflare WAF rules for DDoS protection

**Database:**
- IAM database authentication
- Automated backups (7-day retention)

**Secrets:**
- Migrate to AWS Secrets Manager
- Implement secret rotation

**Input Validation:**
- Comprehensive sanitization (XSS, SQL injection)
- Rate limiting per user

**Logging:**
- CloudWatch Logs for API requests
- CloudTrail for infrastructure changes
- Alerting on suspicious activity


---

## Scalability Considerations

### Current Design

1. **Stateless Backend** - JWT enables horizontal scaling
2. **Database** - Connection pooling, supports read replicas
3. **Caching** - Redis for frequently accessed data
4. **CDN** - Cloudflare for global static asset delivery

### Future Improvements

**Database:**
- Add read replicas for heavy queries
- Consider partitioning tickets table by date
- Sharding for multi-region (if needed)

**Backend:**
- Auto-scaling based on CPU/memory metrics
- Circuit breakers for external services

**Frontend:**
- Code splitting for faster initial load
- Lazy loading for routes

**Observability:**
- Distributed tracing with OpenTelemetry
- Metrics and monitoring via Grafana + Prometheus

**Migration Strategy:**
- Strangler Fig pattern for gradual microservice migration
- Start with event service as independent microservice
- Keep monolith as API gateway during transition


---

## References

* [Go Project Layout](https://github.com/golang-standards/project-layout) - Standard Go structure
* [12-Factor App](https://12factor.net/) - Modern app best practices
* [AWS Well-Architected Framework](https://aws.amazon.com/architecture/well-architected/) - Cloud architecture principles
* [OWASP Top 10](https://owasp.org/www-project-top-ten/) - Security best practices
* [4+1 Views in Modeling System Architecture](https://guides.visual-paradigm.com/4-1-views-in-modeling-system-architecture-with-uml/)