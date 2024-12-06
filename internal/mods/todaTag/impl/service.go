package todaTagImpl

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/repo"
	"go.uber.org/zap"
)

type TodaTagService struct {
	repo repo.IRepo
}

func (s *TodaTagService) Get(ctx context.Context, id uint) (*entity.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.Get(id)
}

func (s *TodaTagService) First(ctx context.Context, querier *dto.TodaTagQuerier) (*entity.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.First(querier)
}

func (s *TodaTagService) Save(ctx context.Context, form *entity.TodaTag) (*entity.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	isCreate := form.Id < 1
	if isCreate {
		form.OwnerUserId = tokenUser.UserId
	} else {
		// check user permission
		_, err := userTodaTagRepo.First(&dto.UserTodaTagQuerier{
			UserId:    &tokenUser.UserId,
			TodaTagId: &form.Id,
		})
		if err != nil {
			globals.LOG.Warn("no permission to update todaTag",
				zap.Any("user", tokenUser),
				zap.Any("form", form),
				zap.Error(err),
			)
			return nil, fiber.ErrForbidden
		}
	}
	form, err := todaTagRepo.Save(form)
	if err != nil {
		return nil, err
	}
	if isCreate {
		// init todaTag with user
		_, err := userTodaTagRepo.Save(&entity.UserTodaTag{
			UserId:    form.OwnerUserId,
			TodaTagId: form.Id,
		})
		if err != nil {
			return nil, err
		}
	}
	return form, nil
}

func (s *TodaTagService) List(ctx context.Context, querier *dto.TodaTagQuerier) ([]*entity.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	querier.OwnerUserId = &tokenUser.UserId
	return todaTagRepo.List(querier)
}

func (s *TodaTagService) Delete(ctx context.Context, id uint) (uint, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	// check user permission
	_, err := userTodaTagRepo.First(&dto.UserTodaTagQuerier{
		UserId:    &tokenUser.UserId,
		TodaTagId: &id,
	})
	if err != nil {
		globals.LOG.Warn("no permission to delete todaTag",
			zap.Any("user", tokenUser),
			zap.Any("id", id),
			zap.Error(err),
		)
		return 0, fiber.ErrForbidden
	}
	return todaTagRepo.Delete(id)
}

func NewTodaTagService(repo repo.IRepo) *TodaTagService {
	return &TodaTagService{
		repo: repo,
	}
}
