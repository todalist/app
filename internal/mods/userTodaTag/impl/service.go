package userTodaTagImpl

import (
	"dailydo.fe1.xyz/internal/mods/userTodaTag"
	"dailydo.fe1.xyz/internal/repo"
	"context"
)

type UserTodaTagService struct {
	repo repo.IRepo
}

func (s *UserTodaTagService) Get(ctx context.Context, id uint) (*userTodaTag.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.Get(id)
}

func (s *UserTodaTagService) Save(ctx context.Context, form *userTodaTag.UserTodaTag) (*userTodaTag.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.Save(form)
}

func (s *UserTodaTagService) List(ctx context.Context, querier *userTodaTag.UserTodaTagQuerier) ([]*userTodaTag.UserTodaTag, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.List(querier)
}

func (s *UserTodaTagService) Delete(ctx context.Context, id uint) (uint, error) {
	userTodaTagRepo := s.repo.GetUserTodaTagRepo(ctx)
	return userTodaTagRepo.Delete(id)
}