package userTodaImpl

import (
	"context"

	"github.com/todalist/app/internal/mods/userToda"
	"github.com/todalist/app/internal/repo"
)

type UserTodaService struct {
	repo repo.IRepo
}

func (s *UserTodaService) Get(ctx context.Context, id uint) (*userToda.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.Get(id)
}

func (s *UserTodaService) First(ctx context.Context, querier *userToda.UserTodaQuerier) (*userToda.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.First(querier)
}

func (s *UserTodaService) Save(ctx context.Context, form *userToda.UserToda) (*userToda.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.Save(form)
}

func (s *UserTodaService) List(ctx context.Context, querier *userToda.UserTodaQuerier) ([]*userToda.UserToda, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.List(querier)
}

func (s *UserTodaService) Delete(ctx context.Context, id uint) (uint, error) {
	userTodaRepo := s.repo.GetUserTodaRepo(ctx)
	return userTodaRepo.Delete(id)
}

func NewUserTodaService(repo repo.IRepo) *UserTodaService {
	return &UserTodaService{
		repo: repo,
	}
}
