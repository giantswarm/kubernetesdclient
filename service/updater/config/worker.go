package config

import (
	"github.com/giantswarm/kubernetesdclient/service/updater/config/aws"
	"github.com/giantswarm/kubernetesdclient/service/updater/config/azure"
)

// Worker configures the Kubernetes worker nodes.
type Worker struct {
	AWS     aws.Worker        `json:"aws"`
	Azure   azure.Worker      `json:"azure"`
	CPU     CPU               `json:"cpu"`
	ID      string            `json:"id"`
	Labels  map[string]string `json:"labels"`
	Memory  Memory            `json:"memory"`
	Volumes Volumes           `json:"volumes"`
}

// DefaultWorker provides a default worker configuration by best effort.
func DefaultWorker() Worker {
	return Worker{
		AWS:     aws.DefaultWorker(),
		CPU:     DefaultCPU(),
		ID:      "",
		Labels:  map[string]string{},
		Memory:  DefaultMemory(),
		Volumes: DefaultVolumes(),
	}
}
