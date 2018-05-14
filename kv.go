package quickstart

import (
	"github.com/hashicorp/consul/api"
)

type KV struct {
	consulKV *api.KV
	client   *api.Client
}

// Get retrieves a KV key
func (k KV) Get(key string) (data []byte, err error) {
	var kvPair *api.KVPair
	if kvPair, _, err = k.consulKV.Get(key, nil); err != nil {
		return
	} else if kvPair != nil {

		data = kvPair.Value
	}
	return
}

// Put writes a key to the consul KV
func (k KV) Put(key string, data []byte) (err error) {
	kvPair := &api.KVPair{
		Key:   key,
		Value: data,
	}

	_, err = k.consulKV.Put(kvPair, nil)
	return
}

// Delete removes a key from consul KV
func (k KV) Delete(key string) (err error) {
	_, err = k.consulKV.Delete(key, nil)
	return
}

// NewKV constructs new KV connection
func NewKV(address string) (k *KV, err error) {
	conf := &api.Config{
		Address: address,
		Scheme:  "http",
	}
	var c *api.Client
	if c, err = api.NewClient(conf); err != nil {
		return
	}

	k = &KV{
		client:   c,
		consulKV: c.KV(),
	}
	return
}
