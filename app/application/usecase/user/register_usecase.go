package user

import (
	"context"

	"github.com/qwt-tmk/todo-rest-api/domain/errors"
	"github.com/qwt-tmk/todo-rest-api/domain/user"
)

type RegisterUsecase struct {
	userRepository    user.UserRepository
	userDomainService user.UserDomainService
}

func NewRegisterUsecase(
	userRepository user.UserRepository,
	userDomainService user.UserDomainService,
) *RegisterUsecase {
	return &RegisterUsecase{
		userRepository:    userRepository,
		userDomainService: userDomainService,
	}
}

func (ru *RegisterUsecase) Run(ctx context.Context, input RegisterUsecaseInputDTO) (*RegisterUsecaseOutputDTO, error) {
	// userインスタンスを生成
	u, err := user.NewUser(
		input.Email,
		input.Name,
		input.Password,
	)
	if err != nil {
		return nil, err
	}

	// userが既に登録済みではないか？
	ok, err := ru.userDomainService.IsExists(ctx, u.GetEmail())
	if err != nil {
		return nil, err
	}

	// emailが既に存在していた場合
	if ok {
		return nil, errors.ErrAlreadyRegistered
	}
	if err := ru.userRepository.Save(ctx, u); err != nil {
		return nil, err
	}

	// DTOに詰めて返す
	return &RegisterUsecaseOutputDTO{
		ID:    u.GetID(),
		Name:  u.GetName(),
		Email: u.GetEmail().Value(),
	}, nil
}
