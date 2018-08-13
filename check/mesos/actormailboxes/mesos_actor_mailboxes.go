package actormailboxes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/adyatlov/bun"
)

const (
	name		= "mesos-actor-mailboxes"
	description	= "Check if actor mailboxes in the Mesos process have sane amount of messages"
	max_events	= 30 // the number of events in an actor's mailbox after which the actor is considered backlogged
)

func init() {
	bun.RegisterCheck(bun.CheckInfo{name, description}, checkMesosActorMailboxes)
}

// Truncated JSON schema of __processes__.
type MesosActor struct {
	Id	string
	Events	[]MailboxEvent
}
type MailboxEvent struct {
}

func checkMesosActorMailboxes(ctx context.Context, b bun.Bundle,
	p chan<- bun.Progress) (bun.Fact, error) {

	// Prepare the output struct.
	fact := bun.Fact{Status: bun.SOK}
	fact.Errors = make([]string, 0)

	// Used to report progress, which is not quite necessary for this
	// lightweight check.
	step := 0

	// Accumulator for error strings.
	var unhealthy strings.Builder

	// Iterate over all hosts in the bundle
	for _, host := range b.Hosts {
		// Check if canceled.
		select {
		case <-ctx.Done():
			return fact, ctx.Err()
		default:
		}

		// For each host.
		func() {
			step++

			file, err := host.OpenFile("processes")
			if err != nil {
				fact.Errors = append(fact.Errors, err.Error())
				return
			}
			defer func() {
				if err := file.Close(); err != nil {
					log.Printf("Error when closing file __processes__: %v", err)
				}
			}()

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				fact.Errors = append(fact.Errors, err.Error())
				return
			}

			var o []MesosActor
			err = json.Unmarshal(bytes, &o)
			if err != nil {
				fact.Errors = append(fact.Errors, err.Error())
				return
			}

			// Walk through all actors and check the number of events in the mailbox.
			for _, a := range o {
				if len(a.Events) > max_events {
					fact.Status = bun.SProblem
					unhealthy.WriteString(
						fmt.Sprintf("(Mesos) %v@%v: mailbox size = %v (> %v)\n", a.Id, host.IP, len(a.Events), max_events))
			    }
			}
		}()
	}


	if unhealthy.Len() > 0 {
		fact.Long = "The following Mesos actor mailboxes are too big:\n" + unhealthy.String()
	} else {
		fact.Long = "All checked Mesos actor mailboxes are fine."
	}
	if fact.Status != bun.SProblem && len(fact.Errors) > 0 {
		fact.Status = bun.SUndefined
	}
	switch fact.Status {
	case bun.SOK:
		fact.Short = "All Mesos actors are fine."
	case bun.SProblem:
		fact.Short = "Some Mesos actors are backlogged."
	case bun.SUndefined:
		fact.Short = "Errors occurred when checking Mesos actor mailboxes."
	}

	return fact, nil
}
