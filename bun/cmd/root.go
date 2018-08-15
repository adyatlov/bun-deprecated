package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var bundlePath string

var rootCmd = &cobra.Command{
	Use:   "bun",
	Short: "DC/OS diagnostics bundle analysis tool",
	Long: "Bun extracts useful facts from hundreds of files in a DC/OS diagnostics bundle" +
		" and searches for some common problems of a DC/OS cluster." +
		"\nMore information is available at https://github.com/adyatlov/bun",
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while detecting a working directory: %v\n", err.Error())
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&bundlePath, "path", "p", wd, "path to the bundle directory")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
