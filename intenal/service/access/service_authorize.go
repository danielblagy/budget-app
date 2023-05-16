package access

import (
	"context"

	"github.com/pkg/errors"
)

var ErrNotAuthorized = errors.New("not authorized")

func (s service) Authorize(ctx context.Context, token string) (string, error) {
	username, err := parseJwtToken(token)
	if err != nil {
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
