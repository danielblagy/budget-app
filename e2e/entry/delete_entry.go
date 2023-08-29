package entry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/danielblagy/budget-app/e2e/common"
	"github.com/danielblagy/budget-app/e2e/util"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func RunEntryDelete(ctx1 context.Context, ctx2 context.Context, t *testing.T) {
	client1 := util.SetupHttpClient()
	username1 := util.SetupUser(ctx1, t, client1)
	ctx1 = context.WithValue(ctx1, common.CtxKeyUsername, username1)
	client2 := util.SetupHttpClient()
	username2 := util.SetupUser(ctx2, t, client2)
	ctx2 = context.WithValue(ctx2, common.CtxKeyUsername, username2)
	testInvalidDeleteRequest(ctx1, ctx2, t, client1, client2)
	testSuccessDeleting(ctx1, t, client1)
}

func testInvalidDeleteRequest(ctx1 context.Context, ctx2 context.Context, t *testing.T, client1 *http.Client, client2 *http.Client) {

	// making 2 entries for 2 users
	categoryID1 := util.CreateNewCategory(ctx1, t, client1, "category 1", model.CategoryTypeExpense)

	util.Post(ctx1, t, client1, "http://localhost:5123/v1/entries", model.CreateEntry{
		CategoryID:  categoryID1,
		Amount:      1000.0,
		Date:        "2023-01-01",
		Description: "first users description",
		Type:        model.EntryTypeExpense,
	})

	categoryID2 := util.CreateNewCategory(ctx2, t, client2, "category 2", model.CategoryTypeExpense)

	status, body2 := util.Post(ctx2, t, client2, "http://localhost:5123/v1/entries", model.CreateEntry{
		CategoryID:  categoryID2,
		Amount:      1000.0,
		Date:        "2023-03-03",
		Description: "second users description",
		Type:        model.EntryTypeExpense,
	})
	require.Equal(t, fiber.StatusOK, status)

	var entry2 model.Entry
	err := json.Unmarshal(body2, &entry2)
	require.NoError(t, err)

	entryID2 := entry2.ID
	t.Logf("entryID2 %d", entryID2)

	t.Logf("User tries to delete other users entry")
	status, body := util.Delete(ctx1, t, client1, fmt.Sprintf("http://localhost:5123/v1/entries/%d", entryID2))
	require.Equal(t, fiber.StatusNotFound, status)
	require.Equal(t, "entry not found", string(body))

	t.Logf("User tries to delete entry that doesnt exist")
	status, body = util.Delete(ctx1, t, client1, "http://localhost:5123/v1/entries/99999")
	require.Equal(t, fiber.StatusNotFound, status)
	require.Equal(t, "entry not found", string(body))

}

func testSuccessDeleting(ctx context.Context, t *testing.T, client *http.Client) {
	username := util.FromCtx(ctx, t, common.CtxKeyUsername)
	categoryID := util.CreateNewCategory(ctx, t, client, "category 3", model.CategoryTypeExpense)
	_, body := util.Post(ctx, t, client, "http://localhost:5123/v1/entries", model.CreateEntry{
		CategoryID:  categoryID,
		Amount:      300.0,
		Date:        "2023-01-01",
		Description: "first users description",
		Type:        model.EntryTypeExpense,
	})

	var entry model.Entry
	json.Unmarshal(body, &entry)
	entryID := entry.ID

	t.Logf("positive: delete an entry")
	status, body := util.Delete(ctx, t, client, fmt.Sprintf("http://localhost:5123/v1/entries/%d", entryID))
	require.Equal(t, fiber.StatusOK, status)
	require.NotEmpty(t, body)

	json.Unmarshal(body, &entry)

	require.NotEmpty(t, entry.ID)
	require.Equal(t, username, entry.UserID)
	require.Equal(t, categoryID, *entry.CategoryID)
	require.Equal(t, 300.0, entry.Amount)
	require.Equal(t, 1, entry.Date.Day())
	require.Equal(t, time.January, entry.Date.Month())
	require.Equal(t, 2023, entry.Date.Year())
	require.Equal(t, "first users description", entry.Description)
	require.Equal(t, model.EntryTypeExpense, entry.Type)

	//check if the entry exists
	status, _ = util.Delete(ctx, t, client, fmt.Sprintf("http://localhost:5123/v1/entries/%d", entryID))
	require.Equal(t, fiber.StatusNotFound, status)

}
