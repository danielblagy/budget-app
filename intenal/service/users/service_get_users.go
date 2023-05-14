package users

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s service) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := pgxscan.Select(ctx, s.db, &users, "select username, email, full_name from users")
	if err != nil {
		return nil, errors.Wrap(err, "can't get users from db")
	}

	return users, nil
}
