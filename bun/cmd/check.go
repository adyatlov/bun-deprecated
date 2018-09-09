package cmd

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/adyatlov/bun"
	"github.com/spf13/cobra"
)

var printLong = false
var bundle *bun.Bundle

func printReport(c bun.Check) {
	printEmptyLine := false
	fmt.Printf("[%v] \"%v\" - %v\n", c.Status, c.Name, c.Summary)
	if printLong {
		if len(c.Problems) > 0 {
			fmt.Println("--------")
			fmt.Println("Problems")
			fmt.Println("--------")
			fmt.Println(strings.Join(c.Problems, "\n"))
			printEmptyLine = true
		}
		if len(c.Errors) > 0 {
			fmt.Println("------")
			fmt.Println("Errors")
			fmt.Println("------")
			fmt.Println(strings.Join(c.Errors, "\n"))
			printEmptyLine = true
		}
		if len(c.OKs) > 0 {
			fmt.Println("---")
			fmt.Println("OKs")
			fmt.Println("---")
			fmt.Println(strings.Join(c.OKs, "\n"))
			printEmptyLine = true
		}
	} else {
		if len(c.Problems) > 0 {
			fmt.Println("--------")
			fmt.Println("Problems")
			fmt.Println("--------")
			fmt.Println(strings.Join(c.Problems, "\n"))
			printEmptyLine = true
		}
		if len(c.Errors) > 0 {
			fmt.Println("------")
			fmt.Println("Errors")
			fmt.Println("------")
			fmt.Println(strings.Join(c.Errors, "\n"))
			printEmptyLine = true
		}
	}
	if printEmptyLine {
		fmt.Print("\n")
	}
}

func preRun(cmd *cobra.Command, args []string) {
	if bundle != nil {
		return
	}
	b, err := bun.NewBundle(context.Background(), bundlePath)
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

func init() {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check DC/OS diagnostics bundle for possible problems",
		Long: `Check DC/OS diagnostics bundle for possible problems.

Specify a subcommand to run a specific check, e.g.` + " `bun check health`." +
			`
Or run all the available checks by not specifying any, i.e.` + " `bun check`.",
		PreRun: preRun,
		Run:    runCheck,
	}
	checkCmd.PersistentFlags().BoolVarP(&printLong, "long", "l", false, "print details")

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
		checkCmd.AddCommand(cmd)
		checkCmd.ValidArgs = append(checkCmd.ValidArgs, check.Name)
		checkCmd.PreRun = preRun
	}
	rootCmd.AddCommand(checkCmd)
}
