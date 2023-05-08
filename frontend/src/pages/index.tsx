import { Fragment, useState } from 'react';
import Head from 'next/head';
import Link from 'next/link';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import type { RecipeSummary } from '@/types/Recipe';

export default function Home({ recipes }: InferGetServerSidePropsType<typeof getServerSideProps>) {
    const [searchTerm, setSearchTerm] = useState('');

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
                <button type="submit" onClick={() => console.log(searchTerm)}>
                    Search
                </button>

                {recipes?.map((recipe) => {
                    const { id, title, description, time, servings } = recipe;
                    return (
                        <Fragment key={`recipe-${id}`}>
                            <div>
                                <h2>{title}</h2>
                                {description && <p>{description}</p>}
                                {time && <p>Time to prepare: {time}</p>}
                                {servings && <p>Servings: {servings}</p>}
                                <Link href={`/recipe/${id}`}>See recipe</Link>
                            </div>
                            <hr />
                        </Fragment>
                    );
                })}
            </main>
        </>
    );
}

interface ServerSideProps {
    recipes: RecipeSummary[];
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
