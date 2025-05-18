package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sawada-naoya/todo/backend/internal/handler"
)

// InitRouterは全ルーティングを定義するエントリーポイント
// handler層（DI済み）を受け取り、Echoにルーティングをバインドする
func InitRouter(e *echo.Echo, h *handler.TaskHandler) {
	// タスク一覧取得
	e.GET("/tasks", h.GetAllTasksHandler)
	// タスク詳細取得
	e.GET("/tasks/:id", h.GetTaskHandler)
	// タスク登録
	e.POST("/tasks", h.CreateTaskHandler)
	// タスク更新
	e.PUT("/tasks/:id", h.UpdateTaskHandler)
	// タスク削除
	e.DELETE("/tasks/:id", h.DeleteTaskHandler)
}