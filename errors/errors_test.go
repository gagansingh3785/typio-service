package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ErrorsGoSuite struct {
	suite.Suite
}

func TestErrorsGoSuite(t *testing.T) {
	suite.Run(t, new(ErrorsGoSuite))
}

func (s *ErrorsGoSuite) TestAs4xxError() {
	tests := []struct {
		name        string
		prepare     func() (ServiceError, error)
		wantCode    string
		wantStatus  int
		wantMessage string
		wantSame    bool
	}{
		{
			name: "service error already 4xx",
			prepare: func() (ServiceError, error) {
				svc := NewServiceError(errors.New("bad thing"), "MY4XX", 429, "rate limit")
				return svc, svc
			},
			wantCode:    "MY4XX",
			wantStatus:  429,
			wantMessage: "rate limit",
			wantSame:    true,
		},
		{
			name: "service error not 4xx",
			prepare: func() (ServiceError, error) {
				svc := NewServiceError(errors.New("internal problem"), "SOMECODE", 500, "details")
				return svc, svc
			},
			wantCode:    "BAD_REQUEST",
			wantStatus:  http.StatusBadRequest,
			wantMessage: "internal problem",
		},
		{
			name: "plain error",
			prepare: func() (ServiceError, error) {
				err := errors.New("unknown error")
				return nil, err
			},
			wantCode:    "BAD_REQUEST",
			wantStatus:  http.StatusBadRequest,
			wantMessage: "unknown error",
		},
	}
	for _, tc := range tests {
		s.Run(tc.name, func() {
			original, input := tc.prepare()
			svcErr := As4xxError(input)
			assert.Implements(s.T(), (*ServiceError)(nil), svcErr)
			assert.Equal(s.T(), tc.wantCode, svcErr.Code())
			assert.Equal(s.T(), tc.wantStatus, svcErr.StatusCode())
			assert.Equal(s.T(), tc.wantMessage, svcErr.Message())
			if tc.wantSame {
				assert.Same(s.T(), original, svcErr)
			} else if original != nil {
				assert.NotSame(s.T(), original, svcErr)
			}
		})
	}
}

func (s *ErrorsGoSuite) TestAsServiceError() {
	tests := []struct {
		name       string
		prepare    func() (ServiceError, error)
		wantCode   string
		wantStatus int
		wantSame   bool
	}{
		{
			name: "input already service error",
			prepare: func() (ServiceError, error) {
				svc := NewServiceError(errors.New("known"), "CODE", 502, "msg")
				return svc, svc
			},
			wantCode:   "CODE",
			wantStatus: 502,
			wantSame:   true,
		},
		{
			name: "input plain error",
			prepare: func() (ServiceError, error) {
				err := errors.New("something failed")
				return nil, err
			},
			wantCode:   "INTERNAL_SERVER_ERROR",
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tc := range tests {
		s.Run(tc.name, func() {
			original, input := tc.prepare()
			svcErr := AsServiceError(input)
			assert.Implements(s.T(), (*ServiceError)(nil), svcErr)
			assert.Equal(s.T(), tc.wantCode, svcErr.Code())
			assert.Equal(s.T(), tc.wantStatus, svcErr.StatusCode())
			if tc.wantSame {
				assert.Same(s.T(), original, svcErr)
			} else if original != nil {
				assert.NotSame(s.T(), original, svcErr)
			}
		})
	}
}
