package todaTagImpl

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
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

func (s *TodaTagService) Save(ctx context.Context, form *dto.TodaTagSaveDTO) (*vo.UserTodaTagVO, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
	isCreate := form.Id < 1
	var (
		userTodaTagId uint
		userId        uint
	)
	if isCreate {
		form.OwnerUserId = tokenUser.UserId
	} else {
		// check user permission
		userTodaTag, err := userTodaTagRepo.First(&dto.UserTodaTagQuerier{
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
		userTodaTagId = userTodaTag.Id
		userId = userTodaTag.UserId
	}
	_, err := todaTagRepo.Save(&form.TodaTag)
	if err != nil {
		return nil, err
	}
	if isCreate {
		// init todaTag with user
		userTodaTag, err := userTodaTagRepo.Save(&entity.UserTodaTag{
			UserId:    form.OwnerUserId,
			TodaTagId: form.Id,
			PinTop:    form.PinTop,
		})
		if err != nil {
			return nil, err
		}
		userTodaTagId = userTodaTag.Id
		userId = userTodaTag.UserId
	}
	v := &vo.UserTodaTagVO{
		Tag:           &form.TodaTag,
		UserTodaTagId: userTodaTagId,
		UserId:        userId,
		PinTop:        form.PinTop,
	}
	return v, nil
}

func (s *TodaTagService) SaveUserTodaTag(ctx context.Context, form *entity.UserTodaTag) (*entity.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
	form.UserId = tokenUser.UserId
	return userTodaTagRepo.Save(form)
}

func (s *TodaTagService) List(ctx context.Context, querier *dto.ListUserTodaTagQuerier) ([]*vo.UserTodaTagVO, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
	querier.UserId = &tokenUser.UserId
	return todaTagRepo.ListUserTodaTag(querier)
}

func (s *TodaTagService) Delete(ctx context.Context, id uint) (uint, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
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
