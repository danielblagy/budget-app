package access

import (
	"context"
	"errors"
	"testing"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/users"
	usersMocks "github.com/danielblagy/budget-app/internal/service/users/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_LogIn(t *testing.T) {
	t.Parallel()

	username := "testusername123"
	password := "password111"

	t.Run("error: user not found", func(t *testing.T) {
		t.Parallel()

		usersService := new(usersMocks.Service)
		usersService.
			On("GetPasswordHash", mock.AnythingOfType("*context.emptyCtx"), username).
			Return("", users.ErrUserNotFound)

		service := NewService(usersService, nil)
		_, err := service.LogIn(context.Background(), &model.Login{
			Username: username,
			Password: password,
		})
		require.ErrorIs(t, err, ErrUserNotFound)
	})

	t.Run("error: from usersService.GetPasswordHash", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		usersService := new(usersMocks.Service)
		usersService.
			On("GetPasswordHash", mock.AnythingOfType("*context.emptyCtx"), username).
			Return("", expectedErr)

		service := NewService(usersService, nil)
		_, err := service.LogIn(context.Background(), &model.Login{
			Username: username,
			Password: password,
		})
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("error: password is incorrect", func(t *testing.T) {
		t.Parallel()

		//passwordHash := "$2a$10$eB3Axm6ikuREMCwYlxgrgOuEqjxL7r20ZIgaWziIL8JajzuXRQ6HW"
		// password and passwordHash don't match
		passwordHash := "$2a$10$eB3Axm6ikuREMCwYlxgrgOuEqjxL7r20ZIgaxxxxxxxxxxxxxxxxxxx"

		usersService := new(usersMocks.Service)
		usersService.
			On("GetPasswordHash", mock.AnythingOfType("*context.emptyCtx"), username).
			Return(passwordHash, nil)

		service := NewService(usersService, nil)
		_, err := service.LogIn(context.Background(), &model.Login{
			Username: username,
			Password: password,
		})
		require.ErrorIs(t, err, ErrIncorrectPassword)
	})

	t.Run("error: can't compare passwords", func(t *testing.T) {
		t.Parallel()

		//passwordHash := "$2a$10$eB3Axm6ikuREMCwYlxgrgOuEqjxL7r20ZIgaWziIL8JajzuXRQ6HW"
		// password has is not valid (too short)
		passwordHash := "xxx"

		usersService := new(usersMocks.Service)
		usersService.
			On("GetPasswordHash", mock.AnythingOfType("*context.emptyCtx"), username).
			Return(passwordHash, nil)

		service := NewService(usersService, nil)
		_, err := service.LogIn(context.Background(), &model.Login{
			Username: username,
			Password: password,
		})
		require.ErrorContains(t, err, "can't compare passwords")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		passwordHash := "$2a$10$eB3Axm6ikuREMCwYlxgrgOuEqjxL7r20ZIgaWziIL8JajzuXRQ6HW"

		usersService := new(usersMocks.Service)
		usersService.
			On("GetPasswordHash", mock.AnythingOfType("*context.emptyCtx"), username).
			Return(passwordHash, nil)

		service := NewService(usersService, nil)
		userTokens, err := service.LogIn(context.Background(), &model.Login{
			Username: username,
			Password: password,
		})
		require.NoError(t, err)
		require.NotEmpty(t, userTokens.AccessToken)
		require.NotEmpty(t, userTokens.RefreshToken)
	})
}
