package country

import (
	"errors"
	"net/http"
)

func handleError(err error) (error, int) {
	var code int
	switch {
	case errors.Is(err, ErrInvalidIP):
		code = http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	return err, code
}
