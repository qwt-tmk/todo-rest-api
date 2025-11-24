package errors

import (
	"errors"
	"fmt"
)

// ユーザー関連のドメインエラー
var (
	ErrAlreadyRegistered = newErrDomain("ErrAlreadyRegistered", "you have been already registered")
	ErrInvalidEmail      = newErrDomain("ErrInvalidEmail", "invalid email address")
	ErrPasswordMismatch  = newErrDomain("ErrPasswordMismatch", "thie password is incorrect")
	ErrPasswordTooShort  = newErrDomain("ErrPasswordTooShort", "password is too short")
	ErrNotFoundUser      = newErrDomain("ErrNotFoundUser", "user not found")
)

// タスク関係のドメインエラー
var (
	ErrNotFoundTask           = newErrDomain("ErrNotFoundTask", "task not found")
	ErrContentEmpty           = newErrDomain("ErrContentempty", "Do not empty the content")
	ErrInvalidTaskState       = newErrDomain("ErrInvalidTaskState", "invalid task state, plase select todo/doing/done")
	ErrForbiddenTaskOperation = newErrDomain("ErrForbiddenTaskOperation", "can't operate others tasks")
)

// ドメインエラー
type ErrDomain struct {
	err error
}

// ドメインエラーのコンストラクタ
func newErrDomain(errType, message string) *ErrDomain {
	return &ErrDomain{
		err: fmt.Errorf("errType: %v , Message : %v", errType, message),
	}
}

// ドメインエラーかどうか
func IsDomainErr(target error) bool {
	var errDomain *ErrDomain
	return errors.As(target, &errDomain)
}

func (e *ErrDomain) Error() string {
	return e.err.Error()
}

// errors.Is をラップ
// パッケージ名の衝突を考慮
func Is(err, target error) bool {
	return errors.Is(err, target)
}
func New(message string) error {
	return errors.New(message)
}
