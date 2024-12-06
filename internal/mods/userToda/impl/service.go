package userTodaImpl

import (
	"context"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
	"github.com/todalist/app/internal/repo"
)

type UserTodaService struct {
	repo repo.IRepo
}

func (s *UserTodaService) Get(ctx context.Context, id uint) (*entity.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.Get(id)
}

func (s *UserTodaService) First(ctx context.Context, querier *dto.UserTodaQuerier) (*entity.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.First(querier)
}

func (s *UserTodaService) Save(ctx context.Context, form *entity.UserToda) (*entity.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.Save(form)
}

func (s *UserTodaService) List(ctx context.Context, querier *dto.UserTodaQuerier) ([]*entity.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.List(querier)
}

func (s *UserTodaService) Delete(ctx context.Context, id uint) (uint, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.Delete(id)
}

func (s *UserTodaService) ListUserToda(ctx context.Context, querier *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	tokenUser := globals.MustGetTokenUserFromContext(ctx)
	querier.UserId = &tokenUser.UserId
	return userTodaRepo.ListUserToda(querier)
}

func NewUserTodaService(repo repo.IRepo) *UserTodaService {
	return &UserTodaService{
		repo: repo,
	}
}
