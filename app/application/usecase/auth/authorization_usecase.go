package auth

import (
	"context"

	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type AuthorizationUsecase struct {
	jwtAuthenticator           JwtAuthenticator
	jwtAuthenticatorRepository JwtAuthenticatorRepository
}

func NewAuthorizationUsecase(
	jwtAuthenticator JwtAuthenticator,
	jwtAuthenticatorRepository JwtAuthenticatorRepository,
) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		jwtAuthenticator:           jwtAuthenticator,
		jwtAuthenticatorRepository: jwtAuthenticatorRepository,
	}
}

func (au *AuthorizationUsecase) Run(ctx context.Context, input AuthorizationInputDTO) (
	*AuthorizationOutputDTO,
	error,
) {
	// 公開鍵で署名済みトークンを検証する
	userID, jti, err := au.jwtAuthenticator.VerifyJwtToken(input.JwtToken)
	if err != nil {
		return nil, err
	}
	// KVSから保存されたjtiを取得
	// ログアウトしていた場合は、nilが返る
	jtiFromKVS, err := au.jwtAuthenticatorRepository.Load(ctx, userID)
	if err != nil {
		return nil, err
	}
	// jtiが一致しない場合はエラー
	// ログアウトしている場合はここでエラーとなる
	if jti != jtiFromKVS {
		return nil, errors.New("invalid JWT IDJ")
	}
	return &AuthorizationOutputDTO{
		UserID: userID,
	}, nil
}
