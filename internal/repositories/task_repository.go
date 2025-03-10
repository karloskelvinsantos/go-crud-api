package repositories

import (
	"context"
	"go-crud-api/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	Create(ctx context.Context, task entities.Task) (primitive.ObjectID, error)
	GetAll(ctx context.Context) ([]entities.Task, error)
	Update(ctx context.Context, id primitive.ObjectID, task *entities.Task) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
