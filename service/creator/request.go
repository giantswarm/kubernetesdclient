package creator

import (
	"time"

	"github.com/giantswarm/versionbundle"

	"github.com/giantswarm/kubernetesdclient/service/creator/request"
)

// Request is the configuration for the service action.
type Request struct {
	APIEndpoint    string                 `json:"api_endpoint"`
	CreateDate     time.Time              `json:"create_date"`
	ID             string                 `json:"id"`
	Name           string                 `json:"name,omitempty"`
	Owner          string                 `json:"owner,omitempty"`
	VersionBundles []versionbundle.Bundle `json:"version_bundles,omitempty"`
	Vault          request.Vault          `json:"vault,omitempty"`
	Workers        []request.Worker       `json:"workers,omitempty"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		APIEndpoint:    "",
		CreateDate:     time.Time{},
		ID:             "",
		Name:           "",
		Owner:          "",
		VersionBundles: []versionbundle.Bundle{},
		Vault:          request.DefaultVault(),
		Workers:        []request.Worker{},
	}
}
