package access

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	persistentStoreMocks "github.com/danielblagy/budget-app/internal/service/persistent-store/mocks"
	usersMocks "github.com/danielblagy/budget-app/internal/service/users/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Authorize(t *testing.T) {
	t.Parallel()

	t.Run("error: can't check if token is blacklisted", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		token, generateTokenErr := generateJwtToken("someuser", accessTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, false, expectedErr)

		service := NewService(nil, persistentStoreService)
		_, err := service.Authorize(context.Background(), token)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't check if token is blacklisted")
	})

	t.Run("error: token is blacklisted", func(t *testing.T) {
		t.Parallel()

		token, generateTokenErr := generateJwtToken("someuser", accessTokenDuration)
		require.NoError(t, generateTokenErr)

		persistentStoreService := new(persistentStoreMocks.Service)
		persistentStoreService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, true, nil)

		service := NewService(nil, persistentStoreService)
		_, err := service.Authorize(context.Background(), token)
		require.ErrorIs(t, err, ErrNotAuthorized)
	})

	t.Run("error: token has expired", func(t *testing.T) {
		t.Parallel()

		currentTime := time.Now()
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			"someuser",
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(-time.Minute)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		token, tokenErr := tokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)

		persistentStoreMocks := new(persistentStoreMocks.Service)
		persistentStoreMocks.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, false, nil)

		service := NewService(nil, persistentStoreMocks)
		_, err := service.Authorize(context.Background(), token)
		require.ErrorIs(t, err, ErrNotAuthorized)
		require.ErrorContains(t, err, "token has expired")
	})

	t.Run("error: token is invalid", func(t *testing.T) {
		t.Parallel()

		currentTime := time.Now()
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			"someuser",
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(-time.Minute)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		token, tokenErr := tokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)
		token += "xxxxx"

		persistentStoreMocks := new(persistentStoreMocks.Service)
		persistentStoreMocks.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, false, nil)

		service := NewService(nil, persistentStoreMocks)
		_, err := service.Authorize(context.Background(), token)
		require.NotNil(t, err)
		require.ErrorContains(t, err, "can't parse jwt token")
	})

	t.Run("error: from usersService.Exists", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		username := "someusernamexxx"

		currentTime := time.Now()
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			username,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(accessTokenDuration)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		token, tokenErr := tokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)

		persistentStoreMocks := new(persistentStoreMocks.Service)
		persistentStoreMocks.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, false, nil)

		usersService := new(usersMocks.Service)
		usersService.
			On("Exists", mock.AnythingOfType("*context.emptyCtx"), username).
			Return(false, expectedErr)

		service := NewService(usersService, persistentStoreMocks)
		_, err := service.Authorize(context.Background(), token)
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("error: user doesn't exist", func(t *testing.T) {
		t.Parallel()

		username := "someusernamexxx"

		currentTime := time.Now()
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			username,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(accessTokenDuration)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		token, tokenErr := tokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)

		persistentStoreMocks := new(persistentStoreMocks.Service)
		persistentStoreMocks.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, false, nil)

		usersService := new(usersMocks.Service)
		usersService.
			On("Exists", mock.AnythingOfType("*context.emptyCtx"), username).
			Return(false, nil)

		service := NewService(usersService, persistentStoreMocks)
		_, err := service.Authorize(context.Background(), token)
		require.ErrorIs(t, err, ErrNotAuthorized)
		require.ErrorContains(t, err, "user doesn't exist")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		username := "someusernamexxx"

		currentTime := time.Now()
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
			username,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(currentTime.Add(accessTokenDuration)),
				IssuedAt:  jwt.NewNumericDate(currentTime),
				NotBefore: jwt.NewNumericDate(currentTime),
			},
		})
		token, tokenErr := tokenObj.SignedString([]byte(os.Getenv(jwtSecretKeyEnvVariable)))
		require.NoError(t, tokenErr)

		persistentStoreMocks := new(persistentStoreMocks.Service)
		persistentStoreMocks.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("token-access:%s", token)).
			Return(nil, false, nil)

		usersService := new(usersMocks.Service)
		usersService.
			On("Exists", mock.AnythingOfType("*context.emptyCtx"), username).
			Return(true, nil)

		service := NewService(usersService, persistentStoreMocks)
		resultingUsername, err := service.Authorize(context.Background(), token)
		require.NoError(t, err)
		require.Equal(t, username, resultingUsername)
	})
}
