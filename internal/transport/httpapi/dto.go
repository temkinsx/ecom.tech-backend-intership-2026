package httpapi

import "ecom.tech-backend-intership-2026/internal/domain"

type todoDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *todoDTO) ToDomain() domain.Todo {
	return domain.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

func toDTO(t domain.Todo) todoDTO {
	return todoDTO{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

func toDTOs(todos []domain.Todo) []todoDTO {
	res := make([]todoDTO, 0, len(todos))
	for _, t := range todos {
		res = append(res, toDTO(t))
	}
	return res
}

type updateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (r *updateTodoRequest) ToDomain() domain.Todo {
	return domain.Todo{
		Title:       r.Title,
		Description: r.Description,
		Completed:   r.Completed,
	}
}
