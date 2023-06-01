package task

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, task Task) (string, error)
	FindAll(ctx context.Context) ([]Task, error)
	FindOne(ctx context.Context, id string) (Task, error)
	Update(ctx context.Context, task Task) error
	Delete(ctx context.Context, id string) error
}
