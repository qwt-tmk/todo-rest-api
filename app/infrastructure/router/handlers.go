package router

import (
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/handler/auth"
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/handler/task"
	"github.com/qwt-tmk/todo-rest-api/adapter/presentation/handler/user"
	"github.com/qwt-tmk/todo-rest-api/adapter/queryservice"
	"github.com/qwt-tmk/todo-rest-api/adapter/repository"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/db/sqlc"
	"github.com/qwt-tmk/todo-rest-api/infrastructure/kvs"

	userDomain "github.com/qwt-tmk/todo-rest-api/domain/user"

	authUsecase "github.com/qwt-tmk/todo-rest-api/application/usecase/auth"
	taskUsecase "github.com/qwt-tmk/todo-rest-api/application/usecase/task"
	userUsecase "github.com/qwt-tmk/todo-rest-api/application/usecase/user"

	authInfra "github.com/qwt-tmk/todo-rest-api/infrastructure/auth"
)

// 認証系ハンドラー
var (
	loginHandler  *auth.LoginHandler
	logoutHandler *auth.LogoutHandler
)

// ユーザー系ハンドラー
var (
	postUserHandler   *user.PostUserHandler
	deleteUserHandler *user.DeleteUserHandler
	updateUserHandler *user.UpdateUserHandler
	getUsersHandler   *user.GetUsersHandler
)

// タスク系ハンドラー
var (
	postTaskHandler        *task.PostTaskHandler
	deleteTaskHandler      *task.DeleteTaskHandler
	updateTaskStateHandler *task.UpdateTaskStateHandler
	getTaskHandler         *task.GetTaskHandler
	getTasksHandler        *task.GetTasksHandler
	getUserTasksHandler    *task.GetUserTasksHandler
)

// ハンドラーを初期化する
func initHandlers() {
	initAuthHandlers()
	initUserHandlers()
	initTaskHandlers()
}

func initAuthHandlers() {
	loginHandler = auth.NewLoginHandler(
		authUsecase.NewLoginUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
			repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommandar()),
			authInfra.NewJwtAuthenticator(),
		),
	)

	logoutHandler = auth.NewLogoutHandler(
		authUsecase.NewLogoutUsecase(
			authInfra.NewJwtAuthenticator(),
			repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommandar()),
		),
	)
}

func initUserHandlers() {
	postUserHandler = user.NewPostUserHandler(
		userUsecase.NewRegisterUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
			userDomain.NewUserDomainService(repository.NewUserRepository(sqlc.NewSqlcQuerier())),
		),
	)

	deleteUserHandler = user.NewDeleteUserHandler(
		userUsecase.NewUnregisterUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)

	getUsersHandler = user.NewGetUsersHandler(
		userUsecase.NewFetchUsersUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)

	updateUserHandler = user.NewUpdateUserHandler(
		userUsecase.NewUpdateProfileUsecase(
			repository.NewUserRepository(sqlc.NewSqlcQuerier()),
		),
	)
}

func initTaskHandlers() {
	postTaskHandler = task.NewPostTaskHandler(
		taskUsecase.NewCreateTaskUsecase(
			repository.NewTaskRepository(sqlc.NewSqlcQuerier()),
		),
	)

	deleteTaskHandler = task.NewDeleteTaskHandler(
		taskUsecase.NewDeleteTaskUsecase(
			repository.NewTaskRepository(sqlc.NewSqlcQuerier()),
		),
	)

	updateTaskStateHandler = task.NewUpdateTaskStateHandler(
		taskUsecase.NewUpdateTaskStateUsecase(
			repository.NewTaskRepository(sqlc.NewSqlcQuerier()),
		),
	)

	getTaskHandler = task.NewGetTaskHandler(
		taskUsecase.NewFetchTaskUsecase(
			queryservice.NewTaskQueryService(sqlc.NewSqlcQuerier()),
		),
	)

	getTasksHandler = task.NewGetTasksHandler(
		taskUsecase.NewFetchTasksUsecase(queryservice.NewTaskQueryService(sqlc.NewSqlcQuerier())),
	)

	getUserTasksHandler = task.NewGetUserTasksHandler(
		taskUsecase.NewFetchUserTasksUsecase(queryservice.NewTaskQueryService(sqlc.NewSqlcQuerier())),
	)
}
