package recipesservice

import (
	"context"
	"errors"

	"github.com/sanchayata-jain/food-blog/internal/recipes/models"
	"github.com/sanchayata-jain/food-blog/internal/storage"
)

// RecipeService struct implements the RecipeService interface.
type RecipeService struct {
	Database *storage.Database
}

func NewRecipeService(database *storage.Database) *RecipeService {
	return &RecipeService{
		Database: database,
	}
}

func (r *RecipeService) CreateRecipe(ctx context.Context, recipe *models.Recipe) error {
	if recipe.ID == "" || recipe.Description == "" || recipe.Ingredients == "" || recipe.Instructions == "" || recipe.Title == "" {
		return errors.New("empty fields")
	}
	query := `
		INSERT INTO recipes (id, title, description, ingredients, instructions)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	_, err := r.Database.DB.ExecContext(ctx, query, recipe.ID, recipe.Title, recipe.Description, recipe.Ingredients, recipe.Instructions)
	if err != nil {
		return err
	}

	return nil
}
