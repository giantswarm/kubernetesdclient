// Package kubernetesdclient implements business logic to request the
// Kubernetesd API.
package kubernetesdclient

import "github.com/go-resty/resty"

// Config represents the configuration used to create a new client object.
type Config struct {
	// Dependencies.
	RestClient *resty.Client
}

// DefaultConfig provides a default configuration to create a new client object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		RestClient: resty.New(),
	}

	return newConfig
}

// New creates a new configured client object.
func New(config Config) (*Client, error) {
	newClient := &Client{
		Config: config,
	}

	return newClient, nil
}

type Client struct {
	Config
}
