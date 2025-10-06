# Database Migrations

This directory contains SQL migration files for the recipe collector database.

## Running Migrations

1. Set your database environment variables:
   ```bash
   export DBUSER="username"
   export DBPASS="password"
   export DBADDR="127.0.0.1"
   ```

2. Run the migration command:
   ```bash
   cd backend
   go run main.go migrate
   ```

## Migration Files

- `001_create_tables.sql` - Creates all required tables for the application
- `001_create_tables_down.sql` - Rollback migration to drop all tables

## Database Schema

The migration creates the following tables:

### recipes
- Main recipe table with title, description, time, servings, etc.

### ingredient_sections
- Groups ingredients into sections (e.g., "For the sauce", "For the dough")
- References recipes table

### ingredients
- Individual ingredients within sections
- References ingredient_sections table

### steps
- Recipe instruction steps
- References recipes table

## Features

- **Foreign Key Constraints**: Ensures data integrity with CASCADE deletes
- **Indexes**: Optimized for common query patterns
- **Full-text Search**: Enables efficient searching of recipe titles and ingredients
- **Migration Tracking**: Keeps track of applied migrations in `schema_migrations` table
- **Timestamps**: All tables include created_at and updated_at timestamps

## Notes

- The migration tool automatically creates a `schema_migrations` table to track applied migrations
- Migrations are applied in alphabetical order by filename
- Already applied migrations are skipped automatically
- Only files ending in `.sql` (not `_down.sql`) are considered for forward migrations