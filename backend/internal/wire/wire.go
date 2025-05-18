//go:build wireinject
// +build wireinject

// このビルドタグは「wireコマンドで使う専用ファイル」であることをGoに伝える
// → go build などの通常ビルドでは無視される（wire専用）

package wire

import (
	"github.com/sawada-naoya/todo/backend/internal/handler"

	"github.com/google/wire"
)

// InitializeHandlerは依存関係を自動構築してTaskHandlerを返すためのエントリーポイント関数
// wire.Build() に渡されたProviderSetを元に、構成を自動生成する
func InitializeHandler() *handler.TaskHandler {
	wire.Build(ProviderSet) // ProviderSet（NewMySQL〜NewTaskHandler）を使って依存解決を定義する
	return nil              // このreturnはダミー。wireコマンドで書き換えられるので問題ない
}
