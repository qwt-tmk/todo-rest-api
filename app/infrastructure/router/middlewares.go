package router

import "net/http"

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
