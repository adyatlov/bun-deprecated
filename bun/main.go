package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adyatlov/bun"
	_ "github.com/adyatlov/bun/check/dcosversion"
	_ "github.com/adyatlov/bun/check/node"
	_ "github.com/adyatlov/bun/file"
)

const printProgress = false

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while detecting a working directory: %v\n", err.Error())
	}
	bundle, err := bun.NewBundle(path)
	if err != nil {
		log.Fatalf("Error while identifying basic bundle parameters: %v\n", err.Error())
	}
	prog := make(chan bun.NamedProgress, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if printProgress {
		go func() {
			for {
				select {
				case p := <-prog:
					log.Printf("%v - %v/%v", p.Stage, p.Step, p.Count)
				case <-ctx.Done():
					log.Println(ctx.Err())
					return
				}
			}
		}()
	}
	report, err := bun.RunCheck(ctx, "dcos-version", bundle, prog)
	if err != nil {
		log.Fatalf("Error while running check %v: %v\n", report.Name, err.Error())
		return
	}
	printReport(report)
	report, err = bun.RunCheck(ctx, "node-count", bundle, prog)
	if err != nil {
		log.Fatalf("Error while running check %v: %v\n", report.Name, err.Error())
		return
	}
	printReport(report)
}

func printReport(r bun.Report) {
	fmt.Printf("%v: %v - %v\n", r.Name, r.Short, r.Status)
	if r.Status == bun.SProblem {
		fmt.Printf("Details:\n%v\n", r.Long)
		if r.Status == bun.SError {
			fmt.Printf("Errors: %v\n", r.Long)
			for i, err := range r.Errors {
				fmt.Printf("Err %v: %v\n", i+1, err)
			}
		}
	}
}
