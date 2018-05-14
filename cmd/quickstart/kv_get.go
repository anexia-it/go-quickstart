package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/anexia-it/go-quickstart"
)

var cmdKVGet = &cobra.Command{
	Use:   "get <key>",
	Short: "Retrieve consul KV key value",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
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

		var data []byte

		if data, err = kv.Get(key); err != nil {
			return
		}

		cmd.Printf("value: %s\n", string(data))

		return
	},
}

func init() {
	cmdKV.AddCommand(cmdKVGet)
}
