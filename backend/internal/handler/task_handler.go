package handler

import (
	"github.com/sawada-naoya/todo/backend/internal/usecase"
	"net/http"
	"encoding/json"
)

// TaskHandlerはタスクに関するHTTPハンドラの構造体
type TaskHandler struct {
	usecase usecase.TaskUsecase
}

// NewTaskHandlerはTaskHandlerを初期化して返す
func NewTaskHandler(u usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		usecase: u,
	}
}

// GetAllTasksHandlerは全タスクを取得してJSONで返すHTTPハンドラ
func (h *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.usecase.GetAllTasks()
	if err != nil {
		http.Error(w, "failed to get tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// tasks（構造体スライス）をJSON形式に変換して、HTTPレスポンスに書き出す
	json.NewEncoder(w).Encode(tasks)
}
