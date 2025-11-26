#!/bin/bash

# Destroy ticketing infrastructure
# WARNING: This will delete everything!

set -e

TERRAFORM_DIR="../terraform"
ENV="${1:-dev}"
VAR_FILE="environments/${ENV}.tfvars"

echo "========================================="
echo "Infrastructure Teardown - $ENV"
echo "========================================="
echo ""

# Basic checks
if ! command -v terraform &> /dev/null; then
    echo "Error: terraform not installed"
    exit 1
fi

cd "$TERRAFORM_DIR" || exit 1

if [ ! -f "$VAR_FILE" ]; then
    echo "Error: $VAR_FILE doesn't exist"
    exit 1
fi

# Need these vars for the destroy plan
read -sp "DB Password: " DB_PASSWORD
echo ""
read -sp "JWT Secret: " JWT_SECRET
echo ""

export TF_VAR_db_password="$DB_PASSWORD"
export TF_VAR_jwt_secret="$JWT_SECRET"

# Show what we're about to destroy
echo ""
echo "Generating destroy plan..."
terraform plan -destroy -var-file="$VAR_FILE"

# Double-check with user
echo ""
echo "WARNING: This will permanently delete:"
echo "- All ECS services and containers"
echo "- RDS database (all data will be lost)"
echo "- S3 bucket and objects"
echo "- ElastiCache Redis cluster"
echo "- VPC and networking resources"
echo ""
read -p "Type 'destroy' to confirm deletion: " CONFIRM

if [ "$CONFIRM" != "destroy" ]; then
    echo "Cancelled - nothing was deleted"
    exit 0
fi

# Actually destroy everything
echo ""
echo "Destroying infrastructure..."
terraform destroy -var-file="$VAR_FILE" -auto-approve

echo ""
echo "Done! All infrastructure has been destroyed."