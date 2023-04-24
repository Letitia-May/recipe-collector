export interface Recipe {
    id: number;
    title: string;
    description?: string;
    time?: string;
    servings?: string;
    url?: string;
    notes?: string;
    rating?: number;
    times_cooked?: number;
    ingredient_sections?: IngredientsSection[];
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
