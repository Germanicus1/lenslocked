package models

import "errors"

var (
	ErrEmailTaken = errors.New("models: email address is already in use")
	ErrNotFound = errors.New("models: resource could not be found")
	ErrAccountNotFound = errors.New("models: no account found")
	ErrPasswordIncorrect = errors.New("models: incorrect password")
	ErrInvalidEmail = errors.New("models: invalid email address")
)
