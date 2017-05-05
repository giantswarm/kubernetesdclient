package request

import (
	"github.com/giantswarm/kubernetesdclient/service/creator/request/aws"
)

// Master configures the Kubernetes master nodes.
type Master struct {
	CPU     CPU        `json:"cpu"`
	Memory  Memory     `json:"memory"`
	Storage Storage    `json:"storage"`
	AWS     aws.Master `json:"aws"`
}

// DefaultMaster provides a default master configuration by best effort.
func DefaultMaster() Master {
	return Master{
		CPU:     DefaultCPU(),
		Memory:  DefaultMemory(),
		Storage: DefaultStorage(),
		AWS:     aws.DefaultMaster(),
	}
}
