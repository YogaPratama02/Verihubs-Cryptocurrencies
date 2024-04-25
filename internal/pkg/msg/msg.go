package msg

import "errors"

const (
	MsgHeaderTokenNotFound = "Header `token` not found"
	MsgUnauthorized        = "You must login before"

	MsgHeaderTokenUnauthorized = "Unauthorized token"
)

var (
	ErrMissingHeaderData = errors.New("missing header data")
	ErrInvalidToken      = errors.New("invalid token")
)
