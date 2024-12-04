package todaImpl

import (
	"dailydo.fe1.xyz/internal/mods/toda"
	"dailydo.fe1.xyz/internal/repo"
	"context"
)

type TodaService struct {
	repo repo.IRepo
}

func (s *TodaService) Get(ctx context.Context, id uint) (*toda.Toda, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	return todaRepo.Get(id)
}

func (s *TodaService) Save(ctx context.Context, form *toda.Toda) (*toda.Toda, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	return todaRepo.Save(form)
}

func (s *TodaService) List(ctx context.Context, querier *toda.TodaQuerier) ([]*toda.Toda, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	return todaRepo.List(querier)
}

func (s *TodaService) Delete(ctx context.Context, id uint) (uint, error) {
	todaRepo := s.repo.GetTodaRepo(ctx)
	return todaRepo.Delete(id)
}

func NewTodaService(repo repo.IRepo) *TodaService {
	return &TodaService{
		repo: repo,
	}
}