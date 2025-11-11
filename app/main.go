package main

import (
	"context"
	"log"

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
