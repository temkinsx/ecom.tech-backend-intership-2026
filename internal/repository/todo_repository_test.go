package repository

import (
	"errors"
	"reflect"
	"sort"
	"testing"

	"ecom.tech-backend-intership-2026/internal/domain"
)

func Test_repo_Create(t *testing.T) {
	type testCase struct {
		name    string
		initial []domain.Todo
		todo    domain.Todo
		wantErr error
	}

	tests := []testCase{
		{
			name:    "create_ok",
			initial: nil,
			todo: domain.Todo{
				ID:          1,
				Title:       "test",
				Description: "test",
			},
			wantErr: nil,
		},
		{
			name: "create_fail_ErrDuplicated",
			initial: []domain.Todo{
				{ID: 1, Title: "test_1"},
			},
			todo: domain.Todo{
				ID:          1,
				Title:       "test_duplicated",
				Description: "test",
			},
			wantErr: domain.ErrDuplicated,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			storage := make(map[int64]domain.Todo, len(tt.initial))
			for _, todo := range tt.initial {
				storage[todo.ID] = todo
			}
			r := &repo{storage: storage}

			err := r.Create(tt.todo)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				if len(r.storage) != len(tt.initial) {
					t.Fatalf("storage len changed: want %d got %d", len(tt.initial), len(r.storage))
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			created, ok := r.storage[tt.todo.ID]
			if !ok {
				t.Fatalf("todo should be stored after Create")
			}
			if !reflect.DeepEqual(tt.todo, created) {
				t.Errorf("stored todo mismatch: want %+v got %+v", tt.todo, created)
			}
			if len(r.storage) != len(tt.initial)+1 {
				t.Fatalf("storage len: want %d got %d", len(tt.initial)+1, len(r.storage))
			}
		})
	}
}

func Test_repo_Todo(t *testing.T) {
	type testCase struct {
		name    string
		initial []domain.Todo
		id      int64
		want    domain.Todo
		wantErr error
	}

	tests := []testCase{
		{
			name: "todo_ok",
			initial: []domain.Todo{
				{ID: 1, Title: "test_1", Description: "test_1"},
				{ID: 2, Title: "test_2", Description: "test_2"},
			},
			id:      1,
			want:    domain.Todo{ID: 1, Title: "test_1", Description: "test_1"},
			wantErr: nil,
		},
		{
			name: "todo_fail_ErrNotFound",
			initial: []domain.Todo{
				{ID: 2, Title: "test_2", Description: "test_2"},
			},
			id:      1,
			wantErr: domain.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			storage := make(map[int64]domain.Todo, len(tt.initial))
			for _, todo := range tt.initial {
				storage[todo.ID] = todo
			}
			r := &repo{storage: storage}

			got, err := r.Todo(tt.id)

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
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %+v got %+v", tt.want, got)
			}
		})
	}
}

func Test_repo_Update(t *testing.T) {
	type testCase struct {
		name    string
		initial []domain.Todo
		todo    domain.Todo
		wantErr error
	}

	tests := []testCase{
		{
			name: "update_ok",
			initial: []domain.Todo{
				{ID: 1, Title: "old", Description: "old", Completed: false},
			},
			todo:    domain.Todo{ID: 1, Title: "new", Description: "new", Completed: true},
			wantErr: nil,
		},
		{
			name: "update_fail_ErrNotFound",
			initial: []domain.Todo{
				{ID: 2, Title: "exists"},
			},
			todo:    domain.Todo{ID: 1, Title: "new"},
			wantErr: domain.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			storage := make(map[int64]domain.Todo, len(tt.initial))
			for _, todo := range tt.initial {
				storage[todo.ID] = todo
			}
			r := &repo{storage: storage}

			err := r.Update(tt.todo)

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

			updated, ok := r.storage[tt.todo.ID]
			if !ok {
				t.Fatalf("updated todo should exist in storage")
			}
			if !reflect.DeepEqual(tt.todo, updated) {
				t.Errorf("updated todo mismatch: want %+v got %+v", tt.todo, updated)
			}
			if len(r.storage) != len(tt.initial) {
				t.Fatalf("storage len: want %d got %d", len(tt.initial), len(r.storage))
			}
		})
	}
}

func Test_repo_Delete(t *testing.T) {
	type testCase struct {
		name    string
		initial []domain.Todo
		id      int64
		wantErr error
	}

	tests := []testCase{
		{
			name: "delete_ok",
			initial: []domain.Todo{
				{ID: 1, Title: "t1"},
				{ID: 2, Title: "t2"},
			},
			id:      1,
			wantErr: nil,
		},
		{
			name: "delete_fail_ErrNotFound",
			initial: []domain.Todo{
				{ID: 2, Title: "t2"},
			},
			id:      1,
			wantErr: domain.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			storage := make(map[int64]domain.Todo, len(tt.initial))
			for _, todo := range tt.initial {
				storage[todo.ID] = todo
			}
			r := &repo{storage: storage}

			err := r.Delete(tt.id)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				if len(r.storage) != len(tt.initial) {
					t.Fatalf("storage len changed: want %d got %d", len(tt.initial), len(r.storage))
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if _, ok := r.storage[tt.id]; ok {
				t.Fatalf("todo id=%d should be removed from storage", tt.id)
			}
			if len(r.storage) != len(tt.initial)-1 {
				t.Fatalf("storage len: want %d got %d", len(tt.initial)-1, len(r.storage))
			}
		})
	}
}

func Test_repo_ListAll(t *testing.T) {
	type testCase struct {
		name    string
		initial []domain.Todo
		wantLen int
	}

	tests := []testCase{
		{
			name:    "list_empty",
			initial: nil,
			wantLen: 0,
		},
		{
			name: "list_two",
			initial: []domain.Todo{
				{ID: 2, Title: "t2"},
				{ID: 1, Title: "t1"},
			},
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			storage := make(map[int64]domain.Todo, len(tt.initial))
			for _, todo := range tt.initial {
				storage[todo.ID] = todo
			}
			r := &repo{storage: storage}

			got := r.ListAll()
			if len(got) != tt.wantLen {
				t.Fatalf("len(ListAll()) want %d got %d", tt.wantLen, len(got))
			}

			want := make([]domain.Todo, 0, len(tt.initial))
			want = append(want, tt.initial...)

			sort.Slice(got, func(i, j int) bool { return got[i].ID < got[j].ID })
			sort.Slice(want, func(i, j int) bool { return want[i].ID < want[j].ID })

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want %+v got %+v", want, got)
			}
		})
	}
}
