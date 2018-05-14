package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/anexia-it/go-quickstart"
)

var cmdKVLockedPut = &cobra.Command{
	Use:   "locked-put <key> <value>",
	Short: "Write consul KV key value while holding the corresponding lock",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var key string

		if len(args) < 2 {
			return fmt.Errorf("Usage: %s\n", cmd.Use)
		}
		key = args[0]
		value := args[1]

		kvAddress, _ := cmd.Flags().GetString("address")
		var kv *quickstart.KV

		if kv, err = getKV(kvAddress); err != nil {
			return
		}

		if err = kv.LockedWrite(key, []byte(value)); err == nil {
			cmd.Println("OK")
		}

		return
	},
}

func init() {
	cmdKV.AddCommand(cmdKVLockedPut)
}
