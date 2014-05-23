package consuldiscovery

// Health is a set of functions for the health of services
type Health interface {
	HealthByNode(nodeName string) ([]HealthServiceCheck, error)
	HealthByService(serviceName string) (HealthServiceNodes, error)
	HealthByState(state string) ([]HealthServiceCheck, error)
}

// [{"Node":{"Node":"drnic.local","Address":"192.168.50.1"},
//  "Service":{"ID":"simple_service","Service":"simple_service","Tags":["tag1","tag2"],"Port":6666},
//  "Checks":[{"Node":"drnic.local","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"","ServiceID":"","ServiceName":""}]}]

// HealthServiceNodes summarizes the health checks for all Nodes for a single Service
type HealthServiceNodes []HealthServiceForNode

// HealthServiceForNode summarizes the health checks for a single Nodes for a single Service
type HealthServiceForNode struct {
	Node    HealthServiceNode
	Service HealthServiceService
	Checks  []HealthServiceCheck
}

// HealthServiceNode indicates a server/node being described by HealthServiceNodes
type HealthServiceNode struct {
	Node    string
	Address string
}

// HealthServiceService indicates a service being described by HealthServiceNodes
type HealthServiceService struct {
	ServiceID   string   `json:"ID"`
	ServiceName string   `json:"Service"`
	ServiceTags []string `json:"Tags"`
	ServicePort uint64   `json:"Port"`
}

// HealthServiceCheck contains a current health check result
type HealthServiceCheck struct {
	Node        string
	CheckID     string
	Name        string
	Status      string
	Nodes       string
	Output      string
	ServiceID   string
	ServiceName string
}

// HealthByNode returns the health checks for a specific node
func (c *Client) HealthByNode(nodeName string) (result []HealthServiceCheck, err error) {
	if err = c.doGET("health/node/"+nodeName, &result); err != nil {
		return
	}
	return
}

// HealthByService returns a list of advertised service names and their tags
func (c *Client) HealthByService(serviceName string) (result HealthServiceNodes, err error) {
	if err = c.doGET("health/service/"+serviceName, &result); err != nil {
		return
	}
	return
}

// HealthByState returns the health checks with a specific state
func (c *Client) HealthByState(state string) (checks []HealthServiceCheck, err error) {
	if err = c.doGET("health/state/"+state, &checks); err != nil {
		return
	}
	return
}
