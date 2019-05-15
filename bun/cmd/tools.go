package cmd

import (
	"fmt"
	"github.com/adyatlov/bun/tools"
	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
	"os"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Bun development tool",
	Long:  "Contains subcommands which help to add new file types and checks.",
}


func findFiles(cmd *cobra.Command, args []string) {
	fileTypes, err := tools.FindFiles(bundlePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	y, err := yaml.Marshal(&fileTypes)
	fmt.Println(string(y))
}

func init() {
	rootCmd.AddCommand(toolsCmd)
	var cmd = &cobra.Command{
		Use:   "find-files",
		Short: "Finds all file types in a given bundle",
		Long: "Finds all file types in a given bundle, suggests names, and" +
			" renders it in a YAML format to the stdout.",
		Run: findFiles,
	}
	toolsCmd.AddCommand(cmd)
}
