package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sawada-naoya/todo/backend/internal/router"
	"github.com/sawada-naoya/todo/backend/internal/wire"
)

func main() {
	// Echoインスタンスを生成
	e := echo.New()

	// Wireを使って依存関係を注入し、*handler.TaskHandler を取得する
	h := wire.InitializeHandler()

	// ルーティングを初期化して、各エンドポイントにハンドラをバインドする
	router.InitRouter(e, h)

	// サーバー起動ログを出力
	log.Println("Starting server on :8080")

	// Echoサーバーをポート8080で起動する（異常終了時はログも出力）
	e.Logger.Fatal(e.Start(":8080"))
}