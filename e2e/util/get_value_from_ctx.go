package util

import (
	"context"
	"testing"

	"github.com/danielblagy/budget-app/e2e/common"
	"github.com/stretchr/testify/require"
)

func FromCtx(ctx context.Context, t *testing.T, key common.CtxKey) string {
	value, ok := ctx.Value(key).(string)
	require.True(t, ok)

	return value
}
