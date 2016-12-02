package deleter

import "github.com/giantswarm/kubernetesdclient/service/deleter/config"

type Request struct {
	Cluster *config.Cluster `json:"cluster"`
}

// DefaultRequest provides a default request by best effort.
func DefaultRequest() Request {
	return Request{
		Cluster: config.NewCluster(),
	}
}
