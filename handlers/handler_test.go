package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gagansingh3785/typio-service/constants"
	serviceerrors "github.com/gagansingh3785/typio-service/errors"
	"github.com/gagansingh3785/typio-service/handlers/request"
	"github.com/gagansingh3785/typio-service/handlers/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type stubRequest struct {
	validateErr    error
	validateCalled *bool
}

func (s *stubRequest) Validate() error {
	if s.validateCalled != nil {
		*s.validateCalled = true
	}
	return s.validateErr
}

type stubResponse struct {
	status          int
	getStatusCalled *bool
}

func (s *stubResponse) GetStatus() int {
	if s.getStatusCalled != nil {
		*s.getStatusCalled = true
	}
	return s.status
}

type HandlerSuite struct {
	suite.Suite
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) TestRequestProcessorError() {
	tests := []struct {
		name           string
		requestErr     error
		expectedStatus int
	}{
		{
			name:           "plain error becomes bad request",
			requestErr:     errors.New("request processor failure"),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error preserved",
			requestErr: serviceerrors.NewServiceError(
				errors.New("unprocessable"),
				"UNPROCESSABLE",
				http.StatusUnprocessableEntity,
				"unprocessable",
			),
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range tests {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			handler := Handler(
				func(ctx context.Context, r *http.Request) (request.RequestType, error) {
					return nil, tc.requestErr
				},
				func(ctx context.Context, _ request.RequestType) (any, error) {
					assert.FailNow(t, "domain processor should not be called when request processor fails")
					return nil, nil
				},
				func(ctx context.Context, _ any) response.ResponseType {
					assert.FailNow(t, "response processor should not be called when request processor fails")
					return &stubResponse{status: http.StatusOK}
				},
			)

			recorder := httptest.NewRecorder()
			httpRequest := httptest.NewRequest(http.MethodGet, "/ping", nil)

			handler(recorder, httpRequest)

			assert.Equal(t, tc.expectedStatus, recorder.Result().StatusCode)
			assert.Equal(t, constants.ContentTypeJSON, recorder.Result().Header.Get(constants.ContentTypeHeader))
		})
	}
}

func (s *HandlerSuite) TestRequestValidationError() {
	tests := []struct {
		name           string
		validateErr    error
		expectedStatus int
	}{
		{
			name:           "plain validation error becomes bad request",
			validateErr:    errors.New("invalid payload"),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error preserved",
			validateErr: serviceerrors.NewServiceError(
				errors.New("too many requests"),
				"THROTTLED",
				http.StatusTooManyRequests,
				"too many requests",
			),
			expectedStatus: http.StatusTooManyRequests,
		},
	}

	for _, tc := range tests {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			validateCalled := false

			handler := Handler(
				func(ctx context.Context, r *http.Request) (request.RequestType, error) {
					return &stubRequest{
						validateErr:    tc.validateErr,
						validateCalled: &validateCalled,
					}, nil
				},
				func(ctx context.Context, _ request.RequestType) (any, error) {
					assert.FailNow(t, "domain processor should not be called when validation fails")
					return nil, nil
				},
				func(ctx context.Context, _ any) response.ResponseType {
					assert.FailNow(t, "response processor should not be called when validation fails")
					return &stubResponse{status: http.StatusOK}
				},
			)

			recorder := httptest.NewRecorder()
			httpRequest := httptest.NewRequest(http.MethodGet, "/ping", nil)

			handler(recorder, httpRequest)

			assert.True(t, validateCalled)
			assert.Equal(t, tc.expectedStatus, recorder.Result().StatusCode)
		})
	}
}

func (s *HandlerSuite) TestDomainProcessorError() {
	tests := []struct {
		name           string
		domainErr      error
		expectedStatus int
	}{
		{
			name:           "plain error becomes internal server error",
			domainErr:      errors.New("domain failure"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "service error preserved",
			domainErr: serviceerrors.NewServiceError(
				errors.New("downstream unavailable"),
				"DOWNSTREAM_UNAVAILABLE",
				http.StatusServiceUnavailable,
				"downstream unavailable",
			),
			expectedStatus: http.StatusServiceUnavailable,
		},
	}

	for _, tc := range tests {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			validateCalled := false
			domainCalled := false
			var processed request.RequestType

			handler := Handler(
				func(ctx context.Context, r *http.Request) (request.RequestType, error) {
					req := &stubRequest{
						validateErr:    nil,
						validateCalled: &validateCalled,
					}
					processed = req
					return req, nil
				},
				func(ctx context.Context, req request.RequestType) (any, error) {
					domainCalled = true
					assert.Same(t, processed, req)
					return nil, tc.domainErr
				},
				func(ctx context.Context, _ any) response.ResponseType {
					assert.FailNow(t, "response processor should not run when domain processor fails")
					return &stubResponse{status: http.StatusOK}
				},
			)

			recorder := httptest.NewRecorder()
			httpRequest := httptest.NewRequest(http.MethodGet, "/ping", nil)

			handler(recorder, httpRequest)

			assert.True(t, validateCalled)
			assert.True(t, domainCalled)
			assert.Equal(t, tc.expectedStatus, recorder.Result().StatusCode)
		})
	}
}

func (s *HandlerSuite) TestSuccess() {
	s.T().Run("success", func(t *testing.T) {
		validateCalled := false
		domainCalled := false
		responseProcessorCalled := false
		getStatusCalled := false

		var processed request.RequestType
		expectedDomain := struct {
			Message string
		}{
			Message: "hello",
		}

		handler := Handler(
			func(ctx context.Context, r *http.Request) (request.RequestType, error) {
				req := &stubRequest{
					validateErr:    nil,
					validateCalled: &validateCalled,
				}
				processed = req
				return req, nil
			},
			func(ctx context.Context, req request.RequestType) (any, error) {
				domainCalled = true
				assert.Same(t, processed, req)
				return expectedDomain, nil
			},
			func(ctx context.Context, domainObj any) response.ResponseType {
				responseProcessorCalled = true
				assert.Equal(t, expectedDomain, domainObj)

				return &stubResponse{
					status:          http.StatusCreated,
					getStatusCalled: &getStatusCalled,
				}
			},
		)

		recorder := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodGet, "/ping", nil)

		handler(recorder, httpRequest)

		assert.True(t, validateCalled)
		assert.True(t, domainCalled)
		assert.True(t, responseProcessorCalled)
		assert.True(t, getStatusCalled)
		assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)
	})
}
