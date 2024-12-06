package todaTagRefImpl

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/repo"
)

type TodaTagRefService struct {
	repo repo.IRepo
}

func (s *TodaTagRefService) Get(ctx context.Context, id uint) (*entity.TodaTagRef, error) {
	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
	return todaTagRefRepo.Get(id)
}

func (s *TodaTagRefService) First(ctx context.Context, querier *dto.TodaTagRefQuerier) (*entity.TodaTagRef, error) {
	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
	return todaTagRefRepo.First(querier)
}

func (s *TodaTagRefService) Save(ctx context.Context, form *entity.TodaTagRef) (*entity.TodaTagRef, error) {
	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
	return todaTagRefRepo.Save(form)
}

func (s *TodaTagRefService) List(ctx context.Context, querier *dto.TodaTagRefQuerier) ([]*entity.TodaTagRef, error) {
	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
	return todaTagRefRepo.List(querier)
}

func (s *TodaTagRefService) Delete(ctx context.Context, id uint) (uint, error) {
	todaTagRefRepo := s.repo.GetTodaTagRefRepo(ctx)
	return todaTagRefRepo.Delete(id)
}

func NewTodaTagRefService(repo repo.IRepo) *TodaTagRefService {
	return &TodaTagRefService{
		repo: repo,
	}
}