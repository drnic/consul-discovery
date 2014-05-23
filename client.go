package consuldiscovery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"regexp"
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

// CatalogServices contains the available service names and their tags
// From API response: {"consul":null,"simple_service":["tag1","tag2"]}
type CatalogServices []CatalogService

// CatalogService contains a single available service name and its tags
type CatalogService struct {
	Name string
	Tags []string
}

// CatalogServiceNodes contains the nodes composing a service
type CatalogServiceNodes []CatalogServiceNode

// CatalogServiceNode describes a single node of a service
// From API response:  {"Node":"drnic.local","Address":"192.168.50.1","ServiceID":"simple_service","ServiceName":"simple_service","ServiceTags":["tag1","tag2"],"ServicePort":6666}
type CatalogServiceNode struct {
	Node        string
	Address     string
	ServiceID   string
	ServiceName string
	ServiceTags []string
	ServicePort uint64
}

// Catalog is a set of functions to find services information
type Catalog interface {
	ServiceList() CatalogServices
	ServiceNodes(name string) CatalogServiceNodes
}

// ServiceList returns a list of advertised service names and their tags
func (c *Client) ServiceList() CatalogServices {
	// {"consul":null,"simple_service":["tag1","tag2"]}
	return CatalogServices{
		CatalogService{Name: "consul"},
		CatalogService{Name: "simple_service"},
	}
}

// ServiceNodes returns a list of nodes composing a service
func (c *Client) ServiceNodes(name string) (nodes CatalogServiceNodes, err error) {
	url := c.pathURL("catalog/service/" + name)
	req := http.Request{
		Method: "GET",
		URL:    url,
	}
	resp, err := c.config.HTTPClient.Do(&req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return nodes, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}
	dumpedResponse, err := httputil.DumpResponse(resp, true)
	fmt.Println(sanitize(string(dumpedResponse)))

	if err != nil {
		return
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	err = json.Unmarshal(jsonBytes, &nodes)
	return
}

// pathUrl is used to generate the HTTP path for a request
func (c *Client) pathURL(endpoint string) *url.URL {
	url := &url.URL{
		Scheme: "http",
		Host:   c.config.Address,
		Path:   path.Join("/v1/", endpoint),
	}
	if c.config.Datacenter != "" {
		query := url.Query()
		query.Set("dc", c.config.Datacenter)
		url.RawQuery = query.Encode()
	}
	return url
}

const (
	privateDataPlaceholder = "[PRIVATE DATA HIDDEN]"
)

func sanitize(input string) (sanitized string) {
	var sanitizeJSON = func(propertyName string, json string) string {
		re := regexp.MustCompile(fmt.Sprintf(`"%s":"[^"]*"`, propertyName))
		return re.ReplaceAllString(json, fmt.Sprintf(`"%s":"`+privateDataPlaceholder+`"`, propertyName))
	}

	re := regexp.MustCompile(`(?m)^Authorization: .*`)
	sanitized = re.ReplaceAllString(input, "Authorization: "+privateDataPlaceholder)
	re = regexp.MustCompile(`password=[^&]*&`)
	sanitized = re.ReplaceAllString(sanitized, "password="+privateDataPlaceholder+"&")

	sanitized = sanitizeJSON("access_token", sanitized)
	sanitized = sanitizeJSON("refresh_token", sanitized)
	sanitized = sanitizeJSON("token", sanitized)

	return
}
