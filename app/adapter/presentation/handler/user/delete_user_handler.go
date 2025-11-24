package user

import (
	"net/http"

	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/presenter"
	"github.com/qwt-tmk/todo-rest-api/application/usecase/user"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type DeleteUserHandler struct {
	unregisterUsecase *user.UnregisterUsecase
}

func NewDeleteUserHandler(unregisterUsecase *user.UnregisterUsecase) *DeleteUserHandler {
	return &DeleteUserHandler{
		unregisterUsecase: unregisterUsecase,
	}
}

// @Summary		ユーザーの退会
// @Description	ユーザーを退会させ、ユーザー情報を削除する
// @Tags			User
// @Produce		json
// @Security		BearerAuth
// @Success		204
// @Failure		400	{object}	presenter.FailureResponse	"不正なリクエスト"
// @Failure		500	{object}	presenter.FailureResponse	"内部サーバーエラー"
// @Router			/users/me [delete]
func (duh *DeleteUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// リクエストスコープのコンテキストからuserIdを取得
	userID := middleware.GetUserID(r.Context())
	input := user.UnregisterUsecaseInputDTO{
		ID: userID,
	}
	ctx := r.Context()
	err := duh.unregisterUsecase.Run(ctx, input)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}

	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	presenter.RespondNoContent(rw)
}
