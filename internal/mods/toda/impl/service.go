package todaImpl

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/repo"
	"go.uber.org/zap"
)

type TodaService struct {
	repo repo.IRepo
}

func (s *TodaService) Get(ctx context.Context, id uint) (*entity.Toda, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	return todaRepo.Get(id)
}

func (s *TodaService) Save(ctx context.Context, form *entity.Toda) (*entity.Toda, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	isCreate := form.Id < 1
	if isCreate {
		form.UserId = tokenUser.UserId
		form.Status = entity.TodaStatusTodo
	} else {
		// check user permission
		_, err := userTodaRepo.First(&dto.UserTodaQuerier{
			UserId: &tokenUser.UserId,
			TodaId: &form.Id,
		})
		if err != nil {
			globals.LOG.Warn("no permission to update toda",
				zap.Any("user", tokenUser),
				zap.Any("form", form),
				zap.Error(err),
			)
			return nil, fiber.ErrForbidden
		}
	}
	if form.Priority == 0 {
		// TODO to support user config
		form.Priority = entity.TodaPriorityLow
	}
	form, err := todaRepo.Save(form)
	if err != nil {
		return nil, err
	}
	if isCreate {
		// init toda with user
		_, err := userTodaRepo.Save(&entity.UserToda{
			UserId: form.UserId,
			TodaId: form.Id,
		})
		if err != nil {
			return nil, err
		}
	}
	return form, nil
}

func (s *TodaService) List(ctx context.Context, querier *dto.TodaQuerier) ([]*entity.Toda, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	querier.UserId = &tokenUser.UserId
	return todaRepo.List(querier)
}

func (s *TodaService) Delete(ctx context.Context, id uint) (uint, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	_, err := userTodaRepo.First(&dto.UserTodaQuerier{
		UserId: &tokenUser.UserId,
		TodaId: &id,
	})
	if err != nil {
		globals.LOG.Warn("no permission to delete toda",
			zap.Any("user", tokenUser),
			zap.Any("id", id),
			zap.Error(err),
		)
		return id, fiber.ErrForbidden
	}
	if err := userTodaRepo.DeleteByTodaId(id); err != nil {
		return id, err
	}
	return todaRepo.Delete(id)
}

func NewTodaService(repo repo.IRepo) *TodaService {
	return &TodaService{
		repo: repo,
	}
}

// func (s *TodaService) fillTodaVO(ctx context.Context, list []*toda.TodaVO) ([]*toda.TodaVO, error) {
// 	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
// 	todaMap := common.ToFieldMap(list, func(t *toda.TodaVO) uint {
// 		return t.Id
// 	})
// 	todaIds := lo.Keys(todaMap)
// 	todaTagVOs, err := todaTagRefRepo.ListTodaTagByTodaIds(todaIds)
// 	if err != nil {
// 		globals.LOG.Error("todaToVO error", zap.Error(err))
// 		return nil, err
// 	}
// 	for _, vo := range todaTagVOs {
// 		t, ok := todaMap[vo.TodaId]
// 		if ok {
// 			t.Tags = append(t.Tags, vo)
// 		}
// 	}
// 	return list, nil
// }
