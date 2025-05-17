package models

import (
	"context"
	"database/sql"
)

type Model struct {
	repo ModelInterface
	DB   *sql.DB
}

type ModelInterface interface {
	Create(ctx context.Context, model any) (any, error)
	Update(ctx context.Context, model any) (any, error)
	Delete(ctx context.Context, model any) error
	GetById(ctx context.Context, id int) (any, error)
}

func NewModel(db *sql.DB) *Model {
	return &Model{DB: db}
}

func (s *Model) Create(ctx context.Context, model any) (any, error) {
	return s.repo.Create(ctx, model)

}

func (s *Model) GetById(ctx context.Context, id int) (any, error) {
	return s.repo.GetById(ctx, id)
}

func (s *Model) Update(ctx context.Context, model *any) (any, error) {
	return s.repo.Update(ctx, model)
}

func (s *Model) Delete(ctx context.Context, model *any) error {
	return s.repo.Delete(ctx, model)
}
