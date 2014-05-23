package consuldiscovery

// Health is a set of functions for the health of services
type Health interface {
	HealthByNode(nodeName string) ([]HealthCheck, error)
	HealthByService(serviceName string) (HealthNodes, error)
	HealthByState(state string) ([]HealthCheck, error)
}

// [{"Node":{"Node":"drnic.local","Address":"192.168.50.1"},
//  "Service":{"ID":"simple_service","Service":"simple_service","Tags":["tag1","tag2"],"Port":6666},
//  "Checks":[{"Node":"drnic.local","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"","ServiceID":"","ServiceName":""}]}]

// HealthNodes summarizes the health checks for all Nodes for a single Service
type HealthNodes []HealthForNode

// HealthForNode summarizes the health checks for a single Nodes for a single Service
type HealthForNode struct {
	Node    HealthNode
	Service HealthService
	Checks  []HealthCheck
}

// HealthNode indicates a server/node being described by HealthNodes
type HealthNode struct {
	Node    string
	Address string
}

// HealthService indicates a service being described by HealthNodes
type HealthService struct {
	ServiceID   string   `json:"ID"`
	ServiceName string   `json:"Service"`
	ServiceTags []string `json:"Tags"`
	ServicePort uint64   `json:"Port"`
}

// HealthCheck contains a current health check result
type HealthCheck struct {
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
func (c *Client) HealthByNode(nodeName string) (result []HealthCheck, err error) {
	err = c.doGET("health/node/"+nodeName, &result)
	return
}

// HealthByService returns a list of advertised service names and their tags
func (c *Client) HealthByService(serviceName string) (result HealthNodes, err error) {
	err = c.doGET("health/service/"+serviceName, &result)
	return
}

// HealthByState returns the health checks with a specific state
func (c *Client) HealthByState(state string) (checks []HealthCheck, err error) {
	err = c.doGET("health/state/"+state, &checks)
	return
}
