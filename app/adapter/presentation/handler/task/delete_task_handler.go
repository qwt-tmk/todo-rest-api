package task

import (
	"net/http"

	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/middleware"
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/presenter"
	"github.com/qwt-tmk/todo-rest-api/application/usecase/task"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type DeleteTaskHandler struct {
	deleteTaskUsecase *task.DeleteTaskUsecase
}

func NewDeleteTaskHandler(deleteTaskUsecase *task.DeleteTaskUsecase) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		deleteTaskUsecase: deleteTaskUsecase,
	}
}

// @Summary		タスクを削除する
// @Description	指定したidのタスクを削除する
// @Tags			Task
// @Produce		json
// @Security		BearerAuth
// @Success		204
// @Failure		400	{object}	presenter.FailureResponse	"不正なリクエスト"
// @Failure		403	{object}	presenter.FailureResponse	"権限エラー"
// @Failure		500	{object}	presenter.FailureResponse	"内部サーバーエラー"
// @Router			/tasks/{id} [delete]
func (dth *DeleteTaskHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// パスパラメータから取得
	id := r.PathValue("id")
	// contextからuserIdを取得
	userID := middleware.GetUserID(r.Context())
	// inputDTOに詰め替える
	input := task.DeleteTaskUsecaseInputDTO{
		ID:     id,
		UserId: userID,
	}
	err := dth.deleteTaskUsecase.Run(r.Context(), input)
	// タスクを削除する権限がない（ログインしているユーザーのタスクでない）場合
	if err != nil && errors.Is(err, errors.ErrForbiddenTaskOperation) {
		presenter.RespondForbidden(rw, err.Error())
		return
	}
	// ドメインエラー
	if err != nil && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	// その他のエラー
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	presenter.RespondNoContent(rw)
}
