package recipesservice

import (
	"context"
	"errors"

	"github.com/sanchayata-jain/food-blog/internal/recipes/models"
	"github.com/sanchayata-jain/food-blog/internal/recipes/repository"
)

type Service interface {
	CreateRecipe(ctx context.Context, recipe *models.Recipe) error
	// Add other methods as needed
}

// RecipeService struct implements the RecipeService interface.
type RecipeService struct {
	// Database *storage.Database
	Repository *repository.RecipeRepo
}

func NewRecipeService(repo *repository.RecipeRepo) *RecipeService {
	return &RecipeService{
		Repository: repo,
	}
}

func (r *RecipeService) CreateRecipe(ctx context.Context, recipe *models.Recipe) error {
	if recipe.ID == "" || recipe.Description == "" || recipe.Ingredients == "" || recipe.Instructions == "" || recipe.Title == "" {
		return errors.New("empty fields")
	}

	err := r.Repository.InsertRecipe(ctx, recipe)
	if err != nil {
		return err
	}

	return nil
}
