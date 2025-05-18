package db

import (
	"database/sql"
	"fmt"
	"os"
)

// 受け取った db を使って、マイグレーション（SQLの実行）だけを行う
// DBに CREATE TABLE や ALTER TABLE を実行する（構造を整える）
func Migrate(db *sql.DB) error {
	sqlByte, err := os.ReadFile("./db/init.sql")
	if err != nil {
		return fmt.Errorf("failed to read init.sql: %w", err)
	}

	_, err = db.Exec(string(sqlByte))
	if err != nil {
		return fmt.Errorf("failed to execute init.sql: %w", err)
	}

	return nil
}