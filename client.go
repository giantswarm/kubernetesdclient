// Package kubernetesdclient implements business logic to request the
// Kubernetesd API.
package kubernetesdclient

import (
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"gopkg.in/resty.v1"

	"github.com/giantswarm/kubernetesdclient/service/creator"
	"github.com/giantswarm/kubernetesdclient/service/deleter"
	"github.com/giantswarm/kubernetesdclient/service/root"
	"github.com/giantswarm/kubernetesdclient/service/searcher"
	"github.com/giantswarm/kubernetesdclient/service/updater"
)

// Config represents the configuration used to create a new client object.
type Config struct {
	Logger     micrologger.Logger
	RestClient *resty.Client

	Address string
}

// DefaultConfig provides a default configuration to create a new client object
// by best effort.
func DefaultConfig() Config {
	return Config{
		Logger:     nil,
		RestClient: nil,

		Address: "",
	}
}

type Client struct {
	Creator  *creator.Service
	Deleter  *deleter.Service
	Root     *root.Service
	Searcher *searcher.Service
	Updater  *updater.Service
}

// New creates a new configured client object.
func New(config Config) (*Client, error) {
	if config.Address == "" {
		return nil, microerror.Maskf(invalidConfigError, "config.Address must not be empty")
	}

	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var creatorService *creator.Service
	{
		creatorConfig := creator.DefaultConfig()

		creatorConfig.Logger = config.Logger
		creatorConfig.RestClient = config.RestClient
		creatorConfig.URL = u

		creatorService, err = creator.New(creatorConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var deleterService *deleter.Service
	{
		deleterConfig := deleter.DefaultConfig()

		deleterConfig.Logger = config.Logger
		deleterConfig.RestClient = config.RestClient
		deleterConfig.URL = u

		deleterService, err = deleter.New(deleterConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var rootService *root.Service
	{
		rootConfig := root.DefaultConfig()

		rootConfig.Logger = config.Logger
		rootConfig.RestClient = config.RestClient
		rootConfig.URL = u

		rootService, err = root.New(rootConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var updaterService *updater.Service
	{
		updaterConfig := updater.DefaultConfig()

		updaterConfig.Logger = config.Logger
		updaterConfig.RestClient = config.RestClient
		updaterConfig.URL = u

		updaterService, err = updater.New(updaterConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var searcherService *searcher.Service
	{
		c := searcher.Config{
			Logger:     config.Logger,
			RestClient: config.RestClient,
			URL:        u,
		}

		searcherService, err = searcher.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newClient := &Client{
		Creator:  creatorService,
		Deleter:  deleterService,
		Root:     rootService,
		Searcher: searcherService,
		Updater:  updaterService,
	}

	return newClient, nil
}
