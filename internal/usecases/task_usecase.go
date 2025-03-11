package usecases

import (
	"context"
	"go-crud-api/internal/entities"
	"go-crud-api/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error)
	GetTasks(ctx context.Context) ([]entities.Task, error)
	UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error
	DeleteTask(ctx context.Context, id primitive.ObjectID) error
}

type taskUseCase struct {
	repo repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) TaskUseCase {
	return &taskUseCase{repo: repo}
}

func (uc *taskUseCase) CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error) {
	return uc.repo.Create(ctx, task)
}

func (uc *taskUseCase) GetTasks(ctx context.Context) ([]entities.Task, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *taskUseCase) UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error {
	return uc.repo.Update(ctx, id, task)
}

func (uc *taskUseCase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	return uc.repo.Delete(ctx, id)
}
