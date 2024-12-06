package userTodaImpl

import (
	"context"

	"github.com/samber/lo"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
	"github.com/todalist/app/internal/repo"
	"go.uber.org/zap"
)

type UserTodaService struct {
	repo repo.IRepo
}

func (s *UserTodaService) ListUserToda(ctx context.Context, querier *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	querier.UserId = &tokenUser.UserId
	list, err := userTodaRepo.ListUserToda(querier)
	if err != nil {
		return nil, err
	}
	if err = s.fillTodaVO(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserTodaService) fillTodaVO(ctx context.Context, list *[]*vo.UserTodaVO) error {

	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
	todaMap := common.ToFieldMap(list, func(t *vo.UserTodaVO) uint {
		return t.TodaVO.Id
	})
	todaIds := lo.Keys(todaMap)
	todaTagVOs, err := todaTagRefRepo.ListTodaTagByTodaIds(todaIds)
	if err != nil {
		globals.LOG.Error("todaToVO error", zap.Error(err))
		return err
	}
	for _, vo := range todaTagVOs {
		t, ok := todaMap[vo.TodaId]
		if ok {
			t.TodaVO.Tags = append(t.TodaVO.Tags, vo)
		}
	}
	return nil
}

func (s *UserTodaService) CreateUserToda(ctx context.Context, form *entity.Toda) (*entity.Toda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	todaRepo := s.repo.GetTodaRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	form.UserId = tokenUser.UserId
	form.Status = entity.TodaStatusTodo
	if form.Priority == 0 {
		// TODO to support user config
		form.Priority = entity.TodaPriorityLow
	}
	form, err := todaRepo.Save(form)
	if err != nil {
		return nil, err
	}
	// init toda with user
	_, err = userTodaRepo.Save(&entity.UserToda{
		UserId: form.UserId,
		TodaId: form.Id,
	})
	if err != nil {
		return nil, err
	}
	return form, nil
}

func NewUserTodaService(repo repo.IRepo) *UserTodaService {
	return &UserTodaService{
		repo: repo,
	}
}
