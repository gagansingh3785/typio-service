package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PingResponseSuite struct {
	suite.Suite
}

func TestPingResponseSuite(t *testing.T) {
	suite.Run(t, new(PingResponseSuite))
}

func (s *PingResponseSuite) TestNewPingResponse() {
	resp := NewPingResponse()

	s.NotNil(resp)

	s.Equal("pong", resp.Response)
}

func (s *PingResponseSuite) TestGetStatus() {
	resp := NewPingResponse()

	s.Equal(http.StatusOK, resp.GetStatus())
}
