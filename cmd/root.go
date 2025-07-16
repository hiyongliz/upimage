package cmd

import (
	"fmt"
	"os"

	"github.com/hiyongliz/upimage/app"

	"github.com/spf13/cobra"
)

var region string

var rootCmd = &cobra.Command{
	Use:     "upimage <image>",
	Short:   "Upload image to Huawei Cloud SWR",
	Long:    `Upload image to Huawei Cloud SWR`,
	Example: `  upimage myregistry/myimage:latest,`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		up, err := app.NewUp(region)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing upimage: %v\n", err)
			os.Exit(1)
		}
		if err := up.Execute(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing upimage: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().SortFlags = true
	rootCmd.Flags().StringVar(&region, "region", "cn-south-1", "SWR region, default is cn-south-1")
}
