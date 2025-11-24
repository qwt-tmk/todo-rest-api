package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qwt-tmk/todo-rest-api/config"
)

// パッケージ変数としてDB接続を管理
var (
	db *sql.DB
)

// パッケージ変数として*sql.DBをセット
func SetDB(d *sql.DB) {
	db = d
}

func GetDB() *sql.DB {
	return db
}

// DBへの接続
func NewDB(ctx context.Context, cfg *config.Config) func() {
	db, close, err := connect(
		ctx,
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Name,
	)
	if err != nil {
		// DB接続失敗は復旧不可
		panic(err)
	}
	SetDB(db)
	return close
}

const (
	maxRetriesCount = 5
	initialDelay    = 5 * time.Second
)

// DBへ接続するヘルパー
// 失敗したらエクポネンシャルバックオフでリトライ
func connect(ctx context.Context, user, password, host, port, name string) (*sql.DB, func(), error) {
	delay := initialDelay
	for i := range maxRetriesCount {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("failed to open db (attempt %d/%d): %v", i+1, maxRetriesCount, err)
		} else {
			if err := db.PingContext(ctx); err == nil {
				// 成功した場合、呼び出し元でdb.Close()を実行できるようにする
				log.Printf("successfully connected to db (attempt %d/%d)", i+1, maxRetriesCount)
			} else {
				// 接続できた場合のエラーメッセージ
				log.Printf("failed to ping db: %v", err)
			}
		}

		// 接続できなかったらリトライ
		log.Printf("failed to connect to db (attempt %d/%d), retrying in %v seconds...", i+1, maxRetriesCount, delay/time.Second)
		time.Sleep(delay)
		// 待機時間を2倍
		delay *= 2
	}

	return nil, nil, fmt.Errorf("failed to connect to db after %d attempts", maxRetriesCount)
}
