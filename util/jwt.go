package util

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"time"
	"design-api/common/env"
)

var jwtSecret = []byte("design")

const EXPIRE_TIME = 360 * time.Hour

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

/**
生成token
 */
func GenerateToken(id int64) (string, int) {
	expireTime := time.Now().Add(EXPIRE_TIME)
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", env.ERROR_AUTH_TOKEN
	}

	return signedToken, env.RESPONSE_SUCCESS
}

/**
解析token
 */
func ParseToken(token string) (*Claims, int) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {

		if JwtExpireValidReg(err.Error()) {
			return nil, env.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}

		return nil, env.ERROR_AUTH_PARSE
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok {
		return nil, env.ERROR_AUTH_PARSE
	}

	if err := tokenClaims.Valid; err == false {

		return nil, env.ERROR_AUTH_VALID
	}

	return claims, env.SUCCESS
}

/**
刷新token
 */
func RefreshToken(token string) (string, int) {
	claims, code := ParseToken(token)

	if code != env.SUCCESS {
		return "", code
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	//claims.ExpiresAt = time.Now().Add(3 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := refToken.SignedString(jwtSecret)
	if err != nil {
		return "", env.ERROR_AUTH_TOKEN
	}

	return newToken, env.SUCCESS
}
