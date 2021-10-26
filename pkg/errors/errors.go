package errors

import "errors"

var (
	DuplicateEmail = errors.New("duplicate email")
	DuplicateNickName = errors.New("duplicate nickname")
	BadAuth = errors.New("bad password or email")
)
