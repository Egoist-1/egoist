package integration

import (
	ijwt "7day/webook/internal/web/jwt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func SetToken() (tokenString string, err error) {
	claims := ijwt.UserClaims{
		Id:   123,
		Ssid: "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err = token.SignedString([]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"))
	return tokenString, err
}
