package quickstart

import (
	"context"
	"net/url"
	"time"

	"github.com/hashicorp/consul/api"
)

type WatchResult struct {
	Pair  *api.KVPair
	Error error
}

func (k KV) Watch(ctx context.Context, prefix string) (res chan *WatchResult) {
	res = make(chan *WatchResult, 128)

	go func() {
		defer close(res)

		currentIndex := uint64(0)

		for {
			// This loop is supposed to run until either the context is canceled or the
			// consul API reports an error
			opts := &api.QueryOptions{
				WaitIndex: currentIndex,
				WaitTime:  time.Minute,
			}

			// Retrieve list of changes keys
			kvPairs, meta, err := k.consulKV.List(prefix, opts.WithContext(ctx))

			// Error may be nested in url.Error
			if err != nil {
				if urlErr, isUrlErr := err.(*url.Error); isUrlErr {
					err = urlErr.Err
				}
			}

			if err == context.Canceled {
				// Watching canceled, return without reporting an error
				return
			}

			if err != nil {
				// Report error and stop watching
				res <- &WatchResult{
					Error: err,
				}
				return
			}

			if currentIndex == meta.LastIndex {
				// No changes, move on
				continue
			}

			// Iterate over all changed KV pairs
			for _, kvPair := range kvPairs {
				if kvPair.CreateIndex <= currentIndex && kvPair.ModifyIndex <= currentIndex {
					// Key was unmodified, but child key might have changed
					continue
				}

				// Write changed paiur to channel
				res <- &WatchResult{
					Pair: kvPair,
				}
			}

			// Update currentIndex for the next iteration of this loop
			currentIndex = meta.LastIndex

		}
	}()

	return
}
