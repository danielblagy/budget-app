package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

// Create a new user, log in as the user. Returns created and logged in username.
func SetupUser(ctx context.Context, t *testing.T, client *http.Client) string {
	t.Logf("setting up user")

	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	username := fmt.Sprintf("%s%s", "e2euser", timestampStr[len(timestampStr)-4:])
	email := fmt.Sprintf("%s%s", timestampStr[len(timestampStr)-4:], "@e2email.com")
	password := timestampStr

	status, body := Post(ctx, t, client, "http://localhost:5123/v1/users", model.User{
		Username: username,
		Email:    email,
		FullName: "Test User",
		Password: password,
	})
	require.Equal(t, fiber.StatusOK, status)
	require.NotNil(t, body)

	var user model.User
	err := json.Unmarshal(body, &user)
	require.NoError(t, err)
	require.Equal(t, username, user.Username)
	require.Equal(t, email, user.Email)
	require.Equal(t, "Test User", user.FullName)

	status, body = Post(ctx, t, client, "http://localhost:5123/v1/access/login", model.Login{
		Username: username,
		Password: password,
	})
	require.Equal(t, fiber.StatusOK, status)
	require.NotNil(t, body)

	t.Logf("user is set up successfully")

	return username
}
