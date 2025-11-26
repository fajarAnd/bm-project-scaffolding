# Development Environment Configuration

project_name = "ticketing"
environment  = "dev"
aws_region   = "us-east-1"

# Networking
vpc_cidr             = "10.0.0.0/16"
availability_zones   = ["us-east-1a", "us-east-1b"]
public_subnet_cidrs  = ["10.0.1.0/24", "10.0.2.0/24"]
private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"]

# Database
db_name            = "ticketing_db"
db_username        = "postgres"
# db_password      = "CHANGE_ME" # Set via environment variable or prompt
db_instance_class  = "db.t3.micro"
db_allocated_storage = 20
db_multi_az        = false

# Storage
s3_bucket_name = "" # Will be auto-generated

# Compute
container_image   = "nginx:latest" # Replace with actual image
container_port    = 8080
ecs_desired_count = 1
ecs_cpu           = "256"
ecs_memory        = "512"

# Application
# jwt_secret = "CHANGE_ME" # Set via environment variable or prompt