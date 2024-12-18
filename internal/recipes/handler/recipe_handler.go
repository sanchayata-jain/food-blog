package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sanchayata-jain/food-blog/internal/recipes/models"
)

type RecipeService interface {
	CreateRecipe(ctx context.Context, recipe *models.Recipe) error
	GetRecipes(ctx context.Context) ([]models.Recipe, error)
}

type Handler struct {
	Service RecipeService
}

func NewHandler(service RecipeService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var recipe *models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipe.ID = uuid.NewString()
	err = h.Service.CreateRecipe(ctx, recipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	recipes, err := h.Service.GetRecipes(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	recipesBytes, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(recipesBytes)
}
