package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	PORT = "8080"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve files in current directory",
	Long:  "Serve static files from current directory",
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir("."))

		http.Handle("/", fs)

		port := fmt.Sprintf(":%s", PORT)

		log.Printf("Listening on %s...", port)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatalf("Fail to serve files: %s", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&PORT, "port", "p", "8080", "Set port for web server to listen on")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}
}
