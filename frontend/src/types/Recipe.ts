export interface RecipeSummary {
    id: number;
    title: string;
    description?: string;
    time?: string;
    servings?: string;
}

export interface Recipe {
    id: number;
    title: string;
    description?: string;
    time?: string;
    servings?: string;
    url?: string;
    notes?: string;
    timesCooked?: number;
    ingredientSections?: IngredientsSection[];
    steps?: Step[];
}

interface IngredientsSection {
    heading: string;
    ingredients: string[];
}

interface Step {
    number: number;
    description: string;
}
