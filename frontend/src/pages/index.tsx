import { Fragment, useState } from 'react';
import Head from 'next/head';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import type { RecipeSummary as RecipeSummaryType } from '@/types/Recipe';
import { RecipeSummary } from 'components/RecipeSummary/RecipeSummary';

export default function Home({ recipes }: InferGetServerSidePropsType<typeof getServerSideProps>) {
    const [searchTerm, setSearchTerm] = useState('');
    const [filteredRecipes, setFilteredRecipes] = useState<RecipeSummaryType[] | null>(recipes);

    const searchRecipes = async () => {
        fetch(`//localhost:8080/recipes/search?query=${searchTerm}`, {
            method: 'GET',
            headers: { Accept: 'application/json' },
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Network response was not OK');
                }
                return response.json();
            })
            .then((data: RecipeSummaryType[]) => setFilteredRecipes(data))
            .catch((error) => {
                console.error('Error:', error);
            });
    };

    return (
        <>
            <Head>
                <title>Recipe collector</title>
                <meta name="description" content="Collection of recipes that I love" />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <main>
                <h1>Recipe collector</h1>

                <label>
                    Search recipe titles:
                    <input
                        type="text"
                        placeholder="Type here..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                    />
                </label>
                <button type="submit" onClick={() => searchRecipes()}>
                    Search
                </button>

                {!filteredRecipes && <p>No recipes found</p>}

                {filteredRecipes?.map((recipe) => {
                    return (
                        <Fragment key={`recipe-${recipe.id}`}>
                            <RecipeSummary recipe={recipe} />
                            <hr />
                        </Fragment>
                    );
                })}
            </main>
        </>
    );
}

interface ServerSideProps {
    recipes: RecipeSummaryType[];
}

export const getServerSideProps: GetServerSideProps<ServerSideProps> = async () => {
    const response = await fetch(`http://localhost:8080/recipes`, {
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

    const recipes = await response.json();

    return {
        props: {
            recipes,
        },
    };
};
