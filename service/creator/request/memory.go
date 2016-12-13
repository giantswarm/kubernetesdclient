package request

// Memory configures the machine memory.
type Memory struct {
	SizeGB string `json:"size_gb"`
}

// DefaultMemory provides a default ram configuration by best effort.
func DefaultMemory() *Memory {
	return &Memory{
		SizeGB: "",
	}
}
