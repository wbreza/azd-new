package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wbreza/azd-new/extensions/extension1/internal/commands"
)

var rootCmd = &cobra.Command{
	Use:   "azd-ext1",
	Short: "Azure Dev Extension 1",
	Long:  `Azure Dev Extension 1 - A sample extension for the Azure Developer CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Azure Dev Extension 1 - Use 'azd-ext1 --help' for more information.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(commands.NewCustomCommand())
	rootCmd.AddCommand(commands.NewIntegrateCommand())
}