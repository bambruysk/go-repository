package userservice

import (
	"context"
	"repo/entity"
)

type Storager interface {
	Create(ctx context.Context, u entity.User) error
	GetByID(ctx context.Context, id entity.ID) (entity.User, error)
}

type service struct {
	storage Storager
}

func New(storage Storager) *service {
	return &service{storage: storage}
}

func (u *service) Create(ctx context.Context, user entity.User) error {
	return u.storage.Create(ctx, user)
}

func (u *service) GetByID(ctx context.Context, id entity.ID) (entity.User, error) {
	return u.storage.GetByID(ctx, id)
}
