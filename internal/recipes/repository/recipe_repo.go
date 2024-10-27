package repository

import (
	"context"
	"database/sql"

	"github.com/sanchayata-jain/food-blog/internal/recipes/models"
)

type RecipeRepo struct {
	Database *sql.DB
}

func NewRecipeRepo(db *sql.DB) *RecipeRepo {
	return &RecipeRepo{db}
}

func (r *RecipeRepo) InsertRecipe(ctx context.Context, recipe *models.Recipe) error {
	query := `
		INSERT INTO recipes (id, title, description, ingredients, instructions)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	_, err := r.Database.ExecContext(ctx, query, recipe.ID, recipe.Title, recipe.Description, recipe.Ingredients, recipe.Instructions)
	if err != nil {
		return err
	}

	return nil
}
