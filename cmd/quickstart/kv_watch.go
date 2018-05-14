package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/anexia-it/go-quickstart"
)

var cmdKVWatch = &cobra.Command{
	Use:   "watch <prefix>",
	Short: "Watch consul KV key prefix for changes",
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

		go func() {
			// Stop watching when we receive CTRL+C
			defer stopWatching()
			sig := <-sigChan
			logger.Info("Signal received", zap.Stringer("signal", sig))
		}()

		res := kv.Watch(stopContext, key)

		for change := range res {
			if change.Error != nil {
				err = change.Error
				logger.Error("watching returned error", zap.Error(err))
				return
			}

			logger.Info("value changed", zap.String("key", change.Pair.Key), zap.String("value", string(change.Pair.Value)))
		}

		return
	},
}

func init() {
	cmdKV.AddCommand(cmdKVWatch)
}
