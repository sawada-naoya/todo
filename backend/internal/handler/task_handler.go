package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawada-naoya/todo/backend/internal/domain"
	"github.com/sawada-naoya/todo/backend/internal/usecase"
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

// GetAllTasksHandlerは全タスクを取得してJSONで返す
func (h *TaskHandler) GetAllTasksHandler(c echo.Context) error {
	tasks, err := h.usecase.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

// 指定IDのタスクを取得しJSONで返す
func (h *TaskHandler) GetTaskHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}
	task, err := h.usecase.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get task"})
	}
	return c.JSON(http.StatusOK, task)
}

// リクエストBodyからタスクを登録する
func (h *TaskHandler) CreateTaskHandler(c echo.Context) error {
	var task domain.Task
	// HTTPリクエストBody（JSON）を domain.Task 構造体に変換している
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	if err := h.usecase.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create task")
	}
	return c.JSON(http.StatusCreated, task)
}

// 指定IDのタスク情報を更新する
func (h *TaskHandler) UpdateTaskHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}

	var task domain.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	task.ID = id
	if err := h.usecase.UpdateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update task"})
	}
	return c.JSON(http.StatusOK, task)
}

// 指定IDのタスクを削除する
func (h *TaskHandler) DeleteTaskHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task id"})
	}
	if err := h.usecase.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}