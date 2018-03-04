package searcher

import (
	"github.com/giantswarm/kubernetesdclient/service/searcher/response"
)

// Response is the return value of this endpoint.
type Response struct {
	ID            string                  `json:"id"`
	ProtocolPorts []response.ProtocolPort `json:"protocol_ports"`
}
