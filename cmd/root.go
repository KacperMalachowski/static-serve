package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve fiels in current directory",
	Long:  "Serve static files from current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}
}
