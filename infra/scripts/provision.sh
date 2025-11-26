#!/bin/bash

# Deploy ticketing infrastructure to AWS
set -e

TERRAFORM_DIR="../terraform"
ENV="${1:-dev}"
VAR_FILE="environments/${ENV}.tfvars"

echo "=== Deploying Ticketing Infrastructure ==="
echo "Environment: $ENV"
echo ""

# Quick sanity checks
if ! command -v terraform &> /dev/null; then
    echo "Error: terraform not found. Install from https://www.terraform.io/downloads"
    exit 1
fi

cd "$TERRAFORM_DIR" || exit 1

if [ ! -f "$VAR_FILE" ]; then
    echo "Error: $VAR_FILE not found"
    exit 1
fi

# Get secrets (don't hardcode these!)
echo "Enter sensitive variables:"
read -sp "DB Password: " DB_PASSWORD
echo ""
read -sp "JWT Secret: " JWT_SECRET
echo ""

export TF_VAR_db_password="$DB_PASSWORD"
export TF_VAR_jwt_secret="$JWT_SECRET"

# Init and validate
echo ""
echo "Initializing terraform..."
terraform init

echo "Validating config..."
terraform validate

# Format if needed
terraform fmt -check > /dev/null 2>&1 || terraform fmt

# Create plan
echo ""
echo "Creating plan..."
terraform plan -var-file="$VAR_FILE" -out=tfplan

# Confirm before applying
echo ""
echo "---"
echo "Ready to deploy to $ENV environment"
read -p "Continue? (yes/no): " CONFIRM

if [ "$CONFIRM" != "yes" ]; then
    echo "Cancelled"
    rm -f tfplan
    exit 0
fi

# Apply the plan
echo ""
echo "Applying infrastructure changes..."
terraform apply tfplan
rm -f tfplan

# Show outputs
echo ""
echo "=== Deployment Complete ==="
terraform output

echo ""
echo "Infrastructure deployed successfully!"
echo "Note: Save the outputs above - you'll need them to configure the backend"