package cmd

import (
	"embed"
	"fmt"
	"os"
	"time"

	"codechallenge/config"
	"codechallenge/logger"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A generator for Cobra based Applications",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute(mf embed.FS) {
	migrationFS = mf
	err := os.Setenv("TZ", time.UTC.String())
	if err != nil {
		panic(err)
	}

	if err = rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is .)")
	cobra.OnInitialize(
		initConfig,
		initLogger, // logger should come after config
	)
	rootCmd.AddCommand(migrateCommand)
}

func initConfig() {
	config.LoadConfig(cfgFile)
}

func initLogger() {
	logger.LoadLogger(config.AppConfig.General.LogLevel)
}
