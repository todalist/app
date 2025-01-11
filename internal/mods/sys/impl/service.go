package sysImpl

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/repo"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type SysServiceImpl struct {
	repo repo.IRepo
}

func (s *SysServiceImpl) PasswordLogin(ctx context.Context, form *dto.PasswordLoginDTO) (*string, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	u, err := userRepo.First(&dto.UserQuerier{Email: &form.Email})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// create new user
		var hash []byte
		if hash, err = bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost); err != nil {
			globals.LOG.Error("generate hash error", zap.Error(err))
			return nil, fiber.ErrUnauthorized
		}
		u = &entity.User{
			Email:    form.Email,
			Username: form.Email,
			Password: string(hash),
		}
		if u, err = userRepo.Save(u); err != nil {
			globals.LOG.Error("create new user error", zap.Error(err))
			return nil, fiber.ErrUnauthorized
		}
	} else {
		// user exists
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Password)); err != nil {
			globals.LOG.Info("password compare failed", zap.Uint("userId", u.Id),
				zap.String("username", form.Email),
				zap.Error(err),
			)
			return nil, fiber.ErrUnauthorized
		}
	}
	return s.tokenGen(u)
}

func (s *SysServiceImpl) tokenGen(u *entity.User) (*string, error) {
	now := time.Now().UTC()
	jwtConfig := globals.CONF.Auth.Jwt
	claims := &globals.AuthenticationClaims{
		TokenUser: globals.TokenUser{
			UserId: u.Id,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(jwtConfig.JwtExpireSec))),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    jwtConfig.JwtIssuer,
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, e := j.SignedString([]byte(jwtConfig.JwtSecret))
	if e != nil {
		globals.LOG.Error("signing token error", zap.Any("jwt", j), zap.Error(e))
		return nil, fiber.ErrInternalServerError
	}
	return &token, nil
}

func NewSysService(repo repo.IRepo) *SysServiceImpl {
	return &SysServiceImpl{
		repo: repo,
	}
}
