package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adyatlov/dbundle"
	_ "github.com/adyatlov/dbundle/check/dcosversion"
	_ "github.com/adyatlov/dbundle/file"
)

const printProgress = false

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while detecting a working directory: %v\n", err.Error())
	}
	bundle, err := dbundle.NewBundle(path)
	if err != nil {
		log.Fatalf("Error while identifying basic bundle parameters: %v\n", err.Error())
	}
	checkName := "dcos-version"
	check, err := dbundle.GetCheck(checkName)
	if err != nil {
		log.Fatalf("Error when getting check %v: %v\n", checkName, err.Error())
	}
	progCh := make(chan dbundle.Progress, 100)
	done := make(chan struct{})
	if printProgress {
		go func() {
			for {
				select {
				case p := <-progCh:
					log.Printf("%v - %v%%", p.Stage, p.Percent)
				case <-done:
					return
				}
			}
		}()
	}
	factCh := make(chan *dbundle.Fact)
	go func() {
		fact, err := check.Run(context.Background(), bundle, progCh)
		if err != nil {
			log.Printf("Error while running check %v: %v\n", check.Name, err.Error())
		}
		factCh <- fact
	}()
	fact := <-factCh
	close(done)
	if fact != nil {
		printFact(fact)
	}
}

func printFact(f *dbundle.Fact) {
	status := "OK"
	if !f.OK {
		status = "ERROR!"
	}
	fmt.Printf("%v: %v - %v\n", f.Name, f.Short, status)
	if !f.OK {
		fmt.Printf("Details:\n%v", f.Details)
	}
}
