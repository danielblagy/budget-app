package access

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/users"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrIncorrectPassword = errors.New("password is incorrect")

func (s service) LogIn(ctx context.Context, login *model.Login) (*model.UserTokens, error) {
	passwordHash, err := s.usersService.GetPasswordHash(ctx, login.Username)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(login.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrIncorrectPassword
		}
		return nil, errors.Wrap(err, "can't compare passwords")
	}

	accessToken, err := generateJwtToken(login.Username, accessTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate access token")
	}

	refreshToken, err := generateJwtToken(login.Username, refreshTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate refresh token")
	}

	return &model.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
