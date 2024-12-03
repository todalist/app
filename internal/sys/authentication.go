package sys

// import (
// 	"dailydo.fe1.xyz/internal/common"
// 	"dailydo.fe1.xyz/internal/globals"
// 	"dailydo.fe1.xyz/internal/models"
// 	"dailydo.fe1.xyz/internal/services"
// 	"github.com/gofiber/fiber/v3"
// 	"github.com/golang-jwt/jwt/v5"
// 	"go.uber.org/zap"
// 	"golang.org/x/crypto/bcrypt"
// 	"time"
// )

// type PasswordLoginForm struct {
// 	Email string `json:"email" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// }

// const LOGIN_LIMIT_KEY = "limit/LOGIN_LIMIT_KEY"

// // user login via password
// func PasswordLoginRoute(c fiber.Ctx) error {
// 	params := new(PasswordLoginForm)
// 	c.Bind().Body(params)
// 	if err := common.Valid(params); err != nil {
// 		return fiber.ErrBadRequest
// 	}
// 	user, err := services.PutGetUser(
// 		&models.User{Email: params.Email},
// 		params.Password,
// 	)
// 	if err != nil || user == nil {
// 		globals.LOG.Info("user login error", zap.Any("err", err))
// 		return fiber.ErrUnauthorized
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
// 		globals.LOG.Info("password compare failed", zap.Uint("userId", user.ID),
// 			zap.String("username", params.Email),
// 			zap.Error(err),
// 		)
// 		return fiber.ErrUnauthorized
// 	}
// 	now := time.Now().UTC()
// 	jwtConfig := globals.CONF.Auth.Jwt
// 	claims := &globals.AuthenticationClaims{
// 		TokenUser: globals.TokenUser{
// 			UserID: user.ID,
// 		},
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(jwtConfig.JwtExpireSec))),
// 			IssuedAt:  jwt.NewNumericDate(now),
// 			Issuer:    jwtConfig.JwtIssuer,
// 		},
// 	}
// 	j := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
// 	token, e := j.SignedString([]byte(jwtConfig.JwtSecret))
// 	if e != nil {
// 		globals.LOG.Error("signing token error", zap.Any("jwt", j), zap.Error(e))
// 		return fiber.ErrInternalServerError
// 	}
// 	return c.JSON(common.Ok(token))
// }
