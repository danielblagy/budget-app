package entry

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/danielblagy/budget-app/e2e/util"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func RunEntryCreate(ctx context.Context, t *testing.T) {
	// client is needed to store http-only cookies (jwt tokens)
	client := util.SetupHttpClient()

	username := util.SetupUser(ctx, t, client)
	ctx = context.WithValue(ctx, "username", username)
	testInvalidRequest(ctx, t, client)
	testSuccessCreation(ctx, t, client)
}

func testInvalidRequest(ctx context.Context, t *testing.T, client *http.Client) {
	// create a new category
	categoryID := util.CreateNewCategory(ctx, t, client, "category 1", model.CategoryTypeExpense)
	amount := float64(time.Now().Unix()) / 1000.0
	date := "2023-08-03"
	description := "category 1 description"
	entryType := model.EntryTypeExpense

	// negative: cannot create entry because category doesn not exist
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

	// negative: cannot create entry because date is in invalid format
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

	// negative: cannot create entry because amount is not valid
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

	// negative: cannot create entry because entry type is neither 'income' nor 'expense'
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

func testSuccessCreation(ctx context.Context, t *testing.T, client *http.Client) {
	// create an expense category
	category1ID := util.CreateNewCategory(ctx, t, client, "category 2", model.CategoryTypeExpense)
	createEntry1Body := model.CreateEntry{
		CategoryID:  category1ID,
		Amount:      500.0,
		Date:        "2023-08-06",
		Description: "This is an expense on the 6th of August 2023, category 2.",
		Type:        model.EntryTypeExpense,
	}

	// positive: create an expense entry with description
	t.Logf("positive: create an expense entry with description")
	status, body := util.Post(ctx, t, client, "http://localhost:5000/v1/entries", createEntry1Body)
	require.Equal(t, fiber.StatusOK, status)
	require.NotEmpty(t, body)

	username, ok := ctx.Value("username").(string)
	require.True(t, ok)

	var entry1 model.Entry
	err := json.Unmarshal(body, &entry1)
	require.NoError(t, err)

	require.NotEmpty(t, entry1.ID)
	require.Equal(t, username, entry1.UserID)
	require.Equal(t, category1ID, *entry1.CategoryID)
	require.Equal(t, 500.0, entry1.Amount)
	require.Equal(t, 6, entry1.Date.Day())
	require.Equal(t, time.August, entry1.Date.Month())
	require.Equal(t, 2023, entry1.Date.Year())
	require.Equal(t, "This is an expense on the 6th of August 2023, category 2.", entry1.Description)
	require.Equal(t, model.EntryTypeExpense, entry1.Type)

	// create an income category
	category2ID := util.CreateNewCategory(ctx, t, client, "category 3", model.CategoryTypeIncome)
	createEntry2Body := model.CreateEntry{
		CategoryID: category2ID,
		Amount:     1800.0,
		Date:       "2023-08-05",
		Type:       model.EntryTypeIncome,
	}

	// positive: create an income entry without description
	t.Logf("positive: create an expense entry without description")
	status, body = util.Post(ctx, t, client, "http://localhost:5000/v1/entries", createEntry2Body)
	require.Equal(t, fiber.StatusOK, status)
	require.NotEmpty(t, body)

	var entry2 model.Entry
	err = json.Unmarshal(body, &entry2)
	require.NoError(t, err)

	require.NotEmpty(t, entry2.ID)
	require.Equal(t, username, entry2.UserID)
	require.Equal(t, category2ID, *entry2.CategoryID)
	require.Equal(t, 1800.0, entry2.Amount)
	require.Equal(t, 5, entry2.Date.Day())
	require.Equal(t, time.August, entry2.Date.Month())
	require.Equal(t, 2023, entry2.Date.Year())
	require.Empty(t, entry2.Description)
	require.Equal(t, model.EntryTypeIncome, entry2.Type)
}
