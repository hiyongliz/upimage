package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version information - will be set by build flags
var (
	Version   = "dev"
	BuildTime = "unknown"
	CommitSHA = "unknown"
)

func init() {
	// Add version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("upimage version %s\n", Version)
			fmt.Printf("Build time: %s\n", BuildTime)
			fmt.Printf("Commit: %s\n", CommitSHA)
		},
	}
	rootCmd.AddCommand(versionCmd)
}
