# Production Environment Configuration

project_name = "ticketing"
environment  = "prod"
aws_region   = "us-east-1"

# Networking
vpc_cidr             = "10.1.0.0/16"
availability_zones   = ["us-east-1a", "us-east-1b", "us-east-1c"]
public_subnet_cidrs  = ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
private_subnet_cidrs = ["10.1.10.0/24", "10.1.11.0/24", "10.1.12.0/24"]

# Database
db_name              = "ticketing_db"
db_username          = "postgres"
# db_password        = "CHANGE_ME" # Set via environment variable or secrets manager
db_instance_class    = "db.t3.small"
db_allocated_storage = 100
db_multi_az          = true

# Storage
s3_bucket_name = "" # Will be auto-generated

# Compute
container_image   = "YOUR_ECR_REPO/ticketing-backend:latest"
container_port    = 8080
ecs_desired_count = 3
ecs_cpu           = "512"
ecs_memory        = "1024"

# Application
# jwt_secret = "CHANGE_ME" # Set via environment variable or secrets manager