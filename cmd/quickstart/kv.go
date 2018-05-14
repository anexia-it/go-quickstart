package main

import (
	"github.com/spf13/cobra"

	"github.com/anexia-it/go-quickstart"
)

var cmdKV = &cobra.Command{
	Use:   "kv",
	Short: "Consul KV commands",
}

func init() {
	cmdKV.PersistentFlags().StringP("address", "a", "127.0.0.1:8500", "Consul KV address")
	cmdRoot.AddCommand(cmdKV)
}

func getKV(address string) (*quickstart.KV, error) {
	return quickstart.NewKV(address)
}
