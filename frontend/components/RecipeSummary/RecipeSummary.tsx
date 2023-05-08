import Link from 'next/link';
import { RecipeSummary as RecipeSummaryType } from '@/types/Recipe';

interface RecipeSummaryProps {
    recipe: RecipeSummaryType;
}

export const RecipeSummary = ({ recipe }: RecipeSummaryProps) => {
    const { id, title, description, time, servings } = recipe;

    return (
        <div>
            <h2>{title}</h2>
            {description && <p>{description}</p>}
            {time && <p>Time to prepare: {time}</p>}
            {servings && <p>Servings: {servings}</p>}
            <Link href={`/recipe/${id}`}>See recipe</Link>
        </div>
    );
};
