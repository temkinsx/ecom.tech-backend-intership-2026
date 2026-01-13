package service

import (
	"context"

	"ecom.tech-backend-intership-2026/internal/domain"
)

type service struct {
	repo domain.TodoRepository
}

func NewTodoService(repo domain.TodoRepository) domain.TodoService {
	return &service{repo: repo}
}

func (svc *service) Create(ctx context.Context, todo domain.Todo) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if todo.Title == "" {
		return domain.ErrValidation
	}
	return svc.repo.Create(todo)
}

func (svc *service) Todo(ctx context.Context, id int64) (domain.Todo, error) {
	if err := ctx.Err(); err != nil {
		return domain.Todo{}, err
	}
	return svc.repo.Todo(id)
}

func (svc *service) ListAll(ctx context.Context) []domain.Todo {
	if err := ctx.Err(); err != nil {
		return nil
	}
	return svc.repo.ListAll()
}

func (svc *service) Update(ctx context.Context, todo domain.Todo) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if todo.Title == "" {
		return domain.ErrValidation
	}
	return svc.repo.Update(todo)
}

func (svc *service) Delete(ctx context.Context, id int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return svc.repo.Delete(id)
}
