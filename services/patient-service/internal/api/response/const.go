package response

import (
	"net/http"
)

const ContentTypeJSON = "application/json"

var (
	intToStatusCode = map[string]int{
		"success":                http.StatusOK,
		"created":                http.StatusCreated,
		"accepted":               http.StatusAccepted,
		"no content":             http.StatusNoContent,
		"bad request":            http.StatusBadRequest,
		"unauthorized":           http.StatusUnauthorized,
		"forbidden":              http.StatusForbidden,
		"not found":              http.StatusNotFound,
		"method not allowed":     http.StatusMethodNotAllowed,
		"not acceptable":         http.StatusNotAcceptable,
		"request timeout":        http.StatusRequestTimeout,
		"conflict":               http.StatusConflict,
		"gone":                   http.StatusGone,
		"length required":        http.StatusLengthRequired,
		"precondition failed":    http.StatusPreconditionFailed,
		"payload too large":      http.StatusRequestEntityTooLarge,
		"uri too long":           http.StatusRequestURITooLong,
		"unsupported media type": http.StatusUnsupportedMediaType,
	}
)

func InnerStatusInHttpCode(status string) int {
	return intToStatusCode[status]
}
