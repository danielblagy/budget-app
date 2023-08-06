package entry

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/danielblagy/budget-app/e2e/util"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func RunEntryCreate(ctx context.Context, t *testing.T) {
	client := util.SetupHttpClient()

	username := util.SetupUser(ctx, t, client)
	ctx = context.WithValue(ctx, "username", username)
	testInvalidRequest(ctx, t, client)
}

func testInvalidRequest(ctx context.Context, t *testing.T, client *http.Client) {
	categoryID := util.CreateNewCategory(ctx, t, client, "category 1", model.CategoryTypeExpense)
	amount := float64(time.Now().Unix()) / 1000.0
	date := "2023-08-03"
	description := "category 1 description"
	entryType := model.EntryTypeExpense

	t.Logf("category not exists")
	status, body := util.Post(ctx, t, client, "http://localhost:5000/v1/entries", model.CreateEntry{
		CategoryID:  -1,
		Amount:      amount,
		Date:        date,
		Description: description,
		Type:        entryType,
	})
	require.Equal(t, fiber.StatusBadRequest, status)
	require.Equal(t, "category not exists", string(body))

	t.Logf("date format not valid")
	status, body = util.Post(ctx, t, client, "http://localhost:5000/v1/entries", model.CreateEntry{
		CategoryID:  categoryID,
		Amount:      amount,
		Date:        "123-56-457",
		Description: description,
		Type:        entryType,
	})
	require.Equal(t, fiber.StatusBadRequest, status)
	require.Equal(t, "date format not valid", string(body))

	t.Logf("amount less than or equal to zero")
	status, body = util.Post(ctx, t, client, "http://localhost:5000/v1/entries", model.CreateEntry{
		CategoryID:  categoryID,
		Amount:      0.0,
		Date:        date,
		Description: description,
		Type:        entryType,
	})
	require.Equal(t, fiber.StatusBadRequest, status)
	require.Equal(t, "amount less than or equal to zero", string(body))

	t.Logf("entry type is not valid")
	status, body = util.Post(ctx, t, client, "http://localhost:5000/v1/entries", model.CreateEntry{
		CategoryID:  categoryID,
		Amount:      amount,
		Date:        date,
		Description: description,
		Type:        model.EntryType("12345"),
	})
	require.Equal(t, fiber.StatusBadRequest, status)
	require.Equal(t, "entry type is not valid", string(body))
}
