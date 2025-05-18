package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
)

// NewMySQL は、MySQLへの接続用 *sql.DB を生成する初期化関数
//
// - Wire経由で repository に依存注入するときに使う
// - main() で直接呼び出して DB に接続するときにも使える
//
// 接続情報は Docker や環境変数経由で注入された値を os.Getenv で取得し、DSN (Data Source Name) を組み立てて接続する
func NewMySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?parseTime=true",
				// NOTE: docker-composeで環境変数が設定されているので、.envファイルがなくても動く
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
  // MySQL に ping を送って、接続できるかチェック
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// NewMySQLをラップして、失敗時にpanicで即落ちする
func MustConnectMySQL() *sql.DB {
	db, err := NewMySQL()
	if err != nil {
		panic(err)
	}
	return db
}