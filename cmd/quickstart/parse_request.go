package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/anexia-it/go-human"
	"github.com/spf13/cobra"

	"github.com/anexia-it/go-quickstart"
)

// encoder is implemented by json.Encoder and human.Encoder
type encoder interface {
	Encode(o interface{}) error
}

var cmdParseRequest = &cobra.Command{
	Use:   "parse-request <path>",
	Short: "Parses a request",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s", cmd.Use)
		}

		// Check if we want JSON output
		wantJSON, _ := cmd.Flags().GetBool("json")

		var enc encoder
		if wantJSON {
			jsonEncoder := json.NewEncoder(cmd.OutOrStdout())
			jsonEncoder.SetIndent("", "  ")
			enc = jsonEncoder
		} else {
			enc, _ = human.NewEncoder(cmd.OutOrStdout())
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

		err = enc.Encode(r)

		return
	},
}

func init() {
	cmdParseRequest.Flags().BoolP("json", "j", false, "JSON output")
	cmdRoot.AddCommand(cmdParseRequest)
}
