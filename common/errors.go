package common

import "errors"

var (
	ErrNoItems = errors.New("orders must have at least one item")
)
