// Package kubernetesdclient implements business logic to request the
// Kubernetesd API.
package kubernetesdclient

import (
	"net/url"

	"github.com/go-resty/resty"

	"github.com/giantswarm/kubernetesdclient/service/creator"
	"github.com/giantswarm/kubernetesdclient/service/root"
)

// Config represents the configuration used to create a new client object.
type Config struct {
	// Dependencies.
	RestClient *resty.Client

	// Settings.
	Address string
}

// DefaultConfig provides a default configuration to create a new client object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		RestClient: resty.New(),

		// Settings.
		Address: "http://127.0.0.1:8080",
	}

	return newConfig
}

// New creates a new configured client object.
func New(config Config) (*Client, error) {
	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, maskAny(err)
	}

	creatorConfig := creator.DefaultConfig()
	creatorConfig.RestClient = config.RestClient
	creatorConfig.URL = u
	newCreatorService, err := creator.New(creatorConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	rootConfig := root.DefaultConfig()
	rootConfig.RestClient = config.RestClient
	rootConfig.URL = u
	newRootService, err := root.New(rootConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	newClient := &Client{
		Creator: newCreatorService,
		Root:    newRootService,
	}

	return newClient, nil
}

type Client struct {
	Creator *creator.Service
	Root    *root.Service
}
