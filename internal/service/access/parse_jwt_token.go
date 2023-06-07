package access

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var errInvalidToken = errors.New("token is invalid")

// returns username
func parseJwtToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(jwtSecretKeyEnvVariable)), nil
	})
	if err != nil {
		return "", errors.Wrap(err, "can't parse jwt token")
	}

	if claims, ok := token.Claims.(*tokenClaims); !ok || !token.Valid {
		return "", errInvalidToken
	} else {
		return claims.Username, nil
	}
}
