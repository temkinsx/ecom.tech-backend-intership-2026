package repository

import (
	"sync"

	"ecom.tech-backend-intership-2026/internal/domain"
)

type repo struct {
	storage map[int64]domain.Todo
	mu      sync.RWMutex
}

func NewTodoRepository() domain.TodoRepository {
	return &repo{storage: map[int64]domain.Todo{}}
}

func (r *repo) Create(todo domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.storage[todo.ID]
	if ok {
		return domain.ErrDuplicated
	}

	r.storage[todo.ID] = todo
	return nil
}

func (r *repo) Todo(id int64) (domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	todo, ok := r.storage[id]
	if !ok {
		return domain.Todo{}, domain.ErrNotFound
	}

	return todo, nil
}

func (r *repo) ListAll() []domain.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res := make([]domain.Todo, 0, len(r.storage))
	for _, v := range r.storage {
		res = append(res, v)
	}
	return res
}

func (r *repo) Update(todo domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.storage[todo.ID]
	if !ok {
		return domain.ErrNotFound
	}

	r.storage[todo.ID] = todo
	return nil
}

func (r *repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.storage[id]
	if !ok {
		return domain.ErrNotFound
	}
	delete(r.storage, id)
	return nil
}
