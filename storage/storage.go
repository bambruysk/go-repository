package storage

import (
	"context"
	"errors"
	"io"
	"repo/entity"
	"repo/storage/inmemory"
	"repo/storage/postgres"
)

type StorageType string

const (
	StorageTypePostgres StorageType = "pg"
	StorageTypeInmemory StorageType = "mem"
)

type Config struct {
	StorageType StorageType
	Postgres    *postgres.Config
}

type Storager interface {
	Create(ctx context.Context, u entity.User) error
	GetByID(ctx context.Context, id entity.ID) (entity.User, error)
	io.Closer
}

func New(cfg *Config) (Storager, error) {
	switch cfg.StorageType {
	case StorageTypePostgres:
		return postgres.NewPostgresStorage(cfg.Postgres)
	case StorageTypeInmemory:
		return inmemory.NewMemoryStorage()
	default:
		return nil, errors.New("unknown storage type")
	}
}
