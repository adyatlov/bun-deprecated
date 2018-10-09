package cmd

import (
	"github.com/spf13/cobra"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Bun development tool",
	Long:  "Contains subcommands which help to add new file types and checks.",
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}
