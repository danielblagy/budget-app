//go:build e2e
// +build e2e

package e2e

import (
	"context"
	"testing"

	"github.com/danielblagy/budget-app/e2e/entry"
)

func Test_Main(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	entry.RunEntryCreate(ctx, t)
}
