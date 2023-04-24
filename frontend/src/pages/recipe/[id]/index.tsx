import Head from 'next/head';
import { Fragment } from 'react';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import type { Recipe as RecipeData } from '@/types/Recipe';

export default function Recipe({ recipe }: InferGetServerSidePropsType<typeof getServerSideProps>) {
    const {
        title,
        description,
        time,
        servings,
        url,
        notes,
        times_cooked,
        ingredient_sections,
        steps,
    } = recipe;

    return (
        <>
            <Head>
                <title>{`Recipe - ${title}`}</title>
                <meta name="description" content={`Recipe for ${title}`} />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <main>
                <h1>{title}</h1>
                {description && <p>{description}</p>}
                {time && <p>Time to prepare: {time}</p>}
                {servings && <p>Servings: {servings}</p>}
                {times_cooked && <p>Number of times cooked: {times_cooked}</p>}

                {ingredient_sections && (
                    <>
                        <h2>Ingredients</h2>
                        {ingredient_sections.map((section, index) => (
                            <Fragment key={`ingredient-section-${index}`}>
                                {section.heading !== 'Ingredients' && <h3>{section.heading}</h3>}
                                <ul>
                                    {section.ingredients.map((ingredient, i) => (
                                        <li key={`ingredient-${i}`}>{ingredient}</li>
                                    ))}
                                </ul>
                            </Fragment>
                        ))}
                    </>
                )}

                {steps && (
                    <>
                        <h2>Instructions</h2>
                        {steps.map((step, index) => (
                            <Fragment key={`step-${index}`}>
                                <h3>{`Step ${step.number}`}</h3>
                                <p>{step.description}</p>
                            </Fragment>
                        ))}
                    </>
                )}

                <hr />
                {notes && <p>Extra notes: {notes}</p>}
                {url && <a href={url}>Original recipe</a>}
            </main>
        </>
    );
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
