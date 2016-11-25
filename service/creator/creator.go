package creator

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"

	"github.com/giantswarm/kubernetesdclient/service/creator/config"
)

const (
	Endpoint = "/v1/clusters"
)

// Config represents the configuration used to create a creator service.
type Config struct {
	// Dependencies.
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new creator
// service by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		RestClient: resty.New(),

		// Settings.
		URL: nil,
	}

	return newConfig
}

type CreateConfig struct {
	Cluster  *config.Cluster
	Customer *config.Customer
}

// DefaultCreateConfig provides a default configuration to create a new cluster
// resource by best effort.
func DefaultCreateConfig() CreateConfig {
	newConfig := CreateConfig{
		Cluster:  config.NewCluster(),
		Customer: config.NewCustomer(),
	}

	return newConfig
}

// New creates a new configured creator service.
func New(config Config) (*Service, error) {
	newService := &Service{
		Config: config,
	}

	// Dependencies.
	if newService.RestClient == nil {
		return nil, maskAnyf(invalidConfigError, "rest client must not be empty")
	}

	// Settings.
	if newService.URL == nil {
		return nil, maskAnyf(invalidConfigError, "URL must not be empty")
	}

	return newService, nil
}

type Service struct {
	Config
}

func (s *Service) Create(createConfig CreateConfig) (*Response, error) {
	u, err := s.URL.Parse(Endpoint)
	if err != nil {
		return nil, maskAny(err)
	}

	var response *Response
	r, err := s.RestClient.R().SetBody(createConfig).SetResult(response).Post(u.String())
	if err != nil {
		return nil, maskAny(err)
	}

	fmt.Printf("r: %#v\n", r)
	fmt.Printf("response: %#v\n", response)

	return response, nil
}
