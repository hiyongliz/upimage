package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/hiyongliz/upimage/app"

	"github.com/spf13/cobra"
)

var (
	region          string
	namespace       string
	registry        string // Added registry variable
	public          bool
	createNamespace bool
	Opts            app.UpOptions
)

var rootCmd = &cobra.Command{
	Use:     "upimage <image>",
	Short:   "Upload image to Huawei Cloud SWR",
	Long:    `Upload image to Huawei Cloud SWR`,
	Example: `  upimage myregistry/myimage:latest,`,
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		// Initialize options with command line flags
		// Validate required options
		if registry != "swr" && registry != "acr" && registry != "tcr" {
			fmt.Fprintf(os.Stderr, "Invalid registry: %s. Supported values are 'swr', 'acr' or 'tcr'.\n", registry)
			os.Exit(1)
		}

		Opts = app.UpOptions{
			Region:          region,
			Namespace:       namespace,
			Public:          public,
			CreateNamespace: createNamespace,
			Registry:        registry,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		up, err := app.NewUp(Opts)
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
	return fang.Execute(context.Background(), rootCmd)
}

func init() {
	rootCmd.Flags().SortFlags = true
	rootCmd.Flags().StringVarP(&region, "region", "r", "cn-south-1", "Registry region, default is cn-south-1")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Registry namespace, default is default")
	rootCmd.Flags().BoolVar(&createNamespace, "create-namespace", true, "create namespace if not exists, default is true")
	rootCmd.Flags().BoolVar(&public, "public", false, "public image, default is false")
	rootCmd.Flags().StringVarP(&registry, "registry", "g", "swr", "Registry, default is swr")
}
