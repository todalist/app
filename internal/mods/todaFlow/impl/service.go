package todaFlowImpl

import (
	"dailydo.fe1.xyz/internal/mods/todaFlow"
	"dailydo.fe1.xyz/internal/repo"
	"context"
)

type TodaFlowService struct {
	repo repo.IRepo
}

func (s *TodaFlowService) Get(ctx context.Context, id uint) (*todaFlow.TodaFlow, error) {
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	return todaFlowRepo.Get(id)
}

func (s *TodaFlowService) Save(ctx context.Context, form *todaFlow.TodaFlow) (*todaFlow.TodaFlow, error) {
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	return todaFlowRepo.Save(form)
}

func (s *TodaFlowService) List(ctx context.Context, querier *todaFlow.TodaFlowQuerier) ([]*todaFlow.TodaFlow, error) {
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	return todaFlowRepo.List(querier)
}

func (s *TodaFlowService) Delete(ctx context.Context, id uint) (uint, error) {
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	return todaFlowRepo.Delete(id)
}

func NewTodaFlowService(repo repo.IRepo) *TodaFlowService {
	return &TodaFlowService{
		repo: repo,
	}
}