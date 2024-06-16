package cmd

import (
	"GoMVCStarterKit/database/migrations"
	"GoMVCStarterKit/src/server"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Start() {
	rootCmd := &cobra.Command{}

	cmd := []*cobra.Command{
		{
			Use:   "db:migrate",
			Short: "Run migration",
			Run: func(cmd *cobra.Command, args []string) {
				migrations.Migrate()
			},
		},
		{
			Use:   "HTTP Server",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				server.Http()
			},
		},
	}
	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error run command using cobra : %v", err)
		os.Exit(1)
	}
}
