package user

import (
	"encoding/json"
	"net/http"

	"github.com/qwt-tmk/pkg/validation"
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/presenter"
	"github.com/qwt-tmk/todo-rest-api/application/usecase/user"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type PostUserHandler struct {
	registerUsecase *user.RegisterUsecase
}

func NewPostUserHandler(registerUsecase *user.RegisterUsecase) *PostUserHandler {
	return &PostUserHandler{
		registerUsecase: registerUsecase,
	}
}

// @Summary ユーザーの登録
// @Description 新しいユーザーを登録する
// @Tags User
// @Accept json
// @Product json
// @Param request body PostUserRequest true "ユーザー登録のための情報"
// @Success 201 {object} presenter.SuccessResponse[PostUserResponse] "登録されたユーザーの情報"
// @Failure 400 {object} presenter.FailureResponse "不正なリクエスト"
// @Failure 500 {object} presenter.FailureResponse "内部サーバーエラー"
// @Router /users [post]
func (puh *PostUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// jsonをデコード
	var params PostUserRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err := validation.NewValidator().Struct(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	// DTOに詰め替える
	input := user.RegisterUsecaseInputDTO{
		Name: params.Name,
		Email: params.Email,
		Password: params.Password,
	}
	// ユースケースに渡して実行
	ctx := r.Context()
	output, err := puh.registerUsecase.Run(ctx, input)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := PostUserResponse{
		ID: output.ID,
		Email: output.Email,
		Name: output.Name,
	}
	presenter.RespondCreated(rw, resp)
}
