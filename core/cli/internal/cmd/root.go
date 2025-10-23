package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azd",
	Short: "Azure Developer CLI",
	Long:  `Azure Developer CLI - A tool for building and deploying modern applications on Azure.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Azure Developer CLI - Use 'azd --help' for more information.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newDeployCommand())
	rootCmd.AddCommand(newInitCommand())
	rootCmd.AddCommand(newProvisionCommand())
}
