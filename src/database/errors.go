package database

import "errors"

var (
	ErrPackIdInvalid    = errors.New("invalid pack id provided")
	ErrPackTitleInvalid = errors.New("invalid pack title provided")
)
