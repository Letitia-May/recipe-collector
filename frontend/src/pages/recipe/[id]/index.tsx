import { useRouter } from 'next/router';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import type { Recipe as RecipeData } from '@/types/Recipe';

export default function Recipe({ recipe }: InferGetServerSidePropsType<typeof getServerSideProps>) {
    const router = useRouter();
    const { id } = router.query;
    console.log(recipe);

    return <p>Recipe: {id}</p>;
}

interface ServerSideProps {
    recipe: RecipeData;
}

export const getServerSideProps: GetServerSideProps<ServerSideProps> = async ({ query }) => {
    const { id } = query;

    const response = await fetch(`http://localhost:8080/recipes/${id}`, {
        method: 'GET',
        headers: { Accept: 'application/json' },
    });

    if (!response.ok) {
        throw new Error('Network response was not OK');
    }

    if (!response.json) {
        return {
            notFound: true,
        };
    }

    const recipe = await response.json();

    return {
        props: {
            recipe,
        },
    };
};
