package errcodes

import "net/http"

const (
	CodeForbidden       = "FORBIDDEN"
	BadCollectionParams = "BAD_COLLECTION_PARAMS"
)

var statusCodes = map[string]int{
	CodeForbidden:       http.StatusForbidden,
	BadCollectionParams: http.StatusBadRequest,
}
