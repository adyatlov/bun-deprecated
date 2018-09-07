package dcosversion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adyatlov/bun"
)

const (
	name        = "dcos-version"
	description = "Verify that all hosts in the cluster have the same DC/OS version installed"
	errDiffVer  = "Versions are different"
)

func init() {
	bun.RegisterCheck(bun.CheckInfo{name, description}, checkVersion)
}

func checkVersion(ctx context.Context, b bun.Bundle,
	p chan<- bun.Progress) (bun.Fact, error) {
	fact := bun.Fact{Status: bun.SOK}
	fact.Errors = make([]string, 0)
	step := 0
	for _, host := range b.Hosts {
		// Check if canceled
		select {
		case <-ctx.Done():
			return fact, ctx.Err()
		default:
		}
		func() {
			// Read version
			file, err := host.OpenFile("dcos-version")
			if err != nil {
				fact.Errors = append(fact.Errors, err.Error())
				return
			}
			defer func() {
				if err := file.Close(); err != nil {
					log.Printf("Error when closing file dcos-version: %v", err)
				}
			}()
			decoder := json.NewDecoder(file)
			verStruct := &struct{ Version string }{}
			if err = decoder.Decode(verStruct); err != nil {
				fact.Errors = append(fact.Errors, err.Error())
				return
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
			step++
		}()
	}
	if fact.Status == bun.SOK && len(fact.Errors) > 0 {
		fact.Status = bun.SUndefined
	}
	return fact, nil
}
