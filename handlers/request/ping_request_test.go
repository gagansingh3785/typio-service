package request

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PingRequestSuite struct {
	suite.Suite
}

func TestPingRequestSuite(t *testing.T) {
	suite.Run(t, new(PingRequestSuite))
}

func (s *PingRequestSuite) TestNewPingRequest() {
	req, err := NewPingRequest(context.Background(), nil)

	s.NoError(err)

	s.NotNil(req)
}

func (s *PingRequestSuite) TestValidate() {
	req := &PingRequest{}

	s.NoError(req.Validate())
}
