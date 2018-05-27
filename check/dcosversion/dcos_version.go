package dcosversion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/adyatlov/dbundle"
)

const name = "dcos-version"
const description = "This check verifies that all hosts in the cluster have the same DC/OS version installed."
const errDiffVer = "Versions are different"

func run(ctx context.Context, b dbundle.Bundle,
	p chan<- dbundle.Progress) (*dbundle.Fact, error) {
	fact := dbundle.Fact{OK: true, Name: name}
	percentInc := 100 / len(b.Hosts)
	prog := dbundle.Progress{Name: name}
	for _, host := range b.Hosts {
		// Report progress
		prog.Stage = "Checking version installed on " + host.IP + " ..."
		select {
		case p <- prog:
		default:
		}
		// Check if canceled
		select {
		case <-ctx.Done():
			return nil, errors.New("Canceled")
		default:
		}
		// Read version
		file, err := host.OpenFile("dcos-version")
		if err != nil {
			return nil, err
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
			return nil, err
		}
		// Compare version
		if fact.OK {
			if fact.Short == "" {
				fact.Short = verStruct.Version
			} else if fact.Short != verStruct.Version {
				fact.OK = false
				fact.Short = errDiffVer
			}
		}
		fact.Details += fmt.Sprintf("%v %v has DC/OS version %v\n",
			host.Type, host.IP, verStruct.Version)
		prog.Percent += percentInc
	}
	prog.Stage = "Done!"
	prog.Percent = 100
	select {
	case p <- prog:
	default:
	}
	return &fact, nil
}

func init() {
	c := dbundle.Check{
		Name:        name,
		Description: description,
		Run:         run,
	}
	dbundle.RegisterCheck(c)
}
