package task

import (
	"encoding/json"
	"net/http"

	"github.com/qwt-tmk/pkg/validation"
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/presenter"
	"github.com/qwt-tmk/todo-rest-api/application/usecase/task"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type PostTaskHandler struct {
	createPostUsecase *task.CreateTaskUsecase
}

func NewPostTaskHandler(createPostUsecase *task.CreateTaskUsecase) *PostTaskHandler {
	return &PostTaskHandler{
		createPostUsecase: createPostUsecase,
	}
}

// @Summary		タスクを作成する
// @Description	内容、タスク状態からユーザーに紐づくタスクを作成する
// @Tags			Task
// @Produce		json
// @Param			request	body	PostTaskRequest	true	"タスク作成のための情報"
// @Security		BearerAuth
// @Success		201	{object}	presenter.SuccessResponse[PostTaskResponse]	"作成したタスクの情報"
// @Failure		400	{object}	presenter.FailureResponse					"不正なリクエスト"
// @Failure		500	{object}	presenter.FailureResponse					"内部サーバーエラー"
// @Router			/tasks [post]
func (pth *PostTaskHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// jsonをデコード
	var params PostTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err := validation.NewValidator().Struct(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	// contextからuserIdを取得
	userID := middleware.GetUserID(r.Context())
	// inputDTOに詰め替える
	input := task.CreateTaskUsecaseInputDTO{
		UserId:  userID,
		Content: params.Content,
		State:   params.State,
	}
	output, err := pth.createPostUsecase.Run(r.Context(), input)
	if err != nil && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := PostTaskResponse{
		ID:      output.ID,
		UserId:  output.UserId,
		Content: output.Content,
		State:   output.State,
	}
	presenter.RespondCreated(rw, resp)
}
