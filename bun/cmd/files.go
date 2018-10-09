package cmd

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"

	"github.com/adyatlov/bun/tools"
	"github.com/spf13/cobra"
)

func init() {
	run := func(cmd *cobra.Command, args []string) {
		fileTypes, err := tools.FindFiles(wd)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		y, err := yaml.Marshal(&fileTypes)
		fmt.Println(string(y))
	}
	var cmd = &cobra.Command{
		Use:   "find-files",
		Short: "Finds all file types in a given bundle",
		Long: "Finds all file types in a given bundle, suggests names, and" +
			" renders it in a YAML format to the stdout.",
		Run: run,
	}
	toolsCmd.AddCommand(cmd)
}
