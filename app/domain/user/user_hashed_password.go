package user

import (
	"unicode/utf8"

	"github.com/qwt-tmk/pkg/hash"
	"github.com/qwt-tmk/todo-rest-api/domain/errors"
)

type HashedPassword struct {
	value string
}

const (
	minPasswordLength = 6
)

func newHashedPassword(value string) (HashedPassword, error) {
	// validation
	if minPasswordLength >= utf8.RuneCountInString(value) {
		return HashedPassword{}, errors.ErrPasswordTooShort
	}

	// hash
	hashed, err := hash.Hash(value)
	if err != nil {
		return HashedPassword{}, err
	}

	return HashedPassword{value: hashed}, nil
}

func reconstructHashedPassword(value string) HashedPassword {
	return HashedPassword{value: value}
}

func (hp HashedPassword) Value() string {
	return hp.value
}

// ハッシュ化されたパスワードと比較
// 集約ルートUserから呼び出す
func (p HashedPassword) compare(target string) error {
	if err := hash.Compare(p.value, target); err != nil {
		return errors.ErrPasswordMismatch
	}
	return nil
}
