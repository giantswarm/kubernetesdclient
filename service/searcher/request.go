package searcher

import config "github.com/giantswarm/kubernetesdclient/service/searcher/request"

// Request is the configuration for the service action.
type Request struct {
	Cluster *config.Cluster `json:"cluster"`
}
