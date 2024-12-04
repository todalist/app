package todaTagImpl

import (
	"dailydo.fe1.xyz/internal/mods/todaTag"
	"dailydo.fe1.xyz/internal/repo"
	"context"
)

type TodaTagService struct {
	repo repo.IRepo
}

func (s *TodaTagService) Get(ctx context.Context, id uint) (*todaTag.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.Get(id)
}

func (s *TodaTagService) First(ctx context.Context, querier *todaTag.TodaTagQuerier) (*todaTag.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.First(querier)
}

func (s *TodaTagService) Save(ctx context.Context, form *todaTag.TodaTag) (*todaTag.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.Save(form)
}

func (s *TodaTagService) List(ctx context.Context, querier *todaTag.TodaTagQuerier) ([]*todaTag.TodaTag, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.List(querier)
}

func (s *TodaTagService) Delete(ctx context.Context, id uint) (uint, error) {
	todaTagRepo := s.repo.GetTodaTagRepo(ctx)
	return todaTagRepo.Delete(id)
}

func NewTodaTagService(repo repo.IRepo) *TodaTagService {
	return &TodaTagService{
		repo: repo,
	}
}