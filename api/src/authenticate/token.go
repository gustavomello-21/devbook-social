package authenticate

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func Generate(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	claims["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	assignedToken, err := token.SigningString()
	if err != nil {
		return "", err
	}
	return assignedToken, nil
}
