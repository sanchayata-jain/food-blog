package recipesservice

import (
	"context"
	"errors"

	"github.com/sanchayata-jain/food-blog/internal/recipes/models"
	"github.com/sanchayata-jain/food-blog/internal/recipes/repository"
)

type Service interface {
	CreateRecipe(ctx context.Context, recipe *models.Recipe) error
	GetRecipes(ctx context.Context) ([]models.Recipe, error)
}

type RecipeService struct {
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

func (r *RecipeService) GetRecipes(ctx context.Context) ([]models.Recipe, error) {
	rows, err := r.Repository.GetRecipes(ctx)
	if err != nil {
		return nil, err
	}

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(&recipe.Title, &recipe.Description, &recipe.Ingredients, &recipe.Instructions)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
