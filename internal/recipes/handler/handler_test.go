package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sanchayata-jain/food-blog/internal/recipes/mocks"
	"github.com/sanchayata-jain/food-blog/internal/recipes/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateRecipe(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  models.Recipe
		setupMocks   func(*mocks.RecipeService)
		expectStatus int
	}{
		{
			name: "successful creation",
			requestBody: models.Recipe{
				Title:        "Test Recipe",
				Description:  "Test Description",
				Ingredients:  "Test Ingredients",
				Instructions: "Test Instructions",
			},
			setupMocks: func(m *mocks.RecipeService) {
				m.On("CreateRecipe", mock.Anything, mock.AnythingOfType("*models.Recipe")).
					Return(nil)
			},
			expectStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockService := new(mocks.RecipeService)
			tt.setupMocks(mockService)
			handler := NewHandler(mockService)

			// Create request
			body, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// Execute
			handler.CreateRecipe(rec, req)

			// Assert
			assert.Equal(t, tt.expectStatus, rec.Code)
		})
	}
}
