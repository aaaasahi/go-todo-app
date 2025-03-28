package services

import (
	"github.com/jmoiron/sqlx"
)

// サービス構造体の定義（Todoサービス）
type TodoService struct {
	db *sqlx.DB // フィールドに*sqlx.DB（データベース接続）を保持
}

// コンストラクタ関数
// 外部から*sqlx.DBを受け取り、サービス構造体を生成
func NewTodoService(db *sqlx.DB) *TodoService {
	return &TodoService{db: db}
}
