package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey      []byte      = []byte("simple secret key") // должен размещаться .toml/.env файле
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		// здесь может быть дополнительная логика для проверки токена
		return SecretKey, nil
	}
)

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: emptyValidFunc,
	SigningMethod:       jwt.SigningMethodHS256,
})
