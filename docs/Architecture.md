# Ticketing System - Architecture Documentation

> **Project Goal:** Create a scalable, maintainable architecture scaffold for a live events ticketing system that demonstrates production-ready patterns




## Table of Contents

1. [System Overview](#system-overview)
2. [Architecture Diagrams](#architecture-diagrams)

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

### 1. System Context (High-Level Architecture)

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

### 2. Backend Layered Architecture

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

### 3. Authentication Flow (Sequence Diagram)

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

### 4. Ticket Purchase Flow (Sequence Diagram)

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

### 5. Deployment Architecture

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

