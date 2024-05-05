package util

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"

	"github.com/jinzhu/now"
)

type Claims struct {
	jwt.RegisteredClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(jwtSecret []byte, subject string, id string, expireDays int, issuer string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.AddDate(0, 0, expireDays)

	claims := Claims{
		jwt.RegisteredClaims{
			Subject:   subject,
			ID:        id,
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// GenerateUserToken generate tokens used for auth
func GenerateUserToken(jwtSecret []byte, authType string, clientId int64, userId int64, loginId int64, expireDays int) (string, error) {
	expireTime := now.With(time.Now()).AddDate(0, 0, expireDays)

	claims := ApiClaims{
		clientId,
		userId,
		loginId,
		jwt.RegisteredClaims{
			ID:        strconv.FormatInt(userId, 10),
			Subject:   authType,
			ExpiresAt: jwt.NewNumericDate(expireTime),
			//Issuer:    conf.GetString("auth.issuer"),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string, jwtSecret []byte) (ApiClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &ApiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*ApiClaims); ok && tokenClaims.Valid {
			return *claims, nil
		}
	}

	return ApiClaims{}, err
}
