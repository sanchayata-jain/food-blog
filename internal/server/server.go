package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.opencensus.io/plugin/ochttp"

	"github.com/sanchayata-jain/food-blog/internal/recipes/handler"
	recipe_service "github.com/sanchayata-jain/food-blog/internal/recipes/recipes_service"
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
	// Initialize the services
	recipeService := recipe_service.NewRecipeService(db)
	recipeHandler := handler.NewHandler(recipeService)
	r.MethodFunc(http.MethodPost, "/create-recipe", recipeHandler.CreateRecipe)

	// gatewayRepo := repo.NewRepository(db)
	// gatewayService := service.NewService(gatewayRepo)
	// gatewayCtrl := handler.NewController(gatewayService)

	// r.MethodFunc(http.MethodPost, "/create-payment-request", gatewayCtrl.CreatePaymentRequest())
	// r.MethodFunc(http.MethodGet, "/get-payment-details/{id}", gatewayCtrl.GetPaymentDetails())

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
