package request

import (
	"github.com/giantswarm/kubernetesdclient/service/creator/request/aws"
	"github.com/giantswarm/kubernetesdclient/service/creator/request/azure"
)

// Master configures the Kubernetes master nodes.
type Master struct {
	AWS     aws.Master   `json:"aws"`
	Azure   azure.Master `json:"azure"`
	CPU     CPU          `json:"cpu"`
	ID      string       `json:"id"`
	Memory  Memory       `json:"memory"`
	Storage Storage      `json:"storage"`
}

// DefaultMaster provides a default master configuration by best effort.
func DefaultMaster() Master {
	return Master{
		AWS:     aws.DefaultMaster(),
		CPU:     DefaultCPU(),
		ID:      "",
		Memory:  DefaultMemory(),
		Storage: DefaultStorage(),
	}
}
