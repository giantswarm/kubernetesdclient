package searcher

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"
	"golang.org/x/net/context"
)

const (
	// Endpoint is the API endpoint of the service this client action interacts
	// with.
	Endpoint = "/v1/clusters/%s/ingress-ports/"
	// Name is the service name being implemented. This can be used for e.g.
	// logging.
	Name = "cluster/searcher"
)

// Config represents the configuration used to create a creator service.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

type Service struct {
	logger     micrologger.Logger
	restClient *resty.Client
	url        *url.URL
}

// New creates a new configured updater service.
func New(config Config) (*Service, error) {
	// Dependencies.
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.RestClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.RestClient must not be empty", config)
	}

	// Settings.
	if config.URL == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.URL must not be empty", config)
	}

	newService := &Service{
		logger:     config.Logger,
		restClient: config.RestClient,
		url:        config.URL,
	}

	return newService, nil
}

func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	u, err := s.url.Parse(fmt.Sprintf(Endpoint, request.Cluster.ID))
	if err != nil {
		return nil, microerror.Mask(err)
	}

	s.logger.Log("debug", fmt.Sprintf("sending GET request to %s", u.String()), "service", Name)
	r, err := s.restClient.R().SetBody(request).SetResult(Response{}).Get(u.String())
	if err != nil {
		return nil, microerror.Mask(err)
	}
	s.logger.Log("debug", fmt.Sprintf("received status code %d", r.StatusCode()), "service", Name)

	if r.StatusCode() == http.StatusNotFound {
		return nil, microerror.Mask(notFoundError)
	} else if r.StatusCode() != http.StatusOK {
		return nil, microerror.Mask(fmt.Errorf(string(r.Body())))
	}

	response := r.Result().(*Response)

	return response, nil
}
