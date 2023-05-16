package access

import "context"

func (s service) Authenticate(ctx context.Context, token string) (bool, error) {
	username, err := parseJwtToken(token)
	if err != nil {
		return false, err
	}

	exists, err := s.usersService.Exists(ctx, username)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	return true, nil
}
