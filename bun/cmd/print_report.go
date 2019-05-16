package cmd

import (
	"fmt"
	"strings"

	"github.com/adyatlov/bun"
)

func printReport(c bun.Check) {
	printEmptyLine := false
	fmt.Printf("[%v] \"%v\" - %v\n", c.Status, c.Name, c.Summary)
	if verbose {
		if len(c.Problems) > 0 {
			fmt.Println("---------------")
			fmt.Println("Problem details")
			fmt.Println("---------------")
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
			fmt.Println("-------")
			fmt.Println("Details")
			fmt.Println("-------")
			fmt.Println(strings.Join(c.OKs, "\n"))
			printEmptyLine = true
		}
	} else {
		if len(c.Problems) > 0 {
			fmt.Println("---------------")
			fmt.Println("Problem details")
			fmt.Println("---------------")
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
