package middleware

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

type jwtMiddleware struct {
	secret string
}

func NewJwtRequestMiddleware(secret string) Middleware {
	return &jwtMiddleware{
		secret: secret,
	}
}

func (middware *jwtMiddleware) WrapHandler(ctx iris.Context) error {
	if ctx.Path() == "/auth" {
		return nil
	}

	requestId := ctx.URLParam("requestId")
	if len(requestId) == 0 {
		return errors.New("request id empty")
	}

	tokenStr := ctx.URLParam("token")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(middware.secret), nil
	})

	if err != nil || !token.Valid {
		return errors.New("token error")
	}
	return nil
}
