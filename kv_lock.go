package quickstart

import (
	"github.com/hashicorp/consul/api"
)

// NewLock creates a new lock
//
// The lock key is constructed by prefixing the target key with a locks/ prefix
func (k KV) NewLock(key string) (*api.Lock, error) {
	return k.client.LockKey("locks/" + key)
}

// LockedWrite writes to a key after obtaining the corresponding lock
func (k KV) LockedWrite(key string, value []byte) (err error) {
	var lock *api.Lock

	// Create the lock
	if lock, err = k.NewLock(key); err != nil {
		return
	}
	defer lock.Destroy()

	// Acquire the lock
	if _, err = lock.Lock(nil); err != nil {
		return
	}
	defer lock.Unlock()

	return k.Put(key, value)
}
