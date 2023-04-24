export interface Recipe {
    id: number;
    title: string;
    description?: string;
    time?: string;
    servings?: string;
    url?: string;
    notes?: string;
    rating?: number;
    timesCooked?: number;
}
