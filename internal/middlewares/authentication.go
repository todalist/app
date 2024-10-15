package middlewares

import (
	"strings"
	"time"
	"dailydo.fe1.xyz/internal/globals"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func AuthenticationMiddleware(conf *globals.AuthenticationConfig) fiber.Handler {
	jwtConfig := conf.Jwt
	whiteList := conf.WhiteList
	whiteListSet := map[string]string{}
	whiteListMather := []string{}
	for _, l := range whiteList {
		if match, ok := strings.CutSuffix(l, "**"); ok {
			whiteListMather = append(whiteListMather, match)
		} else {
			whiteListSet[l] = l
		}
	}
	return func(c fiber.Ctx) (err error) {
		// bearer
		bearer := c.Get(fiber.HeaderAuthorization)
		if bearer == "" {
			path := string(c.Request().URI().Path())
			if _, ok := whiteListSet[path]; ok {
				// skipped
				return c.Next()
			}
			// may need optimizing when list is growing large
			for _, match := range whiteListMather {
				if strings.HasPrefix(path, match) {
					// skipped
					return c.Next()
				}
			}
			return fiber.ErrUnauthorized
		}
		splits := strings.Split(bearer, " ")
		if len(splits) != 2 {
			globals.LOG.Info("invalid bearer", zap.String("bearer", bearer))
			return fiber.ErrUnauthorized
		}
		tokenStr := splits[1]
		var token *jwt.Token
		if token, err = jwt.ParseWithClaims(tokenStr, new(globals.AuthenticationClaims),
			func(t *jwt.Token) (interface{}, error) { return []byte(jwtConfig.JwtSecret), nil },
			jwt.WithExpirationRequired(),
			jwt.WithIssuedAt(),
			jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Name}),
			jwt.WithStrictDecoding(),
			jwt.WithTimeFunc(func() time.Time { return time.Now().UTC() }),
			jwt.WithIssuer(jwtConfig.JwtIssuer),
		); err != nil {
			globals.LOG.Info("jwt token parse", zap.Error(err))
			return fiber.ErrUnauthorized
		}
		tokenUser := token.Claims.(*globals.AuthenticationClaims)
		globals.SetTokenUser(&tokenUser.TokenUser, c)
		return c.Next()
	}
}
