# Ticketing System Scaffold

Production-ready scaffold for a live events ticketing system with complete local development infrastructure.

## Project Structure

```
.
├── backend/           # Go REST API
├── frontend/          # React application
├── infra/            # Infrastructure automation (Terraform)
├── docs/             # Architecture documentation
├── scripts/          # Database initialization scripts
├── docker-compose.yaml
└── .env.example
```

## Quick Start

### Prerequisites
- Docker Desktop installed
- Git

### One-Command Setup

```bash
# Clone and start everything
git clone <repository-url>
cd bm-project-scaffolding
cp .env.example .env
docker-compose up -d
```

This starts all services:
- **PostgreSQL** (port 5434) - Database
- **Redis** (port 6380) - Cache
- **MinIO** (ports 9000, 9001) - Object storage (S3-compatible)
- **Backend API** (port 8091) - Go backend
- **Frontend** (port 5173) - React application

### Access Points

- Frontend: http://localhost:5173
- Backend API: http://localhost:8091
- MinIO Console: http://localhost:9001 (credentials: minioadmin / minioadmin123)
- API Health Check: http://localhost:8091/health

## Architecture

See detailed documentation:
- [`docs/Architecture.md`](docs/Architecture.md) - System architecture and design decisions
- [`backend/README.md`](backend/README.md) - Backend structure and API details
- [`frontend/README.md`](frontend/README.md) - Frontend setup and routing

## Production Deployment

Deploy to AWS using Terraform automation in [`infra/`](infra/):

```bash
cd infra/scripts
./provision.sh prod
```

This creates:
- **VPC** with public/private subnets (Multi-AZ)
- **ECS Fargate** for backend (auto-scaling 2-10 tasks)
- **RDS PostgreSQL** (encrypted, automated backups)
- **ElastiCache Redis** for caching
- **S3** for object storage
- **ALB** for load balancing

After deployment, get endpoints:
```bash
cd ../terraform
terraform output
```

Update backend `.env` with the outputs (RDS endpoint, Redis URL, S3 bucket). Frontend deploys to S3 + CloudFront (configure separately or use Vercel/Netlify).

Full deployment guide: [`infra/README.md`](infra/README.md)
