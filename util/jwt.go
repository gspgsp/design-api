package util

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SIGNED_STRING = "design_api"

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(id, username string) (string, error) {

	expireTime := time.Now().Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte("testqhsjwt"))
}
