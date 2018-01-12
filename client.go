// Package kubernetesdclient implements business logic to request the
// Kubernetesd API.
package kubernetesdclient

import (
	"net/url"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/go-resty/resty"

	"github.com/giantswarm/kubernetesdclient/service/creator"
	"github.com/giantswarm/kubernetesdclient/service/deleter"
	"github.com/giantswarm/kubernetesdclient/service/root"
	"github.com/giantswarm/kubernetesdclient/service/updater"
)

const (
	// DefaultRetryCount is the default number of times to retry a failed network call.
	DefaultRetryCount = 5

	// DefaultTimeout is the default timeout for network calls.
	DefaultTimeout = 5 * time.Second
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
	newRestyClient := resty.New().
		SetTimeout(DefaultTimeout).
		SetRetryCount(DefaultRetryCount)

	newConfig := Config{
		// Dependencies.
		RestClient: newRestyClient,

		// Settings.
		Address: "http://127.0.0.1:8080",
	}

	return newConfig
}

type Client struct {
	Creator *creator.Service
	Deleter *deleter.Service
	Root    *root.Service
	Updater *updater.Service
}

// New creates a new configured client object.
func New(config Config) (*Client, error) {
	// Dependencies.
	if config.RestClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "rest client must not be empty")
	}

	// Settings.
	if config.Address == "" {
		return nil, microerror.Maskf(invalidConfigError, "address must not be empty")
	}

	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	creatorConfig := creator.DefaultConfig()
	creatorConfig.RestClient = config.RestClient
	creatorConfig.URL = u
	newCreatorService, err := creator.New(creatorConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	deleterConfig := deleter.DefaultConfig()
	deleterConfig.RestClient = config.RestClient
	deleterConfig.URL = u
	newDeleterService, err := deleter.New(deleterConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	rootConfig := root.DefaultConfig()
	rootConfig.RestClient = config.RestClient
	rootConfig.URL = u
	newRootService, err := root.New(rootConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	updaterConfig := updater.DefaultConfig()
	updaterConfig.RestClient = config.RestClient
	updaterConfig.URL = u
	newUpdaterService, err := updater.New(updaterConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	newClient := &Client{
		Creator: newCreatorService,
		Deleter: newDeleterService,
		Root:    newRootService,
		Updater: newUpdaterService,
	}

	return newClient, nil
}
