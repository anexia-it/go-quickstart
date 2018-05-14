package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/anexia-it/go-quickstart"
)

var cmdKVDelete = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete consul KV key",
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

		if err = kv.Delete(key); err == nil {
			cmd.Println("OK")
		}

		return
	},
}

func init() {
	cmdKV.AddCommand(cmdKVDelete)
}
