package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gagansingh3785/typio-service/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PingHandlerSuite struct {
	suite.Suite
}

func TestPingHandlerSuite(t *testing.T) {
	suite.Run(t, new(PingHandlerSuite))
}

func (s *PingHandlerSuite) TestPingHandlerStatusOK() {
	handler := PingHandler()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)

	handler(recorder, request)

	resp := recorder.Result()
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(s.T(), constants.ContentTypeJSON, resp.Header.Get(constants.ContentTypeHeader))
	assert.JSONEq(s.T(), `{"response": "pong"}`, recorder.Body.String())
}
