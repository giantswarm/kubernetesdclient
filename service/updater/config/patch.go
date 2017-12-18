package config

// Patch is the cluster specific configuration.
type Patch struct {
	ReleaseVersion string   `json:"release_version,omitempty"`
	Workers        []Worker `json:"workers,omitempty"`
}

// DefaultPatch provides a default patch by best effort.
func DefaultPatch() Patch {
	return Patch{
		ReleaseVersion: "",
		Workers:        []Worker{},
	}
}
