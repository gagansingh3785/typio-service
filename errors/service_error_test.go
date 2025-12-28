package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceErrorSuite struct {
	suite.Suite
}

func TestServiceErrorSuite(t *testing.T) {
	suite.Run(t, new(ServiceErrorSuite))
}

func (s *ServiceErrorSuite) TestServiceError_Error() {
	baseErr := errors.New("something went wrong")
	svcErr := NewServiceError(baseErr, "ERR_CODE", 500, "error occurred")
	assert.Equal(s.T(), baseErr.Error(), svcErr.Error())
}

func (s *ServiceErrorSuite) TestServiceError_StatusCode() {
	baseErr := errors.New("fail")
	svcErr := NewServiceError(baseErr, "ANY_CODE", 418, "msg")
	assert.Equal(s.T(), 418, svcErr.StatusCode())
}

func (s *ServiceErrorSuite) TestServiceError_Message() {
	tests := []struct {
		message     string
		err         error
		wantMessage string
	}{
		{message: "custom message", wantMessage: "custom message"},
		{message: "", err: errors.New("some error"), wantMessage: "some error"},
	}
	for _, tc := range tests {
		svcErr := NewServiceError(tc.err, "ANY_CODE", 401, tc.message)
		assert.Equal(s.T(), tc.wantMessage, svcErr.Message())
	}
}

func (s *ServiceErrorSuite) TestServiceError_Code() {
	baseErr := errors.New("fail")
	svcErr := NewServiceError(baseErr, "MY_CODE", 404, "msg")
	assert.Equal(s.T(), "MY_CODE", svcErr.Code())
}

func (s *ServiceErrorSuite) TestServiceError_Is4xxError() {
	tests := []struct {
		statusCode int
		want4xx    bool
	}{
		{399, false},
		{400, true},
		{404, true},
		{499, true},
		{500, false},
	}

	for _, tc := range tests {
		baseErr := errors.New("err")
		se := NewServiceError(baseErr, "C", tc.statusCode, "msg").(*serviceError)
		assert.Equal(s.T(), tc.want4xx, se.Is4xxError())
	}
}

func (s *ServiceErrorSuite) TestNewServiceError_ReturnsServiceError() {
	err := errors.New("err!")
	svcErr := NewServiceError(err, "CODE", 400, "message")
	assert.NotNil(s.T(), svcErr)
	assert.IsType(s.T(), &serviceError{}, svcErr)
}
