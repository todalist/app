package todaImpl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
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

func (s *TodaService) Save(ctx context.Context, form *dto.TodaSaveDTO) (*vo.UserTodaVO, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
	isCreate := form.Id < 1
	if isCreate {
		form.OwnerUserId = tokenUser.UserId
		form.Status = entity.TodaStatusTodo
	} else {
		// check user permission
		err := s.checkIfWritable(ctx, form.Id)
		if err != nil {
			return nil, err
		}
	}
	if form.Priority == 0 {
		// TODO to support user config
		form.Priority = entity.TodaPriorityLow
	}
	toda, err := todaRepo.Save(&form.Toda)
	if err != nil {
		return nil, err
	}
	if isCreate {
		// init toda with user
		_, err := userTodaRepo.Save(&entity.UserToda{
			UserId: form.OwnerUserId,
			TodaId: form.Id,
		})
		if err != nil {
			return nil, err
		}
	}
	userTodaVO := &vo.UserTodaVO{
		TodaVO: &vo.TodaVO{
			Toda: toda,
		},
	}
	s.fillTodaVO(ctx, []*vo.UserTodaVO{userTodaVO})
	return userTodaVO, nil
}

// checkIfWritable checks if the user in the context has permission to write to the specified toda (id).
// If the user has permission, it returns nil, otherwise it returns fiber.ErrForbidden.
func (s *TodaService) checkIfWritable(ctx context.Context, id uint) error {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
	_, err := userTodaRepo.First(&dto.UserTodaQuerier{
		UserId: &tokenUser.UserId,
		TodaId: &id,
	})
	if err != nil {
		globals.LOG.Warn("no permission to update toda",
			zap.Any("user", tokenUser),
			zap.Any("id", id),
			zap.Error(err),
		)
		return fiber.ErrForbidden
	}
	return nil
}

func (s *TodaService) Delete(ctx context.Context, id uint) (uint, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
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

func (s *TodaService) List(ctx context.Context, querier *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)
	querier.UserId = &tokenUser.UserId
	list, err := todaRepo.ListUserToda(querier)
	if err != nil {
		return nil, err
	}
	if err = s.fillTodaVO(ctx, list); err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaService) fillTodaVO(ctx context.Context, list []*vo.UserTodaVO) error {
	if len(list) == 0 {
		return nil
	}
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	todaMap := common.ToFieldMap(list, func(t *vo.UserTodaVO) uint {
		return t.Toda.Id
	})
	todaIds := lo.Keys(todaMap)
	todaTagVOs, err := todaTagRepo.ListTodaTagByTodaIds(todaIds)
	if err != nil {
		globals.LOG.Error("todaToVO error", zap.Error(err))
		return err
	}
	for _, vo := range todaTagVOs {
		t, ok := todaMap[vo.TodaId]
		if ok {
			t.Tags = append(t.Tags, vo)
		}
	}
	return nil
}

// FlowToda updates the status of a Toda entity and records the change in the TodaFlow entity.
// It first verifies if the user has permission to modify the specified Toda by checking writability.
// If writable, it creates a TodaFlow record with the previous and new status, then updates the Toda status.
// Returns the ID of the Toda if successful, or an error if any step fails.
func (s *TodaService) FlowToda(ctx context.Context, form *dto.FlowTodaDTO) (*uint, error) {
	now := time.Now().UTC()
	todaRepo := s.repo.GetTodaRepo(ctx)
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	tokenUser := globals.MustTokenUserFromCtx(ctx)

	toda, err := todaRepo.Get(form.TodaId)
	if err != nil {
		return nil, err
	}
	err = s.checkIfWritable(ctx, form.TodaId)
	if err != nil {
		return nil, err
	}
	todaFlow := &entity.TodaFlow{
		TodaId: toda.Id,
		UserId: tokenUser.UserId,
		Prev:   toda.Status,
		Next:   form.Status,
	}
	_, err = todaFlowRepo.Save(todaFlow)
	if err != nil {
		return nil, err
	}
	toda.Status = form.Status
	if form.Status == entity.TodaStatusFinished {
		toda.CompletedAt = &now
	}
	if _, err := todaRepo.Save(toda); err != nil {
		return nil, err
	}
	return &form.TodaId, nil
}

func NewTodaService(repo repo.IRepo) *TodaService {
	return &TodaService{
		repo: repo,
	}
}
