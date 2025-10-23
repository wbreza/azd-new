package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wbreza/azd-new/core/sdk"
)

// newDeployCommand creates a new deploy command.
func newDeployCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy [app-name]",
		Short: "Deploy your application to Azure",
		Long:  `Deploy your application to Azure using the configured settings.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			appName := args[0]

			// Create SDK client with default config
			config := map[string]string{
				"subscription": "default",
				"region":       "eastus",
			}
			client := sdk.NewClient(config)

			fmt.Printf("Deploying application: %s\n", appName)
			if err := client.Deploy(appName); err != nil {
				return fmt.Errorf("deployment failed: %w", err)
			}

			fmt.Printf("Successfully deployed %s\n", appName)
			return nil
		},
	}
}
