package inmemory

import (
	"context"
	"errors"
	"repo/entity"
	"sync"
)

type memoryStorage struct {
	data  map[entity.ID]entity.User
	mutex sync.RWMutex
}

var (
	ErrUserNotFound = errors.New("user not found")
)

// NewMemoryStorage создает новое хранилище данных в памяти.
func NewMemoryStorage() (*memoryStorage, error) {
	return &memoryStorage{
		data: make(map[entity.ID]entity.User),
	}, nil
}

func (s *memoryStorage) Create(_ context.Context, u entity.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[u.ID] = u
	return nil
}

func (s *memoryStorage) GetByID(_ context.Context, id entity.ID) (entity.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	u, ok := s.data[id]
	if !ok {
		return entity.User{}, ErrUserNotFound
	}
	return u, nil
}

func (s *memoryStorage) Close() error {
	return nil
}
