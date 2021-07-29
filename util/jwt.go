package util

import (
	"design-api/common/env"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
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
func GenerateToken(id int64) (interface{}, int) {
	expireTime := time.Now().Add(EXPIRE_TIME)
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenMap := make(map[string]interface{})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", env.ERROR_AUTH_TOKEN
	}

	tokenMap["access_token"] = signedToken
	tokenMap["expire_at"] = strconv.FormatInt(claims.ExpiresAt, 10)

	return tokenMap, env.RESPONSE_SUCCESS
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
func RefreshToken(token string) (interface{}, int) {
	claims, code := ParseToken(token)

	tokenMap := make(map[string]interface{})

	if code != env.SUCCESS {
		return tokenMap, code
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	//claims.ExpiresAt = time.Now().Add(3 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := refToken.SignedString(jwtSecret)
	if err != nil {
		return tokenMap, env.ERROR_AUTH_TOKEN
	}

	tokenMap["access_token"] = newToken
	tokenMap["expire_at"] = strconv.FormatInt(claims.ExpiresAt, 10)

	return tokenMap, env.RESPONSE_SUCCESS
}
