package creator

import (
	"time"

	"github.com/giantswarm/versionbundle"

	"github.com/giantswarm/kubernetesdclient/service/creator/request"
)

// Request is the configuration for the service action.
type Request struct {
	APIEndpoint       string                 `json:"api_endpoint"`
	AvailabilityZones int                    `json:"availability_zones"`
	CreateDate        time.Time              `json:"create_date"`
	ID                string                 `json:"id"`
	Masters           []request.Master       `json:"masters,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Owner             string                 `json:"owner,omitempty"`
	Region            string                 `json:"region,omitempty"`
	Scaling           request.Scaling        `json:"scaling,omitempty"`
	VersionBundles    []versionbundle.Bundle `json:"version_bundles,omitempty"`
	Workers           []request.Worker       `json:"workers,omitempty"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		APIEndpoint:       "",
		AvailabilityZones: 0,
		CreateDate:        time.Time{},
		ID:                "",
		Masters:           []request.Master{},
		Name:              "",
		Owner:             "",
		Region:            "",
		Scaling:           request.Scaling{},
		VersionBundles:    []versionbundle.Bundle{},
		Workers:           []request.Worker{},
	}
}
