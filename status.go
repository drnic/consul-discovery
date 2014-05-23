package consuldiscovery

// Status is a set of functions to get information about the status of the Consul cluster
type Status interface {
	StatusLeader() (string, error)
	StatusPeers() ([]string, error)
}

// StatusLeader gets the Raft leader for the datacenter
func (c *Client) StatusLeader() (leader string, err error) {
	err = c.doGET("status/leader", &leader)
	return
}

// StatusPeers gets the Raft peers for the datacenter
func (c *Client) StatusPeers() (peers []string, err error) {
	err = c.doGET("status/peers", &peers)
	return
}
