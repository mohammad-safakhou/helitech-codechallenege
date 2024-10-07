package cmd

import (
	"codechallenge/internal/gateway/rest"
	"codechallenge/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "serve",
	Short: "launching the http rest listen server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Info("http rest server is starting")
		rest.Start()
	},
}
