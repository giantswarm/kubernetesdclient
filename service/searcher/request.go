package searcher

import "github.com/giantswarm/kubernetesdclient/service/deleter/config"

// Request is the configuration for the service action.
type Request struct {
	Cluster *config.Cluster `json:"cluster"`
}
