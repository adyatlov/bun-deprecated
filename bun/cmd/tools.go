package cmd

import (
	"bytes"
	"fmt"
	"github.com/adyatlov/bun/tools"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	escape, err := cmd.Flags().GetBool("escape");
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if escape {
		y = bytes.ReplaceAll(y, []byte("`"), []byte("`+ \"`\" +`"))
	}
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
	cmd.Flags().BoolP("escape", "e", false, "Escape back ticks for using in the files_yaml.go")
	toolsCmd.AddCommand(cmd)
}
