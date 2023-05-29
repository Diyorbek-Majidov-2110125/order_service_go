package postgres

import (
	"context"
	"fmt"
	"practice1/order_service_go/config"
	"practice1/order_service_go/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db      *pgxpool.Pool
	order   storage.OrderRepoI
	product storage.ProductRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product != nil {
		s.product = NewProductRepo(s.db)
	}
	return s.product
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order != nil {
		s.order = NewOrderRepo(s.db)
	}
	return s.order
}
