package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wbreza/azd-new/core/sdk"
)

// NewCustomCommand creates a new custom command specific to extension1.
func NewCustomCommand() *cobra.Command {
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

// NewIntegrateCommand creates a new integrate command.
func NewIntegrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "integrate [service]",
		Short: "Integrate with external services",
		Long:  `Integrate your application with external services using extension1.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service := args[0]
			
			fmt.Printf("Integrating with service: %s\n", service)
			
			// Integration logic specific to extension1
			fmt.Printf("Successfully integrated with %s using extension1\n", service)
			return nil
		},
	}
}