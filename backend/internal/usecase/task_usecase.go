package usecase

import (
	domain "github.com/sawada-naoya/todo/backend/internal/domain"
	repository "github.com/sawada-naoya/todo/backend/internal/repository"
)

// TaskUsecaseはタスクに関するユースケースのインターフェースを定義する
// ハンドラー層から呼び出され、ビジネスロジックを担う
type TaskUsecase interface {
	GetAllTasks() ([]domain.Task, error)
	GetTask(id int) (*domain.Task, error)
	CreateTask(task *domain.Task) error
	UpdateTask(task *domain.Task) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

// NewTaskUsecaseはTaskUsecaseの実装構造体を生成する
func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		repo: repo,
	}
}

// 全タスクを取得して返すユースケース処理
func (u *taskUsecase) GetAllTasks() ([]domain.Task, error) {
	return u.repo.GetAll()
}

// 指定されたIDのタスクを取得して返すユースケース処理
func (u *taskUsecase) GetTask(id int) (*domain.Task, error) {
	return u.repo.FindByID(id)
}

// 新しいタスクを作成して登録するユースケース処理
func (u *taskUsecase) CreateTask(task *domain.Task) error {
	return u.repo.Create(task)
}

// 既存のタスク情報を更新するユースケース処理
func (u *taskUsecase) UpdateTask(task *domain.Task) error {
	return u.repo.Update(task)
}

// 指定されたIDのタスクを削除するユースケース処理
func (u *taskUsecase) DeleteTask(id int) error {
	return u.repo.Delete(id)
}