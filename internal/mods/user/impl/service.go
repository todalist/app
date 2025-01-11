package userImpl

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/repo"
)

type UserService struct {
	repo repo.IRepo
}

func (s *UserService) Get(ctx context.Context, id uint) (*entity.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.Get(id)
}

func (s *UserService) First(ctx context.Context, querier *dto.UserQuerier) (*entity.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.First(querier)
}

func (s *UserService) Save(ctx context.Context, form *entity.User) (*entity.User, error) {
	userRepo := s.repo.GetUserRepo(ctx)
	return userRepo.Save(form)
}

func (s *UserService) List(ctx context.Context, querier *dto.UserQuerier) ([]*entity.User, error) {
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
