package repository_test

import (
	"log"
	"os"
	"testing"

	"github.com/qwt-tmk/todo-rest-api/infrastructure/db"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/db/container"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/kvs"
)

func TestMain(m *testing.M) {
	// dockertestコンテナを起動
	pool, resource := container.NewDockertestContaier()
	log.Println("success to start dockertest container")
	// DBに接続
	testDB := container.NewDB(pool, resource)
	log.Println("success to connect test-db")
	// マイグレーションを適用させる
	container.SetupDB()
	log.Println("success to apply migrations")
	// dbパッケージ変数にテスト用DBをセット
	db.SetDB(testDB)
	log.Println("dockertest & test-db settings complete")

	// テスト用のredisサーバーを起動
	cli := kvs.NewRedisTestClient()

	code := m.Run()

	container.RemoveDockertestContainer(pool, resource)
	log.Println("success to remove dockertest container")
	testDB.Close()

	cli.Close()

	os.Exit(code)
}
