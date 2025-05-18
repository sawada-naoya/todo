package repository

import (
	"database/sql"

	"github.com/sawada-naoya/todo/backend/internal/domain"
)

// TaskRepositoryはタスクに関するデータ操作の契約（interface）を定義する
type TaskRepository interface {
	GetAll() ([]domain.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

// NewTaskRepositoryはTaskRepositoryの実装を返す
// *sql.DBを受け取り、DBアクセスできる構造体を返す
func NewRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// GetAllはtasksテーブルからすべてのタスクを取得する
func (r *taskRepository) GetAll() ([]domain.Task, error) {
	rows, err := r.db.Query("SELECT id, title, is_done, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.IsDone, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}