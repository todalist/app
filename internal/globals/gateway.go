package globals

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type TokenUser struct {
	UserId uint
}

type AuthenticationClaims struct {
	TokenUser
	jwt.RegisteredClaims
}

type TokenUserKey struct{}

func SetTokenUser(tokenUser *TokenUser, c fiber.Ctx) {
	c.Locals(TokenUserKey{}, tokenUser)
}

func GetTokenUser(c fiber.Ctx) (*TokenUser, error) {
	t := c.Locals(TokenUserKey{}).(*TokenUser)
	if t == nil {
		return t, errors.New("token user not found")
	}
	return t, nil
}

func MustGetTokenUser(c fiber.Ctx) *TokenUser {
	t, e := GetTokenUser(c)
	if e != nil {
		LOG.Error("no token user found")
		panic(e)
	}
	return t
}

func MustGetTokenUserContext(c fiber.Ctx) context.Context {
	tokenUser := MustGetTokenUser(c)
	return context.WithValue(c.Context(), TokenUserKey{}, tokenUser)
}

func MustGetTokenUserFromContext(c context.Context) *TokenUser {
	v := c.Value(TokenUserKey{})
	if v == nil {
		LOG.Panic("no token user found from given context", zap.Any("ctx", c))
	}
	tokenUser, ok := v.(*TokenUser)
	if !ok {
		LOG.Panic("type error of token user", zap.Any("ctx", c))
	}
	return tokenUser
}
