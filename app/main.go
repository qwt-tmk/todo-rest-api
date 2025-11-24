// @title						TODO API
// @version					1.0
// @description				This is TODO API by golang.
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
package main

import (
	"context"
	"log"

	_ "github.com/qwt-tmk/todo-rest-api/docs"

	"github.com/qwt-tmk/todo-rest-api/config"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/db"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/kvs"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config : %v", err)
	}

	run(ctx, cfg)
}

// run関数の中で様々な接続処理を行う
func run(ctx context.Context, cfg *config.Config) {
	close := db.NewDB(ctx, cfg)
	defer close()

	redisClient := kvs.NewRedisClient(ctx, cfg)
	defer redisClient.Close()
}
