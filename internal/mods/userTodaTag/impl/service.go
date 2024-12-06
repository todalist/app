package userTodaTagImpl

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/repo"
)

type UserTodaTagService struct {
	repo repo.IRepo
}

func (s *UserTodaTagService) Get(ctx context.Context, id uint) (*entity.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.Get(id)
}

func (s *UserTodaTagService) First(ctx context.Context, querier *dto.UserTodaTagQuerier) (*entity.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.First(querier)
}

func (s *UserTodaTagService) Save(ctx context.Context, form *entity.UserTodaTag) (*entity.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.Save(form)
}

func (s *UserTodaTagService) List(ctx context.Context, querier *dto.UserTodaTagQuerier) ([]*entity.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.List(querier)
}

func (s *UserTodaTagService) Delete(ctx context.Context, id uint) (uint, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.Delete(id)
}

func NewUserTodaTagService(repo repo.IRepo) *UserTodaTagService {
	return &UserTodaTagService{
		repo: repo,
	}
}
