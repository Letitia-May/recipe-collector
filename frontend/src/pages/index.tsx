import { Fragment } from 'react';
import Head from 'next/head';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import type { Recipe } from '@/types/Recipe';

export default function Home({ recipes }: InferGetServerSidePropsType<typeof getServerSideProps>) {
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

                {recipes?.map((recipe) => {
                    const {
                        id,
                        title,
                        description,
                        time,
                        servings,
                        url,
                        notes,
                        rating,
                        timesCooked,
                    } = recipe;
                    return (
                        <Fragment key={`recipe-${id}`}>
                            <div>
                                <h2>{title}</h2>
                                {description && <p>{description}</p>}
                                {time && <p>Time to prepare: {time}</p>}
                                {servings && <p>Servings: {servings}</p>}
                                {notes && <p>Extra notes: {notes}</p>}
                                {rating && <p>{rating}</p>}
                                {timesCooked && <p>Number of times cooked: {timesCooked}</p>}
                                {url && <a href={url}>Original recipe</a>}
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
    recipes: Recipe[];
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
