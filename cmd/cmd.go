package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0x5w4/kredit-plus/cmd/server"
	"github.com/0x5w4/kredit-plus/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kredit",
	Short: "Kredit service.",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the kredit service.",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(config.LoadConfig())

		if err := s.Run(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
