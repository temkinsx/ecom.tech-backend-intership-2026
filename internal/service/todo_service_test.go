package service

import (
	"context"
	"errors"
	"testing"

	"ecom.tech-backend-intership-2026/internal/domain"
	"ecom.tech-backend-intership-2026/internal/repository"
)

func Test_service_Create_Validation(t *testing.T) {
	tests := []struct {
		name    string
		todo    domain.Todo
		wantErr bool
	}{
		{
			name:    "create_validation_ok",
			todo:    domain.Todo{ID: 1, Title: "test"},
			wantErr: false,
		},
		{
			name:    "create_validation_error",
			todo:    domain.Todo{ID: 1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			svc := &service{
				repo: repository.NewTodoRepository(),
			}

			err := svc.Create(ctx, tt.todo)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !errors.Is(err, domain.ErrValidation) {
					t.Fatalf("expected ErrValidation, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func Test_service_Update(t *testing.T) {
	tests := []struct {
		name    string
		initial []domain.Todo
		todo    domain.Todo
		wantErr error
	}{
		{
			name:    "update_ok",
			initial: []domain.Todo{{ID: 1, Title: "old"}},
			todo:    domain.Todo{ID: 1, Title: "new"},
			wantErr: nil,
		},
		{
			name:    "update_validation_error",
			initial: []domain.Todo{{ID: 1, Title: "old"}},
			todo:    domain.Todo{ID: 1},
			wantErr: domain.ErrValidation,
		},
		{
			name:    "update_not_found",
			initial: nil,
			todo:    domain.Todo{ID: 1, Title: "new"},
			wantErr: domain.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			repo := repository.NewTodoRepository()
			svc := &service{repo: repo}

			for _, td := range tt.initial {
				if err := repo.Create(td); err != nil {
					t.Fatalf("seed repo: %v", err)
				}
			}

			err := svc.Update(ctx, tt.todo)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
