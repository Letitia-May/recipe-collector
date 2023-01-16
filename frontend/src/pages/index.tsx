import { Fragment, useEffect, useState } from 'react';
import Head from 'next/head'
import styled from 'styled-components';
import type { Recipe } from '../types/Recipe';

const RecipeWrapper = styled.div`
  padding: 16px;
`;

export default function Home() {
  const [recipes, setRecipes] = useState<Recipe[] | []>([]);

  useEffect(() => {
    fetch('//localhost:8080/recipes', {
      method: 'GET',
      headers: { Accept: 'application/json' },
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not OK');
      }
      return response.json();
    })
    .then((data: Recipe[]) => setRecipes(data))
    .catch(error => {
      console.error('Error:', error);
    });
  }, []);

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

        {recipes?.map(recipe => {
          const {id, title, description, time, servings, url, notes, rating, timesCooked} = recipe;
          return (
            <Fragment key={`recipe-${id}`}>
              <RecipeWrapper>
                <h2>{title}</h2>
                {description && <p>{description}</p>}
                {time && <p>Time to prepare: {time}</p>}
                {servings && <p>Servings: {servings}</p>}
                {notes && <p>Extra notes: {notes}</p>}
                {rating && <p>{rating}</p>}
                {timesCooked && <p>Number of times cooked: {timesCooked}</p>}
                {url && <a href={url}>Original recipe</a>}
              </RecipeWrapper>
              <hr />
            </Fragment>
          )
        })}
      </main>
    </>
  )
}
