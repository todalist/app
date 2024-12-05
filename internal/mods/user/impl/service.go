package userImpl

import (
	"context"

	"github.com/todalist/app/internal/mods/user"
	"github.com/todalist/app/internal/repo"
)

type UserService struct {
	repo repo.IRepo
}

func (s *UserService) Get(ctx context.Context, id uint) (*user.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.Get(id)
}

func (s *UserService) First(ctx context.Context, querier *user.UserQuerier) (*user.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.First(querier)
}

func (s *UserService) Save(ctx context.Context, form *user.User) (*user.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.Save(form)
}

func (s *UserService) List(ctx context.Context, querier *user.UserQuerier) ([]*user.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.List(querier)
}

func (s *UserService) Delete(ctx context.Context, id uint) (uint, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.Delete(id)
}

func NewUserService(repo repo.IRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
