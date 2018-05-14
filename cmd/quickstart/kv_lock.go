package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/anexia-it/go-quickstart"
)

var cmdKVLock = &cobra.Command{
	Use:   "lock",
	Short: "Creates and holds a lock until the command is canceled",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		logger := quickstart.GetRootLogger().Named("watch")
		var key string

		if len(args) < 1 {
			return fmt.Errorf("Usage: %s\n", cmd.Use)
		}
		key = args[0]

		kvAddress, _ := cmd.Flags().GetString("address")
		var kv *quickstart.KV

		if kv, err = getKV(kvAddress); err != nil {
			return
		}

		stopContext, stopWatching := context.WithCancel(context.Background())
		defer stopWatching()

		// Watch for CTRL+C
		sigChan := make(chan os.Signal, 1)
		defer close(sigChan)
		signal.Notify(sigChan, os.Interrupt)

		var lock *api.Lock

		if lock, err = kv.NewLock(key); err != nil {
			return
		}
		defer lock.Destroy()

		go func() {
			// Stop watching when we receive CTRL+C
			defer stopWatching()
			sig := <-sigChan
			logger.Info("Signal received", zap.Stringer("signal", sig))
		}()

		if _, err = lock.Lock(nil); err != nil {
			return
		}

		logger.Info("Lock acquired", zap.String("key", key))
		defer func() {
			lock.Unlock()
			logger.Info("Lock released", zap.String("key", key))
		}()

		// Wait for the context to be canceled
		<-stopContext.Done()
		return
	},
}

func init() {
	cmdKV.AddCommand(cmdKVLock)
}
