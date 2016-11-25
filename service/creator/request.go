package creator

import "github.com/giantswarm/kubernetesdclient/service/creator/config"

type Request struct {
	Cluster  config.Cluster
	Customer config.Customer
}
