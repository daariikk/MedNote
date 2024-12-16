package repository

import "errors"

var (
	ErrorNotFound         = errors.New("reminder not found")
	ErrorDeleteFailed     = errors.New("no reminder entries to delete")
	ErrorInsertFailed     = errors.New("failed to insert record")
	ErrorInvalidInput     = errors.New("invalid input data")
	ErrorTypeNotSupported = errors.New("type not supported")
)
