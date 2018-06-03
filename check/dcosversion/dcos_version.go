package dcosversion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adyatlov/bun"
)

const name = "dcos-version"
const description = "This check verifies that all hosts in the cluster have the same DC/OS version installed."
const errDiffVer = "Versions are different"

func check(ctx context.Context, b bun.Bundle, p chan<- bun.Progress) (bun.Fact, error) {
	fact := bun.Fact{Status: bun.SOK}
	step := 0
	for _, host := range b.Hosts {
		step++
		// Check if canceled
		select {
		case <-ctx.Done():
			return fact, ctx.Err()
		default:
		}
		// Read version
		file, err := host.OpenFile("dcos-version")
		if err != nil {
			if fact.Errors == nil {
				fact.Errors = make([]string, 0)
			}
			fact.Errors = append(fact.Errors, err.Error())
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("Error occurred when closing file %v: %v",
					file.Name(), err.Error())
			}
		}()
		decoder := json.NewDecoder(file)
		verStruct := &struct{ Version string }{}
		if err = decoder.Decode(verStruct); err != nil {
			return fact, err
		}
		// Compare version
		if fact.Status == bun.SOK {
			if fact.Short == "" {
				fact.Short = verStruct.Version
			} else if fact.Short != verStruct.Version {
				fact.Status = bun.SProblem
				fact.Short = errDiffVer
			}
		}
		fact.Long += fmt.Sprintf("%v %v has DC/OS version %v\n",
			host.Type, host.IP, verStruct.Version)
		// Report progress
		bun.SendProg(p, "Checked version installed on "+host.IP, step, len(b.Hosts))
	}
	if fact.Status == bun.SOK && len(fact.Errors) > 0 {
		fact.Status = bun.SError
	}
	return fact, nil
}

func init() {
	c := bun.CheckInfo{
		Name:        name,
		Description: description,
	}
	bun.RegisterCheck(c, check)
}
