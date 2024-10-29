package services

import (
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(filter *models.User, maskSensitive bool) (user *models.User, err error) {
	user = new(models.User)
	if err = globals.DB.Where(filter).First(user).Error; err != nil {
		return nil, err
	}
	if maskSensitive {
		maskUser(user)
	}
	return user, nil
}

func maskUser(user *models.User) {
	user.Password = ""
}

func PutGetUser(filter *models.User, pwd string) (user *models.User, err error) {
	if user, err = GetUser(filter, false); err == nil {
		return
	}
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return
	}
	username := filter.Username
	if username == "" {
		username = filter.Email
	}
	user = &models.User{
		Username: username,
		Email:    filter.Email,
		Password: string(hash),
	}
	if err = globals.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return
}

func RewritePassword(filter *models.User, prev string, next string) error {
	user, err := GetUser(filter, false)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(prev)); err != nil {
		globals.LOG.Info("password compare failed", zap.Uint("userId", user.ID),
			zap.String("username", filter.Username),
			zap.Error(err),
		)
		return fiber.ErrUnauthorized
	}
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(next), bcrypt.DefaultCost); err != nil {
		return err
	}
	user.Password = string(hash)
	if err = globals.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}
