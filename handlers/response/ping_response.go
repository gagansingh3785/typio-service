package response

import (
	"net/http"
)

type PingResponse struct {
	Response string `json:"response"`
}

func NewPingResponse() *PingResponse {
	return &PingResponse{
		Response: "pong",
	}
}

func (resp *PingResponse) GetStatus() int {
	return http.StatusOK
}
