package quickstart

import "github.com/hashicorp/consul/api"

// ConsulAPI represents a consul API client
type ConsulAPI struct {
	*api.Client
}

// NewConsulAPI returns a new consul API object
func NewConsulAPI(address string) (c *ConsulAPI, err error) {
	var consulClient *api.Client

	clientConfig := &api.Config{
		Address: address,
		Scheme:  "http",
	}

	// Construct the client
	if consulClient, err = api.NewClient(clientConfig); err != nil {
		return
	}

	// Construct the ConsulAPI object
	c = &ConsulAPI{
		Client: consulClient,
	}
	return
}
