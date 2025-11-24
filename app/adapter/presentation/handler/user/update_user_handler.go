package user

import (
	"encoding/json"
	"net/http"

	"github.com/qwt-tmk/pkg/validation"
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/presenter"
	"github.com/qwt-tmk/todo-rest-api/application/usecase/user"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type UpdateUserHandler struct {
	updateUserUsecase *user.UpdateProfileUsecase
}

func NewUpdateUserHandler(updateUserHandler *user.UpdateProfileUsecase) *UpdateUserHandler {
	return &UpdateUserHandler{
		updateUserUsecase: updateUserHandler,
	}
}

func (uuh *UpdateUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	var params UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err := validation.NewValidator().Struct(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	input := user.UpdateProfileUsecaseInputDTO{
		ID:    userID,
		Email: params.Email,
		Name:  params.Name,
	}
	ctx := r.Context()
	output, err := uuh.updateUserUsecase.Run(ctx, input)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := UpdateUserResponse{
		ID:    output.ID,
		Email: output.Email,
		Name:  output.Name,
	}
	presenter.RespondOK(rw, resp)
}
