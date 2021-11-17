package consuldiscovery

// Catalog is a set of functions to find services information
type Catalog interface {
	CatalogServices() CatalogServices
	CatalogServiceByName(name string) CatalogServiceByName
}

// catalogServicesResponse maps GET /v1/catalog/services response
// From API response: {"consul":null,"simple_service":["tag1","tag2"]}
type catalogServicesResponse map[string][]string

// CatalogServices contains the available service names and their tags
type CatalogServices []CatalogService

// CatalogService contains a single available service name and its tags
type CatalogService struct {
	Name string
	Tags []string
}

// CatalogServiceByName contains the nodes composing a service
type CatalogServiceByName []CatalogServiceNode

// CatalogServiceNode describes a single node of a service
// From API response:  {"Node":"drnic.local","Address":"192.168.50.1","ServiceID":"simple_service","ServiceName":"simple_service","ServiceTags":["tag1","tag2"],"ServicePort":6666}
type CatalogServiceNode struct {
	Node        string
	Address     string
	ServiceID   string
	ServiceName string
	ServiceTags []string
	ServicePort uint64

	TaggedAddresses        map[string]string
	NodeMeta               map[string]string
	ID                     string
	Datacenter             string
	ServiceAddress         string
	ServiceTaggedAddresses map[string]ServiceAddress
	ServiceMeta            map[string]string
	Namespace              string
	Partition              string
}

type ServiceAddress struct {
	Address string
	Port    int
}

// CatalogServices returns a list of advertised service names and their tags
func (c *Client) CatalogServices() (result CatalogServices, err error) {
	services := catalogServicesResponse{}
	if err = c.doGET("catalog/services", &services); err != nil {
		return
	}

	// Convert {"name" => ["tags"]} into []CatalogService
	for name, tags := range services {
		service := CatalogService{Name: name, Tags: tags}
		result = append(result, service)
	}

	return
}

// CatalogServiceByName returns a list of nodes composing a service
func (c *Client) CatalogServiceByName(name string) (nodes CatalogServiceByName, err error) {
	err = c.doGET("catalog/service/"+name, &nodes)
	return
}
