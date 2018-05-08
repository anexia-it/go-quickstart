package main

import (
	"github.com/spf13/cobra"

	"github.com/anexia-it/go-quickstart"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "print version information",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Printf("quickstart v%s\n", quickstart.VersionString())
	},
}

func init() {
	cmdRoot.AddCommand(cmdVersion)
}
