package todaFlowImpl

import (
	"context"

	"github.com/todalist/app/internal/mods/todaFlow"
	"github.com/todalist/app/internal/repo"
)

type TodaFlowService struct {
	repo repo.IRepo
}

func (s *TodaFlowService) Get(ctx context.Context, id uint) (*todaFlow.TodaFlow, error) {
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	return todaFlowRepo.Get(id)
}

func (s *TodaFlowService) First(ctx context.Context, querier *todaFlow.TodaFlowQuerier) (*todaFlow.TodaFlow, error) {
	todaFlowRepo := s.repo.GetTodaFlowRepo(ctx)
	return todaFlowRepo.First(querier)
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
