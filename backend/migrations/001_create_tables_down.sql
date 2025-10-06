-- Migration: 001_create_tables_down.sql
-- Description: Rollback migration to drop all tables created in 001_create_tables.sql
-- Date: 2025-10-06

-- Drop tables in reverse order (due to foreign key constraints)
DROP TABLE IF EXISTS steps;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS ingredient_sections;
DROP TABLE IF EXISTS recipes;