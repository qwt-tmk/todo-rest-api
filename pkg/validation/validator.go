package validation

import "github.com/go-playground/validator/v10"

// シングルトンインスタンスとして用意
var validate *validator.Validate

func NewValidator() *validator.Validate {
	if validate != nil {
		return validate
	}
	// 非ポインタ構造体でもrequiredタグが有効になる
	// v11からデフォルトになるのでオプションをつけている
	validate = validator.New(validator.WithRequiredStructEnabled())
	return validate
}
