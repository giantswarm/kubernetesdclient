package deleter

import "github.com/giantswarm/kubernetesdclient/service/deleter/config"

type Request struct {
	Cluster *config.Cluster `json:"cluster"`
}

// DefaultDeleteConfig provides a default configuration to create a new cluster
// resource by best effort.
func DefaultDeleteConfig() Request {
	newConfig := Request{
		Cluster: config.NewCluster(),
	}

	return newConfig
}
