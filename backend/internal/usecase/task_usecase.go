package usecase

import (
	domain "github.com/sawada-naoya/todo/backend/internal/domain"
	repository "github.com/sawada-naoya/todo/backend/internal/repository"
)

// TaskUsecaseはタスク取得処理のユースケース
// インターフェースにより上位層から疎結合に呼び出せるようにする
type TaskUsecase interface {
	GetAllTasks() ([]domain.Task, error)
}

type taskUsecase struct {
	repo repository.TaskRepository
}

// NewTaskUsecaseはTaskUsecaseの実装を返す
func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		repo: repo,
	}
}

// GetAllTasksは全タスクを取得して返すユースケース処理
func (u *taskUsecase) GetAllTasks() ([]domain.Task, error) {
	return u.repo.GetAll()
}