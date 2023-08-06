package util

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

// CreateNewCategory creates a new category for the logged in user. Returns category ID.
func CreateNewCategory(ctx context.Context, t *testing.T, client *http.Client, name string, categoryType model.CategoryType) int64 {
	t.Logf("create category name: '%s', type: '%s'", name, categoryType)

	status, body := Post(ctx, t, client, "http://localhost:5000/v1/categories", model.CreateCategory{
		Name: name,
		Type: categoryType,
	})
	require.Equal(t, fiber.StatusOK, status)
	require.NotEmpty(t, body)

	var category model.Category
	err := json.Unmarshal(body, &category)
	require.NoError(t, err)
	require.Equal(t, name, category.Name)
	require.Equal(t, categoryType, category.Type)
	username, ok := ctx.Value("username").(string)
	require.True(t, ok)
	require.Equal(t, username, category.UserID)

	return category.ID
}
