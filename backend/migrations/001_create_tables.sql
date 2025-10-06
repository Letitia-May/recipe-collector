-- Migration: 001_create_tables.sql
-- Description: Create all required tables for the recipe collector application
-- Date: 2025-10-06

-- Create recipes table (main table)
CREATE TABLE recipes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    time VARCHAR(100),
    servings VARCHAR(50),
    url TEXT,
    notes TEXT,
    times_cooked BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create ingredient_sections table (groups ingredients like "For the sauce")
CREATE TABLE ingredient_sections (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    recipe_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
);

-- Create ingredients table (individual ingredients within sections)
CREATE TABLE ingredients (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    ingredient_section_id BIGINT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (ingredient_section_id) REFERENCES ingredient_sections(id) ON DELETE CASCADE
);

-- Create steps table (recipe instructions)
CREATE TABLE steps (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    recipe_id BIGINT NOT NULL,
    number BIGINT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
    UNIQUE KEY unique_recipe_step (recipe_id, number)
);

-- Create indexes for better query performance
CREATE INDEX idx_recipes_title ON recipes(title);
CREATE INDEX idx_ingredient_sections_recipe_id ON ingredient_sections(recipe_id);
CREATE INDEX idx_ingredients_section_id ON ingredients(ingredient_section_id);
CREATE INDEX idx_steps_recipe_id ON steps(recipe_id);
CREATE INDEX idx_steps_number ON steps(recipe_id, number);

-- Create full-text search index for recipe titles and ingredient descriptions
-- This will improve search performance for the SearchRecipes function
CREATE FULLTEXT INDEX idx_recipes_title_fulltext ON recipes(title);
CREATE FULLTEXT INDEX idx_ingredients_description_fulltext ON ingredients(description);