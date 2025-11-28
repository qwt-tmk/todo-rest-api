package integration

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/qwt-tmk/todo-rest-api/infrastructure/db"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/db/container"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/kvs"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/router"
)

// NOTE: パッケージ変数としてmuxをテスト関数から使用する
var mux http.Handler

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

	// ルーティングを初期化
	mux = router.NewMux()

	code := m.Run()

	// cleaning up
	container.RemoveDockertestContainer(pool, resource)
	log.Println("success to remove docker test container")
	testDB.Close()
	cli.Close()

	os.Exit(code)
}
