package middleware

import "net/http"

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// 開発環境用なので全てのリクエストを通すようにする
		rw.Header().Set("Access-Control-Allow-Origin", "*") // 本番環境でこの設定は禁止
		rw.Header().Set("Access-Control-Allow-Headers", "Contet-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Authorization")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, OPTIONS")
		h.ServeHTTP(rw, r)
	})
}
