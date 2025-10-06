-- Migration: 002_seed_sample_recipe_down.sql
-- Description: Remove sample recipe added by the seed migration
-- Date: 2025-10-06

-- Remove the sample recipe and all related data (CASCADE will handle related tables)
DELETE FROM recipes WHERE title = 'Really Quick Broccoli Pasta';