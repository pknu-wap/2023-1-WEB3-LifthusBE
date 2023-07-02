package helper

import (
	"errors"
	"fmt"
	"lifthus-auth/common/lifthus"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWTWithHMAC(tokenString string) (claims jwt.MapClaims, expired bool, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(lifthus.HusSecretKey), nil
	})
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if token.Valid {
			return claims, false, nil
		} else {
			return claims, true, nil
		}
	} else {
		return nil, false, fmt.Errorf("invalid token")
	}
}
