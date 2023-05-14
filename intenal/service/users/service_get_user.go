package users

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s service) GetUser(ctx context.Context, userID int) (*model.User, error) {
	var user model.User
	err := pgxscan.Get(ctx, s.db, &user, fmt.Sprintf("select username, email, full_name from users where id = %d", userID))
	if err != nil {
		return nil, errors.Wrap(err, "can't get user from db")
	}

	return &user, nil
}
