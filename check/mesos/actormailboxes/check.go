package actormailboxes

import (
	"fmt"
	"strings"

	"github.com/adyatlov/bun"
	"github.com/adyatlov/bun/file/mesos/actormailboxesfile"
)

// number of events in an actor's mailbox after which the actor is
// considered backlogged
const maxEvents = 30

func init() {
	builder := bun.CheckBuilder{
		Name: "mesos-actor-mailboxes",
		Description: "Check if actor mailboxes in the Mesos process " +
			"have a reasonable amount of messages",
		OKSummary:          "All Mesos actors are fine.",
		ProblemSummary:     "Some Mesos actors are backlogged.",
		ForEachMaster:      check,
		ForEachAgent:       check,
		ForEachPublicAgent: check,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}

func check(host bun.Host) (ok bool, details interface{}, err error) {
	actors := []actormailboxesfile.MesosActor{}
	if err = host.ReadJSON("processes", &actors); err != nil {
		return
	}
	u := []string{}
	for _, a := range actors {
		if len(a.Events) > maxEvents {
			u = append(u, fmt.Sprintf("(Mesos) %v@%v: mailbox size = %v (> %v)",
				a.ID, host.IP, len(a.Events), maxEvents))
		}
	}
	if len(u) > 0 {
		details = "The following Mesos actor mailboxes are too big:\n" +
			strings.Join(u, "\n")
		return
	}
	ok = true
	return
}
