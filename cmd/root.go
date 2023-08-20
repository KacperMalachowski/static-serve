package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve fiels in current directory",
	Long:  "Serve static files from current directory",
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir("."))

		http.Handle("/", fs)

		log.Print("Listening on :3000...")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatalf("Fail to serve files: %s", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}
}
