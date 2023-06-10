package access

import (
	"context"
	"errors"
	"fmt"
	"testing"

	cacheMocks "github.com/danielblagy/budget-app/internal/service/cache/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_LogOut(t *testing.T) {
	t.Parallel()

	t.Run("error: can't blacklist access token", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		username := "someuserxxx"

		accessToken, tokenGenerateErr := generateJwtToken(username, accessTokenDuration)
		require.NoError(t, tokenGenerateErr)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On(
				"SetWithExpiration",
				mock.AnythingOfType("*context.emptyCtx"),
				fmt.Sprintf("token-access:%s", accessToken),
				accessToken,
				accessTokenDuration,
			).
			Return(expectedErr)

		service := NewService(nil, cacheService)
		err := service.LogOut(context.Background(), accessToken, "")
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't blacklist access token")
	})

	t.Run("error: can't blacklist refresh token", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		username := "someuserxxx"

		accessToken, tokenGenerateErr := generateJwtToken(username, accessTokenDuration)
		require.NoError(t, tokenGenerateErr)

		refreshToken, tokenGenerateErr := generateJwtToken(username, refreshTokenDuration)
		require.NoError(t, tokenGenerateErr)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On(
				"SetWithExpiration",
				mock.AnythingOfType("*context.emptyCtx"),
				fmt.Sprintf("token-access:%s", accessToken),
				accessToken,
				accessTokenDuration,
			).
			Return(nil)
		cacheService.
			On(
				"SetWithExpiration",
				mock.AnythingOfType("*context.emptyCtx"),
				fmt.Sprintf("token-refresh:%s", refreshToken),
				refreshToken,
				refreshTokenDuration,
			).
			Return(expectedErr)

		service := NewService(nil, cacheService)
		err := service.LogOut(context.Background(), accessToken, refreshToken)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't blacklist refresh token")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		username := "someuserxxx"

		accessToken, tokenGenerateErr := generateJwtToken(username, accessTokenDuration)
		require.NoError(t, tokenGenerateErr)

		refreshToken, tokenGenerateErr := generateJwtToken(username, refreshTokenDuration)
		require.NoError(t, tokenGenerateErr)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On(
				"SetWithExpiration",
				mock.AnythingOfType("*context.emptyCtx"),
				fmt.Sprintf("token-access:%s", accessToken),
				accessToken,
				accessTokenDuration,
			).
			Return(nil)
		cacheService.
			On(
				"SetWithExpiration",
				mock.AnythingOfType("*context.emptyCtx"),
				fmt.Sprintf("token-refresh:%s", refreshToken),
				refreshToken,
				refreshTokenDuration,
			).
			Return(nil)

		service := NewService(nil, cacheService)
		err := service.LogOut(context.Background(), accessToken, refreshToken)
		require.NoError(t, err)
	})
}
