package client

import "errors"

var (
	ErrStatus = errors.New("server replied with status: ")
)
