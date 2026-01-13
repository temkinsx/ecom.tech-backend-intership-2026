package domain

import "context"

type Todo struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
}

type TodoRepository interface {
	Create(todo Todo) error
	Todo(id int64) (Todo, error)
	ListAll() []Todo
	Update(todo Todo) error
	Delete(id int64) error
}

type TodoService interface {
	Create(ctx context.Context, todo Todo) error
	Todo(ctx context.Context, id int64) (Todo, error)
	ListAll(ctx context.Context) []Todo
	Update(ctx context.Context, todo Todo) error
	Delete(ctx context.Context, id int64) error
}
