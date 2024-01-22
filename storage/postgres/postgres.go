package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"repo/entity"
	"time"
)

type Config struct {
	DatabaseDSN    string
	ConnectTimeout time.Duration
}

type postgresStorage struct {
	db *pgx.Conn
}

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrQueryExecution = errors.New("query execution failed")
	ErrConnectTimeout = errors.New("connect timeout")
)

// NewPostgresStorage создает новое хранилище данных в PostgreSQL.
func NewPostgresStorage(cfg *Config) (*postgresStorage, error) {
	connCtx, cancel := context.WithTimeoutCause(context.Background(), cfg.ConnectTimeout, ErrConnectTimeout)
	defer cancel()
	conn, err := pgx.Connect(connCtx, cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("postgres connect: %w", err)
	}

	return &postgresStorage{
		db: conn,
	}, nil
}

func (s *postgresStorage) Create(ctx context.Context, u entity.User) error {
	query := "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(ctx, query, u.ID, u.Name, u.Email)
	if err != nil {
		return ErrQueryExecution
	}

	return nil
}

func (s *postgresStorage) GetByID(ctx context.Context, id entity.ID) (entity.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := s.db.QueryRow(ctx, query, id)

	var u entity.User
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.User{}, ErrUserNotFound
	}
	if err != nil {
		return entity.User{}, ErrQueryExecution
	}

	return u, nil
}

func (s *postgresStorage) Close() error {
	err := s.db.Close(context.Background())

	if err != nil {
		return fmt.Errorf("postgres close: %w", err)
	}

	return nil
}
