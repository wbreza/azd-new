package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newInitCommand creates a new init command.
func newInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init [template]",
		Short: "Initialize a new project",
		Long:  `Initialize a new project from a template.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			template := "default"
			if len(args) > 0 {
				template = args[0]
			}

			fmt.Printf("Initializing project with template: %s\n", template)
			// Initialize project logic would go here
			fmt.Println("Project initialized successfully!")
			return nil
		},
	}
}
