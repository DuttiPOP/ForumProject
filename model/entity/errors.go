package entity

import "errors"

var (
	ErrNotOwner = errors.New("user is not the owner")
)
