package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"recipe-collector/backend/queries"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type recipesResource struct {
	db *sql.DB
}

func NewRecipesResource(db *sql.DB) recipesResource {
	return recipesResource{db: db}
}

func (rr recipesResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/", rr.getAllRecipesHandler)
	r.Post("/", rr.createRecipeHandler)
	r.Get("/search", rr.searchRecipesHandler)
	r.Get("/{recipeID}", rr.getRecipeHandler)

	return r
}

func (rr recipesResource) getAllRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := queries.GetAllRecipes(rr.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	recipesJson, err := json.Marshal(recipes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writeResponse(w, recipesJson)
}

func (rr recipesResource) getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeID, err := strconv.ParseInt(chi.URLParam(r, "recipeID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println(err)
		return
	}

	recipe, err := queries.GetRecipe(rr.db, recipeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if recipe == nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	recipeJson, err := json.Marshal(recipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writeResponse(w, recipeJson)
}

func (rr recipesResource) searchRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := queries.SearchRecipes(rr.db, r.URL.Query().Get("query"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	recipesJson, err := json.Marshal(recipes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writeResponse(w, recipesJson)
}

func (rr recipesResource) createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipeData queries.RecipeInput

	if err := json.NewDecoder(r.Body).Decode(&recipeData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding recipe:", err)
		return
	}

	if recipeData.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "title is required"}`))
		return
	}

	recipeID, err := queries.CreateRecipe(rr.db, &recipeData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error creating recipe:", err)
		return
	}

	response := map[string]interface{}{
		"id":      recipeID,
		"message": "Recipe created successfully",
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, responseJson)
}

func writeResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(data)
}
