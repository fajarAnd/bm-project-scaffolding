-- Initialize databases for the ticketing system
-- This script runs once when PostgreSQL container is first created

-- Create ticketing database
CREATE DATABASE ticketing_db;

-- Grant privileges to the default user
GRANT ALL PRIVILEGES ON DATABASE ticketing_db TO postgres;

-- Connect to ticketing_db to create extensions
\c ticketing_db;

-- Enable UUID extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create a comment on the database
COMMENT ON DATABASE ticketing_db IS 'Database for ticketing system - events, tickets, and user management';