package users

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s service) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	rows, err := s.db.Query(ctx, fmt.Sprintf("insert into users (username, email, full_name) values ('%s', '%s', '%s') returning username, email, full_name", user.Username, user.Email, user.FullName))
	if err != nil {
		return nil, errors.Wrap(err, "can't insert into db")
	}

	var createdUser model.User
	err = pgxscan.ScanOne(&createdUser, rows)
	if err != nil {
		return nil, errors.Wrap(err, "can't return inserted entity")
	}

	return &createdUser, nil
}
