package access

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
)

func (s service) Refresh(ctx context.Context, accessToken, refreshToken string) (*model.UserTokens, error) {
	_, ok, err := s.persistentStoreService.Get(ctx, fmt.Sprintf("token-refresh:%s", refreshToken))
	if err != nil {
		return nil, errors.Wrap(err, "can't check if refresh token is blacklisted")
	}
	if ok {
		return nil, ErrNotAuthorized
	}

	username, err := parseJwtToken(refreshToken)
	if err != nil {
		if errors.Is(err, errTokenExpired) {
			return nil, errors.Wrap(ErrNotAuthorized, "refresh token has expired")
		}
		if errors.Is(err, errInvalidToken) {
			return nil, errors.Wrap(ErrNotAuthorized, "refresh token is invalid")
		}
		return nil, err
	}

	err = s.persistentStoreService.Set(ctx, fmt.Sprintf("token-access:%s", accessToken), accessToken, accessTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't blacklist access token")
	}

	err = s.persistentStoreService.Set(ctx, fmt.Sprintf("token-refresh:%s", refreshToken), refreshToken, refreshTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't blacklist refresh token")
	}

	newAccessToken, err := generateJwtToken(username, accessTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate access token")
	}

	newRefreshToken, err := generateJwtToken(username, refreshTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate refresh token")
	}

	return &model.UserTokens{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
