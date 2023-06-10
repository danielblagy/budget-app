package access

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

var ErrNotAuthorized = errors.New("not authorized")

func (s service) Authorize(ctx context.Context, token string) (string, error) {
	_, ok, err := s.cacheService.Get(ctx, fmt.Sprintf("token-access:%s", token))
	if err != nil {
		return "", errors.Wrap(err, "can't check if token is blacklisted")
	}
	if ok {
		return "", ErrNotAuthorized
	}

	username, err := parseJwtToken(token)
	if err != nil {
		if errors.Is(err, errTokenExpired) {
			return "", errors.Wrap(ErrNotAuthorized, "token has expired")
		}
		if errors.Is(err, errInvalidToken) {
			return "", errors.Wrap(ErrNotAuthorized, "token is invalid")
		}
		return "", err
	}

	exists, err := s.usersService.Exists(ctx, username)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.Wrap(ErrNotAuthorized, "user doesn't exist")
	}

	return username, nil
}
