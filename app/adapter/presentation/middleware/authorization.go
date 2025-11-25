package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/presenter"
	"github.com/qwt-tmk/todo-rest-api/application/usecase/auth"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type userIDKey struct{}

func Authorization(authorizationUsecase *auth.AuthorizationUsecase) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			jwtToken, err := getJwtToken(r)
			if err != nil {
				presenter.RespondUnAuthorized(rw, err.Error())
				return
			}
			input := auth.AuthorizationInputDTO{JwtToken: jwtToken}
			output, err := authorizationUsecase.Run(r.Context(), input)
			if err != nil {
				presenter.RespondUnAuthorized(rw, err.Error())
				return
			}
			userID := output.UserID
			// コンテキストにuserIDを含める
			ctx := context.WithValue(r.Context(), userIDKey{}, userID)
			// 後続の処理（ハンドラ）を実行する
			h.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

func getJwtToken(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("authorization Header is missing")
	}
	// スペースを区切りとして最大2つに分割
	parts := strings.SplitN(authorizationHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header format")
	}
	return parts[1], nil
}

// コンテキストに設定したユーザーIDを取得する
// ハンドラーに使用する
func GetUserID(context context.Context) string {
	return context.Value(userIDKey{}).(string)
}
