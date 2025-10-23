package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azd-ext2",
	Short: "Azure Dev Extension 2",
	Long:  `Azure Dev Extension 2 - A monitoring and analytics extension for the Azure Developer CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Azure Dev Extension 2 - Use 'azd-ext2 --help' for more information.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newMonitorCommand())
	rootCmd.AddCommand(newAnalyzeCommand())
}
