package access

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	persistentStoreMocks "github.com/danielblagy/budget-app/internal/service/persistent-store/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Refresh(t *testing.T) {
	t.Parallel()

	t.Run("error: can't check if refresh token is blacklisted", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		refreshToken, generateTokenErr := generateJwtToken("someuser", refreshTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, false, expectedErr)

		service := NewService(nil, persistentStoreService)
		_, err := service.Refresh(context.Background(), "", refreshToken)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't check if refresh token is blacklisted")
	})

	t.Run("error: refresh token is blacklisted", func(t *testing.T) {
		t.Parallel()

		refreshToken, generateTokenErr := generateJwtToken("someuser", refreshTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, true, nil)

		service := NewService(nil, persistentStoreService)
		_, err := service.Refresh(context.Background(), "", refreshToken)
		require.ErrorIs(t, err, ErrNotAuthorized)
	})

	t.Run("error: refresh token has expired", func(t *testing.T) {
		t.Parallel()

		currentTime := time.Now()
		refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			"someuser",
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(-time.Minute)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		refreshToken, tokenErr := refreshTokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, false, nil)

		service := NewService(nil, persistentStoreService)
		_, err := service.Refresh(context.Background(), "", refreshToken)
		require.ErrorIs(t, err, ErrNotAuthorized)
		require.ErrorContains(t, err, "refresh token has expired")
	})

	t.Run("error: refresh token is invalid", func(t *testing.T) {
		t.Parallel()

		currentTime := time.Now()
		refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			"someuser",
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(-time.Minute)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		refreshToken, tokenErr := refreshTokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)
		refreshToken += "xxxxx"

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, false, nil)

		service := NewService(nil, persistentStoreService)
		_, err := service.Refresh(context.Background(), "", refreshToken)
		require.NotNil(t, err)
		require.ErrorContains(t, err, "can't parse jwt token")
	})

	t.Run("error: can't blacklist access token", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		accessToken, generateTokenErr := generateJwtToken("someuser", accessTokenDuration)
		require.NoError(t, generateTokenErr)
		refreshToken, generateTokenErr := generateJwtToken("someuser", refreshTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, false, nil)
		persistentStoreService.
			On(
				"Set",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("string"),
				accessToken,
				accessTokenDuration,
			).
			Return(expectedErr)

		service := NewService(nil, persistentStoreService)
		_, err := service.Refresh(context.Background(), accessToken, refreshToken)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't blacklist access token")
	})

	t.Run("error: can't blacklist refresh token", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		accessToken, generateTokenErr := generateJwtToken("someuser", accessTokenDuration)
		require.NoError(t, generateTokenErr)
		refreshToken, generateTokenErr := generateJwtToken("someuser", refreshTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, false, nil)
		persistentStoreService.
			On(
				"Set",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("string"),
				accessToken,
				accessTokenDuration,
			).
			Return(nil)
		persistentStoreService.
			On(
				"Set",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("string"),
				refreshToken,
				refreshTokenDuration,
			).
			Return(expectedErr)

		service := NewService(nil, persistentStoreService)
		_, err := service.Refresh(context.Background(), accessToken, refreshToken)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't blacklist refresh token")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		accessToken, generateTokenErr := generateJwtToken("someuser", accessTokenDuration)
		require.NoError(t, generateTokenErr)
		refreshToken, generateTokenErr := generateJwtToken("someuser", refreshTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-refresh:%s", refreshToken)).
			Return(nil, false, nil)
		persistentStoreService.
			On(
				"Set",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("string"),
				accessToken,
				accessTokenDuration,
			).
			Return(nil)
		persistentStoreService.
			On(
				"Set",
				mock.AnythingOfType("*context.emptyCtx"),
				mock.AnythingOfType("string"),
				refreshToken,
				refreshTokenDuration,
			).
			Return(nil)

		service := NewService(nil, persistentStoreService)
		newTokens, err := service.Refresh(context.Background(), accessToken, refreshToken)
		require.NoError(t, err)
		require.NotNil(t, newTokens)
		require.NotEmpty(t, newTokens.AccessToken)
		require.NotEmpty(t, newTokens.RefreshToken)
	})
}
