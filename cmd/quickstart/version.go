package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/anexia-it/go-quickstart"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "print version information",
	Run: func(cmd *cobra.Command, _ []string) {
		logger := quickstart.GetRootLogger().Named("version")
		logger.Info("quickstart version information", zap.String("version", quickstart.VersionString()))
	},
}

func init() {
	cmdRoot.AddCommand(cmdVersion)
}
