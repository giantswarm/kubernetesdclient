package searcher

import (
	"github.com/giantswarm/kubernetesd/server/endpoint/searcher/response"
)

// Response is the return value of this endpoint.
type Response struct {
	ID            string                  `json:"id"`
	ProtocolPorts []response.ProtocolPort `json:"protocol_ports"`
}

// DefaultResponse provides a default response by best effort.
func DefaultResponse() *Response {
	return &Response{
		ID:            "",
		ProtocolPorts: []response.ProtocolPort{},
	}
}
