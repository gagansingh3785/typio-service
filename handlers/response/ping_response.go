package response

import (
	"net/http"
)

type PingResponse struct{}

func NewPingResponse() *PingResponse {
	return &PingResponse{}
}

func (resp *PingResponse) GetStatus() int {
	return http.StatusOK
}
