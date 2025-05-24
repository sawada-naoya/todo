package repository

import (
	"database/sql"

	"github.com/sawada-naoya/todo/backend/internal/models"
)

// TaskRepositoryはタスクに関するデータ操作の契約（interface）を定義する
// CRUD全てを扱う
// - GetAll: 全件取得
// - FindByID: ID指定で1件取得
// - Create: 新規作成
// - Update: タイトルやステータスの更新
// - Delete: 削除

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	FindByID(id int) (*models.Task, error)
	Create(task *models.Task) error
	UpdateIsDone(id int, isDone bool) error
	Delete(id int) error
}

type taskRepository struct {
	db *sql.DB
}

// NewTaskRepositoryはTaskRepositoryの実装を返す
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

// GetAllはtasksテーブルからすべてのタスクを取得する
func (r *taskRepository) GetAll() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT id, title, is_done, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.IsDone, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// FindByIDは指定されたIDのタスクを1件取得する
func (r *taskRepository) FindByID(id int) (*models.Task, error) {
	row := r.db.QueryRow("SELECT id, title, is_done, created_at, updated_at FROM tasks WHERE id = ?", id)
	var t models.Task
	if err := row.Scan(&t.ID, &t.Title, &t.IsDone, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	return &t, nil
}

// Createは新しいタスクをtasksテーブルに挿入する
func (r *taskRepository) Create(task *models.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, is_done) VALUES (?, ?)", task.Title, task.IsDone)
	return err
}

// Updateは既存のタスクの内容を更新する
func (r *taskRepository) UpdateIsDone(id int, isDone bool) error {
	_, err := r.db.Exec(
		"UPDATE tasks SET is_done = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		isDone, id,
	)
	return err
}

// Deleteは指定されたIDのタスクを削除する
func (r *taskRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
