package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/adyatlov/bun"
	_ "github.com/adyatlov/bun/import"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check DC/OS diagnostics bundle for possible problems",
	Long: `Check DC/OS diagnostics bundle for possible problems.

Specify a subcommand to run a specific check, e.g.` + " `bun check health`." +
		`
Or run all the available checks by not spcifying any: ` + "`bun check`.",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	if err := cobra.NoArgs(cmd, args); err != nil {
		cmd.Help()
		return
	}
	checks := bun.Checks()
	sort.Slice(checks, func(i, j int) bool {
		return checks[i].Name < checks[j].Name
	})
	ctx := context.Background()
	prog := make(chan bun.NamedProgress)
	for _, check := range checks {
		report, err := bun.RunCheck(ctx,
			check.Name,
			bundle,
			prog)
		if err != nil {
			fmt.Printf("Error while running check %v: %v\n", check.Name, err.Error())
		}
		printReport(report)
	}
}

const printLong = false

func printReport(r bun.Report) {
	fmt.Printf("%v: %v - %v\n", r.Status, r.Name, r.Short)
	if r.Status == bun.SProblem || printLong {
		fmt.Printf("Details:\n%v\n", r.Long)
	}
	if len(r.Errors) > 0 {
		fmt.Printf("Errors: \n")
		for i, err := range r.Errors {
			fmt.Printf("%v: %v\n", i+1, err)
		}
	}
}

var bundle bun.Bundle

func init() {
	rootCmd.AddCommand(checkCmd)

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while detecting a working directory: %v\n", err.Error())
	}
	bundle, err = bun.NewBundle(path)
	if err != nil {
		log.Fatalf("Error while identifying basic bundle parameters: %v\n", err.Error())
	}
	ctx := context.Background()
	prog := make(chan bun.NamedProgress)
	for _, check := range bun.Checks() {
		run := func(cmd *cobra.Command, args []string) {
			report, err := bun.RunCheck(ctx,
				cmd.Name(),
				bundle,
				prog)
			if err != nil {
				fmt.Printf("Error while running check %v: %v\n", check.Name, err.Error())
			}
			printReport(report)
		}
		var cmd = &cobra.Command{
			Use:   check.Name,
			Short: check.Description,
			Long:  check.Name,
			Run:   run,
		}
		checkCmd.AddCommand(cmd)
	}
}
