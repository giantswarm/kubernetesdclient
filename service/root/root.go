package root

import (
	"net/url"

	"github.com/go-resty/resty"
	"golang.org/x/net/context"
)

const (
	Endpoint = "/"
)

// Config represents the configuration used to create a root service.
type Config struct {
	// Dependencies.
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new root service
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		RestClient: resty.New(),

		// Settings.
		URL: nil,
	}

	return newConfig
}

// New creates a new configured root service.
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

func (s *Service) Get(ctx context.Context) (*Response, error) {
	u, err := s.URL.Parse(Endpoint)
	if err != nil {
		return nil, maskAny(err)
	}

	r, err := s.RestClient.R().SetResult(&Response{}).Get(u.String())
	if err != nil {
		return nil, maskAny(err)
	}

	response := r.Result().(*Response)

	return response, nil
}
