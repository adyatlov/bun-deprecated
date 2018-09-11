package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/adyatlov/bun"
	"github.com/spf13/cobra"
)

var bundlePath string
var bundle *bun.Bundle
var printLong = false

var rootCmd = &cobra.Command{
	Use:   "bun",
	Short: "DC/OS diagnostics bundle analysis tool",
	Long: "Bun extracts useful facts from hundreds of files in the DC/OS diagnostics bundle\n" +
		"and searches for some common problems of the DC/OS cluster.\n" +
		"\nSpecify a subcommand to run a specific check, e.g. `bun health`\n" +
		"or run all the available checks by not specifying any, i.e. `bun`.\n" +
		"\nMore information is available at https://github.com/adyatlov/bun",
	PreRun: preRun,
	Run:    runCheck,
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while detecting a working directory: %v\n", err.Error())
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&bundlePath, "path", "p", wd,
		"path to the bundle directory")
	rootCmd.PersistentFlags().BoolVarP(&printLong, "long", "l", false,
		"print details")
	// Adding registered checks as commands.
	for _, check := range bun.Checks() {
		run := func(cmd *cobra.Command, args []string) {
			check.Run(*bundle)
			printReport(check)
			return
		}
		var cmd = &cobra.Command{
			Use:    check.Name,
			Short:  check.Description,
			Long:   check.Description,
			PreRun: preRun,
			Run:    run,
		}
		rootCmd.AddCommand(cmd)
		rootCmd.ValidArgs = append(rootCmd.ValidArgs, check.Name)
		rootCmd.PreRun = preRun
	}
}

func preRun(cmd *cobra.Command, args []string) {
	if bundle != nil {
		return
	}
	b, err := bun.NewBundle(bundlePath)
	if err != nil {
		fmt.Printf("Cannot find a bundle: %v\n", err.Error())
		os.Exit(1)
	}
	bundle = &b
}

func runCheck(cmd *cobra.Command, args []string) {
	if err := cobra.OnlyValidArgs(cmd, args); err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Run '%v --help' for usage.\n", cmd.CommandPath())
		os.Exit(1)
	}
	checks := bun.Checks()
	sort.Slice(checks, func(i, j int) bool {
		return checks[i].Name < checks[j].Name
	})
	for _, check := range checks {
		check.Run(*bundle)
		printReport(check)
	}
}

// Execute starts Bun.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
