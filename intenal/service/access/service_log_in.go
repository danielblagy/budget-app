package access

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrIncorrectPassword = errors.New("password is incorrect")

func (s service) LogIn(ctx context.Context, login *model.Login) error {
	var passwordHash string
	err := pgxscan.Get(ctx, s.db, &passwordHash, fmt.Sprintf("select password_hash from users where username = '%s'", login.Username))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return errors.Wrap(err, "can't get user from db")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(login.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrIncorrectPassword
		}
		return errors.Wrap(err, "can't compare passwords")
	}

	return nil
}
