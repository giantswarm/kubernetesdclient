package creator

import "github.com/giantswarm/kubernetesdclient/service/creator/config"

type Request struct {
	Cluster  *config.Cluster  `json:"cluster"`
	Customer *config.Customer `json:"customer"`
}

// DefaultRequest provides a default request by best effort.
func DefaultRequest() Request {
	return Request{
		Cluster:  config.NewCluster(),
		Customer: config.NewCustomer(),
	}
}
