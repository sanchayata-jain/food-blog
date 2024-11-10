package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.opencensus.io/plugin/ochttp"

	"github.com/sanchayata-jain/food-blog/internal/recipes/handler"
	"github.com/sanchayata-jain/food-blog/internal/recipes/repository"
	recipeservice "github.com/sanchayata-jain/food-blog/internal/recipes/service"
	"github.com/sanchayata-jain/food-blog/internal/storage"

	_ "github.com/lib/pq"
)

const (
	port string = "8080"
	addr string = "0.0.0.0:8080"
)

type Server interface {
	ListenandServe() error
	Shutdown(ctx context.Context) error
}

type server struct {
	serv *http.Server
}

func New() Server {
	db, err := storage.NewDatabase()
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	err = storage.CreateExtenstion(db)
	if err != nil {
		log.Fatalf("failed to create extenstion: %v", err)
	}
	err = storage.CreateRecipesTable(db)
	if err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}

	r := chi.NewRouter()
	recipeRepo := repository.NewRecipeRepo(db.DB)
	recipeService := recipeservice.NewRecipeService(recipeRepo)
	recipeHandler := handler.NewHandler(recipeService)
	r.MethodFunc(http.MethodGet, "/get-all-recipes", recipeHandler.GetRecipes)
	r.MethodFunc(http.MethodPost, "/create-recipe", recipeHandler.CreateRecipe)

	return &server{
		serv: &http.Server{
			Addr:    addr,
			Handler: r,
		},
	}

}

func (s *server) ListenandServe() error {
	log.Print("Server is running on port 8080")
	return http.ListenAndServe(addr, &ochttp.Handler{
		Handler: s.serv.Handler,
	})
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}
