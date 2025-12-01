package entity

import "errors"

var (
	ErrInvalidAnswer = errors.New("invalid answer value")
	ErrInvalidChoose = errors.New("invalid choose value")
)
