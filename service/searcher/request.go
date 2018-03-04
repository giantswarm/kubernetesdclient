package searcher

import "github.com/giantswarm/kubernetesdclient/service/searcher/config"

// Request is the configuration for the service action.
type Request struct {
	Cluster *config.Cluster `json:"cluster"`
}
