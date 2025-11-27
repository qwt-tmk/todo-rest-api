package router

import (
	"net/http"

	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/middleware"
	"github.com/qwt-tmk/todo-rest-api/adapter/repository"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/auth"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/kvs"

	authUsecase "github.com/qwt-tmk/todo-rest-api/application/usecase/auth"
)

// ミドルウェア
var (
	authorization func(h http.Handler) http.Handler
	logger        func(h http.Handler) http.Handler
)

// ミドルウェアを初期化する
func initMiddlewares() {
	authorization = middleware.Authorization(
		authUsecase.NewAuthorizationUsecase(
			auth.NewJwtAuthenticator(),
			repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommandar()),
		),
	)
	logger = middleware.Logger
}

// 適用させたい順で、ミドルウェアを引数に入れる
// composeMiddlewares(M1, M2, M3)とした場合、M1(M2(M3))といったようにラップされたミドルウェアを返す
func composeMiddlewares(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := range middlewares {
			h = middlewares[len(middlewares)-(i+1)](h)
		}
		return h
	}
}
