package consuldiscovery

import (
	"net/http"
	"time"
)

// Config is used to configure the creation of a client
type Config struct {
	// Address is the address of the Consul server
	Address string

	// Datacenter to use. If not provided, the default agent datacenter is used.
	Datacenter string

	// HTTPClient is the client to use. Default will be
	// used if not provided.
	HTTPClient *http.Client

	// WaitTime limits how long a Watch will block. If not provided,
	// the agent default values will be used.
	WaitTime time.Duration
}

// Client provides a client to Consul for Service data
type Client struct {
	config Config
}

// NewClient returns a new
func NewClient(config *Config) (*Client, error) {
	client := &Client{
		config: *config,
	}
	return client, nil
}

// DefaultConfig returns a default configuration for the client
func DefaultConfig() *Config {
	return &Config{
		Address:    "127.0.0.1:8500",
		HTTPClient: http.DefaultClient,
	}
}

type CatalogServices []CatalogService

type CatalogService struct {
	Name string
	Tags []string
}

type Catalog interface {
	ServiceList() CatalogServices
}

func (client *Client) ServiceList() CatalogServices {
	return CatalogServices{
		CatalogService{Name: "consul"},
		CatalogService{Name: "simple_service"},
	}
}
