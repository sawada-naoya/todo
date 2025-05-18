package domain

import "time"

// Taskはタスク情報を保持するエンティティ
// DBのtasksテーブルの構造に対応させている
type Task struct {
	ID 				int
	Title 		string
	IsDone 		bool
	CreatedAt time.Time
	UpdatedAt time.Time
}