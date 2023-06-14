package access

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (s service) LogOut(ctx context.Context, accessToken, refreshToken string) error {
	err := s.persistentStoreService.Set(ctx, fmt.Sprintf("token-access:%s", accessToken), accessToken, accessTokenDuration)
	if err != nil {
		return errors.Wrap(err, "can't blacklist access token")
	}

	err = s.persistentStoreService.Set(ctx, fmt.Sprintf("token-refresh:%s", refreshToken), refreshToken, refreshTokenDuration)
	if err != nil {
		return errors.Wrap(err, "can't blacklist refresh token")
	}

	return nil
}
