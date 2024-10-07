package cmd

import (
	"codechallenge/db"
	"embed"
	"github.com/spf13/cobra"
)

var (
	migrationFS embed.FS
)

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate schemas to database",
	Run: func(_ *cobra.Command, _ []string) {
		dtb := db.New(migrationFS)
		if err := dtb.Migrate(); err != nil {
			panic(err)
		}
	},
}
