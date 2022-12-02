package validator

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

// 共通的なvalidationを記載する
// 個別modelと対応するようなvalidationはファイル切り出しする

func PasswordValidate(password string) error {

	err := errors.New("invalid format for password")

	if utf8.RuneCountInString(password) < 8 {
		return err
	}
	if !regexp.MustCompile(`[A-Z]+`).MatchString(password) {
		return err
	}
	if !regexp.MustCompile(`[a-z]+`).MatchString(password) {
		return err
	}
	if !regexp.MustCompile(`[0-9]+`).MatchString(password) {
		return err
	}
	if !regexp.MustCompile(`[a-zA-Z0-9!"#$%&'()*+,\-./:;<=>?@[\]^_{|}~` + "`" + `]+`).MatchString(password) {
		return err
	}

	return nil
}
