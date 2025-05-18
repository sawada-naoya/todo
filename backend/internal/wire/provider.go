package wire

import (
	"github.com/google/wire"
	"github.com/sawada-naoya/todo/backend/internal/handler"
	"github.com/sawada-naoya/todo/backend/internal/infrastructure"
	"github.com/sawada-naoya/todo/backend/internal/repository"
	"github.com/sawada-naoya/todo/backend/internal/usecase"
)

// ProviderSetはWireが使う依存解決セット
var ProviderSet = wire.NewSet(
	infrastructure.MustConnectMySQL,// *sql.DB
	repository.NewTaskRepository,   // TaskRepository
	usecase.NewTaskUsecase,         // TaskUsecase
	handler.NewTaskHandler,         // *TaskHandler
)