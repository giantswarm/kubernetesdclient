package config

import "github.com/giantswarm/versionbundle"

// Patch is the cluster specific configuration.
type Patch struct {
	ReleaseVersion string                 `json:"release_version,omitempty"`
	Scaling        Scaling                `json:"scaling,omitempty"`
	VersionBundles []versionbundle.Bundle `json:"version_bundles,omitempty"`
	Workers        []Worker               `json:"workers,omitempty"`
}

// DefaultPatch provides a default patch by best effort.
func DefaultPatch() Patch {
	return Patch{
		ReleaseVersion: "",
		Scaling:        Scaling{},
		VersionBundles: []versionbundle.Bundle{},
		Workers:        []Worker{},
	}
}
