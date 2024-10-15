package globals

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type TokenUser struct {
	UserId uint
}

type AuthenticationClaims struct {
	TokenUser
	jwt.RegisteredClaims
}

const (
	_USER_LOCAL_KEY = "TOKEN-USER"
)

func SetTokenUser(tokenUser *TokenUser, c fiber.Ctx) {
	c.Locals(_USER_LOCAL_KEY, tokenUser)
}

func GetTokenUser(c fiber.Ctx) (*TokenUser, error) {
	t := c.Locals(_USER_LOCAL_KEY).(*TokenUser)
	if t == nil {
		return t, errors.New("token user not found")
	}
	return t, nil
}

func MustGetTokenUser(c fiber.Ctx) *TokenUser {
	t, e := GetTokenUser(c)
	if e != nil {
		LOG.Fatal("no token user found")
	}
	return t
}
