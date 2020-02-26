package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lcoutinho/luizalabs-client-api/config"
	"time"
)

type SimpleJwtAuth struct {
}

func (sv *SimpleJwtAuth) GenerateToken(resource string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(config.JWT_TIME_EXPIRE).Unix()
	claims["resource"] = resource
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(config.JWT_KEY))

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", tokenString)
}
func (sv *SimpleJwtAuth) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})
}
