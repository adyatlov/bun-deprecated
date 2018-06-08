package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/adyatlov/bun"
	_ "github.com/adyatlov/bun/import"
	"github.com/spf13/cobra"
)

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

var bundle *bun.Bundle

func preRun(cmd *cobra.Command, args []string) {
	if bundle != nil {
		return
	}
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while detecting a working directory: %v\n", err.Error())
		os.Exit(1)
	}
	b, err := bun.NewBundle(path)
	if err != nil {
		fmt.Printf("Error while identifying basic bundle parameters: %v\n", err.Error())
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
		report, err := bun.RunCheckSimple(check.Name, *bundle)
		if err != nil {
			fmt.Printf("Error while running check %v: %v", check.Name, err.Error())
		}
		printReport(report)
	}
	return
}

func init() {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check DC/OS diagnostics bundle for possible problems",
		Long: `Check DC/OS diagnostics bundle for possible problems.

Specify a subcommand to run a specific check, e.g.` + " `bun check health`." +
			`
Or run all the available checks by not spcifying any, i.e.` + " `bun check`.",
		PreRun: preRun,
		Run:    runCheck,
	}

	for _, check := range bun.Checks() {
		run := func(cmd *cobra.Command, args []string) {
			report, err := bun.RunCheckSimple(cmd.Name(), *bundle)
			if err != nil {
				fmt.Println(err.Error())
			}
			printReport(report)
			return
		}
		var cmd = &cobra.Command{
			Use:    check.Name,
			Short:  check.Description,
			Long:   check.Description,
			PreRun: preRun,
			Run:    run,
		}
		checkCmd.AddCommand(cmd)
		checkCmd.ValidArgs = append(checkCmd.ValidArgs, check.Name)
		checkCmd.PreRun = preRun
	}
	rootCmd.AddCommand(checkCmd)
}
