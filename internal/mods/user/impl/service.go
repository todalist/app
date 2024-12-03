package userImpl

import (
	"dailydo.fe1.xyz/internal/mods/user"
	"dailydo.fe1.xyz/internal/store"
	"context"
)

type UserService struct {
	store store.IStore
}

func (s *UserService) Get(ctx context.Context, id uint) (*user.User, error) {
	userStore := s.store.GetUserStore(ctx)
	return userStore.Get(id)
}

func (s *UserService) Save(ctx context.Context, form *user.User) (*user.User, error) {
	userStore := s.store.GetUserStore(ctx)
	return userStore.Save(form)
}

func (s *UserService) List(ctx context.Context, querier *user.UserQuerier) ([]*user.User, error) {
	userStore := s.store.GetUserStore(ctx)
	return userStore.List(querier)
}

func (s *UserService) Delete(ctx context.Context, id uint) (uint, error) {
	userStore := s.store.GetUserStore(ctx)
	return userStore.Delete(id)
}