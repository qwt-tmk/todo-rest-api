package testhelper

import (
	"context"
	"testing"
	"time"

	"github.com/qwt-tmk/todo-rest-api/adapter/repository"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/auth"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/kvs"
)

// ログイン状態をセットアップする
// userID: JTIのペアをKVSに保存する
// 生成したJWTトークンを返す
func LoginForTest(t *testing.T, userID string) string {
	t.Helper()

	jwtAuthenticator := auth.NewJwtAuthenticator()
	// トークン生成
	jwtToken, err := jwtAuthenticator.GenerateJwtToken(userID, "jti")
	if err != nil {
		t.Fatalf("error occured in jwtAuthenticator.GenerateJwtToken(): %v", err)
	}
	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommandar())
	// Redisに保存
	jwtAuthenticatorRepository.Save(context.Background(), time.Duration(2*time.Hour), userID, "jti")
	return jwtToken
}

func LogoutForTest(t *testing.T, userID string) {
	t.Helper()

	jwtAuthenticatorRepository := repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommandar())
	if err := jwtAuthenticatorRepository.Delete(context.Background(), userID); err != nil {
		t.Fatalf("error occured in jwtAuthenticatorRepository.Delete() :%v", err)
	}
}
