package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"ecom.tech-backend-intership-2026/internal/domain"
)

type Handler struct {
	svc domain.TodoService
}

func NewTodoHandler(svc domain.TodoService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req todoDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	todo := req.ToDomain()

	if err := h.svc.Create(r.Context(), todo); err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, toDTO(todo))
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.svc.ListAll(r.Context())
	writeJSON(w, http.StatusOK, toDTOs(todos))
}

func (h *Handler) TodoByID(w http.ResponseWriter, r *http.Request) {
	id, ok, err := parseID(r.URL.Path)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if !ok {
		writeJSONError(w, http.StatusNotFound, "not found")
		return
	}

	todo, err := h.svc.Todo(r.Context(), id)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toDTO(todo))
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, ok, err := parseID(r.URL.Path)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if !ok {
		writeJSONError(w, http.StatusNotFound, "not found")
		return
	}

	var req updateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	todo := req.ToDomain()
	todo.ID = id
	if err := h.svc.Update(r.Context(), todo); err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toDTO(todo))
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, ok, err := parseID(r.URL.Path)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if !ok {
		writeJSONError(w, http.StatusNotFound, "not found")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func parseID(path string) (int64, bool, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 2 || parts[0] != "todos" {
		return 0, false, nil
	}
	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, true, err
	}
	return id, true, nil
}

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrValidation):
		writeJSONError(w, http.StatusBadRequest, "validation error")
	case errors.Is(err, domain.ErrDuplicated):
		writeJSONError(w, http.StatusConflict, "todo already exists")
	case errors.Is(err, domain.ErrNotFound):
		writeJSONError(w, http.StatusNotFound, "todo not found")
	default:
		writeJSONError(w, http.StatusInternalServerError, "internal error")
	}
}
