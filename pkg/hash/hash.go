package hash

import "golang.org/x/crypto/bcrypt"

func Hash(tareget string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(tareget), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Compare(hashed, target string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(target))
	return err
}
