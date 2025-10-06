-- Migration: 002_seed_sample_recipe.sql
-- Description: Add sample recipe if the database is empty
-- Date: 2025-10-06

-- Insert a sample recipe only if no recipes exist
INSERT INTO recipes (title, description, time, servings, url, notes, times_cooked)
SELECT 
    'Really Quick Broccoli Pasta',
    "This is a great emergency meal OR carb + veg side for all those times when your cupboards are bare except for broccoli, pasta, and some kind of cheese. It's saucy without using tons of oil, and there's loads of sub options.",
    '20 minutes',
    '4 serves',
    'https://www.recipetineats.com/quick-broccoli-pasta/#recipe',
    "Cheese - any melting cheese fine here, preferably flavoured like cheddar, Monterey Jack, tasty, gruyere, Swiss. Mozzarella also fine but you'll probably need more salt.",
    2
WHERE NOT EXISTS (SELECT 1 FROM recipes);

-- Get the recipe ID for the sample recipe we just inserted
SET @recipe_id = (SELECT id FROM recipes WHERE title = 'Really Quick Broccoli Pasta' LIMIT 1);

-- Insert ingredient sections only if the recipe was inserted
INSERT INTO ingredient_sections (recipe_id, name)
SELECT @recipe_id, 'Main Ingredients'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 'Pasta Sauce'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 'Serving'
WHERE @recipe_id IS NOT NULL;

-- Get the section IDs
SET @main_ingredients_id = (SELECT id FROM ingredient_sections WHERE recipe_id = @recipe_id AND name = 'Main Ingredients' LIMIT 1);
SET @pasta_sauce_id = (SELECT id FROM ingredient_sections WHERE recipe_id = @recipe_id AND name = 'Pasta Sauce' LIMIT 1);
SET @serving_id = (SELECT id FROM ingredient_sections WHERE recipe_id = @recipe_id AND name = 'Serving' LIMIT 1);

-- Insert ingredients for main ingredients section
INSERT INTO ingredients (ingredient_section_id, description)
SELECT @main_ingredients_id, '350g / 12 oz dried short pasta (I used small shells)'
WHERE @main_ingredients_id IS NOT NULL
UNION ALL
SELECT @main_ingredients_id, '2 broccoli heads (BIG!)'
WHERE @main_ingredients_id IS NOT NULL
UNION ALL
SELECT @main_ingredients_id, '1 cup shredded cheese (or more!)'
WHERE @main_ingredients_id IS NOT NULL;

-- Insert ingredients for pasta sauce section
INSERT INTO ingredients (ingredient_section_id, description)
SELECT @pasta_sauce_id, '2 tsp lemon zest'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '2 tbsp lemon juice (or more!)'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '5 tbsp extra virgin olive oil'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '1/3 cup parmesan, finely grated'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '2 garlic cloves, minced'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '1 tsp mixed dried herbs (or fresh!)'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '1/2 tsp+ red pepper flakes (add more if you want spicy!)'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '1 tsp sugar'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '3/4 tsp salt'
WHERE @pasta_sauce_id IS NOT NULL
UNION ALL
SELECT @pasta_sauce_id, '1/2 tsp pepper'
WHERE @pasta_sauce_id IS NOT NULL;

-- Insert ingredients for serving section
INSERT INTO ingredients (ingredient_section_id, description)
SELECT @serving_id, 'More parmesan'
WHERE @serving_id IS NOT NULL;

-- Insert recipe steps
INSERT INTO steps (recipe_id, number, description)
SELECT @recipe_id, 1, 'Cook pasta: Boil a large pot of water with 2 tsp salt, add pasta.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 2, 'Chop broccoli: Chop broccoli into small florets.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 3, 'Cook broccoli: Add broccoli into water 1 - 2 minutes before pasta is cooked.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 4, 'Sauce ingredients in jar: Place Sauce ingredients in a jar with lid.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 5, 'Reserve pasta water: SCOOP OUT 1 cup pasta cooking water just before draining. Then drain and return pasta back into same pot on turned off stove.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 6, 'Add pasta water into Sauce: Add 1/2 cup pasta water to Pasta Sauce jar, shake well.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 7, 'Add Sauce & Cheese to pasta: Pour Sauce and add cheese into pot with pasta.'
WHERE @recipe_id IS NOT NULL
UNION ALL
SELECT @recipe_id, 8, 'Stir then serve! Stir vigorously, adding more pasta water if required. Add more salt and pepper if required. Serve immediately, garnished with parmesan.'
WHERE @recipe_id IS NOT NULL;