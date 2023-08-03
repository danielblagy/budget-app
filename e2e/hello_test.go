//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func E2e_Hello(t *testing.T) {
	t.Parallel()

	status, body, errs := fiber.Get("localhost:5000/v1/entries/expense").Bytes()
	require.Equal(t, fiber.StatusUnauthorized, status)
	require.NotNil(t, body)
	require.Empty(t, errs)
}
