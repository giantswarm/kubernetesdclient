package response

// ProtocolPort represents a mapping from a port on the host cluster to the ingress
// of a guest cluster for a given protocol.
type ProtocolPort struct {
	IngressPort int    `json:"ingress_port"`
	LBPort      int    `json:"lb_port"`
	Protocol    string `json:"protocol"`
}

// DefaultProtocolPort provides a default ProtocolPort by best effort.
func DefaultProtocolPort() ProtocolPort {
	return ProtocolPort{
		IngressPort: 0,
		LBPort:      0,
		Protocol:    "",
	}
}
