package devicemanager

import "errors"

var (
	ErrAuthentication  = errors.New("authentication failed: invalid token")
	ErrMalformedEntity = errors.New("malformed entity")
	ErrNotFound        = errors.New("entity not found")
	ErrQueryFailed     = errors.New("query failed")
	ErrDB              = errors.New("database error")
)
