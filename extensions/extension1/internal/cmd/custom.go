package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wbreza/azd-new/core/sdk"
)

// newCustomCommand creates a new custom command specific to extension1.
func newCustomCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "custom [resource-name]",
		Short: "Perform custom extension1 operations",
		Long:  `Perform custom operations specific to extension1 functionality.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resourceName := args[0]

			// Use the SDK client
			config := map[string]string{
				"extension": "extension1",
				"mode":      "custom",
			}
			client := sdk.NewClient(config)

			fmt.Printf("Running custom operation on resource: %s\n", resourceName)

			// Extension1 specific logic
			client.SetConfig("operation", "custom-ext1")

			fmt.Printf("Extension1 custom operation completed for %s\n", resourceName)
			return nil
		},
	}
}
