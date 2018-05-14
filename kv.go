package quickstart

import (
  "github.com/hashicorp/consul/api"
)

type KV struct {
  consulClient *api.Client
}

// NewKV constructs new KV connection
func NewKV(address string) (k *KV, err error) {
	conf := &api.Config{
		Address: address,
		Scheme: "http",
	}
	var c *api.Client
	if c, err = api.NewClient(conf); err != nil {
          return
	}

	k = &KV{
		consulClient: c,
	}
	return
}
