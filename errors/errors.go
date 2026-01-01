package errors

import (
	"errors"
	"net/http"
)

var (
	GenericBadRequestError = NewServiceError(errors.New("bad request"), "BAD_REQUEST", http.StatusBadRequest, "bad request")

	NoParagraphsFoundError = NewServiceError(errors.New("no paragraphs found"), "NO_PARAGRAPHS_FOUND", http.StatusNotFound, "no paragraphs found")
)

func As4xxError(err error) ServiceError {
	var svcError *serviceError

	if errors.As(err, &svcError) && svcError.Is4xxError() {
		return svcError
	}

	return NewServiceError(

		err,

		"BAD_REQUEST",

		http.StatusBadRequest,

		err.Error(),
	)
}

func AsServiceError(err error) ServiceError {
	var svcError *serviceError

	if errors.As(err, &svcError) {
		return svcError
	}

	return NewServiceError(

		err,

		"INTERNAL_SERVER_ERROR",

		http.StatusInternalServerError,

		err.Error(),
	)
}
