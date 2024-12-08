package repository

import "errors"

var (
	ErrorNotFound         = errors.New("record not found")
	ErrorDeleteFailed     = errors.New("no record entries to delete")
	ErrorInsertFailed     = errors.New("failed to insert record")
	ErrorInvalidInput     = errors.New("invalid input data")
	ErrorTypeNotSupported = errors.New("type not supported")
)
