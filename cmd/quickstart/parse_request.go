package main

import (
	"fmt"
	"os"

	"github.com/anexia-it/go-quickstart"
	"github.com/spf13/cobra"
)

var cmdParseRequest = &cobra.Command{
	Use:   "parse-request <path>",
	Short: "Parses a request",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		logger := quickstart.GetRootLogger().Named("parse_request")
		if len(args) != 1 {
			return fmt.Errorf("usage: %s", cmd.Use)
		}

		// Try to open the file
		var f *os.File
		if f, err = os.Open(args[0]); err != nil {
			return
		}
		// Ensure the file is closed when we leave this function scope
		defer f.Close()

		var r *quickstart.Request
		if r, err = quickstart.DecodeRequest(f); err != nil {
			return
		}

		logger.Sugar().Infof("Request: %s\n", r)

		return
	},
}

func init() {
	cmdRoot.AddCommand(cmdParseRequest)
}
