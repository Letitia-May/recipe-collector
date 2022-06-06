import React, { useEffect, useState } from "react"

interface Recipe {
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

const IndexPage = () => {
  const [recipes, setRecipes] = useState<Recipe[] | []>([])

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
    <main>
      <title>All recipes</title>
      <h1>
        Recipe collector
      </h1>

      {recipes?.map(recipe => {
        const {id, title, description, time, servings, url, notes, rating, timesCooked} = recipe;
        return (
          <div key={`recipe-${id}`}>
            <h2>{title}</h2>
            {description && <p>{description}</p>}
            {time && <p>Time to prepare: {time}</p>}
            {servings && <p>Servings: {servings}</p>}
            {url && <a href={url}>Link to recipe</a>}
            {notes && <p>Extra notes: {notes}</p>}
            {rating && <p>{rating}</p>}
            {timesCooked && <p>Number of times cooked: {timesCooked}</p>}
            <hr />
          </div>
        )})}
    </main>
  )
}

export default IndexPage
