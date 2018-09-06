package bun

import (
	"context"
	"fmt"
	"strings"
)

// HostCheck checks an individual host. It returns status, details, and error if
// the function cannot perform a check. If the returned error is not nil, then
// the status is ignored.
type HostCheck func(Host) (bool, string, error)

// HostCheckResult is a reult of a single host check.
type HostCheckResult struct {
	IP      string
	OK      bool
	Details string
	Err     error
}

// AtomicCheck builds Check as a sum of HostChecks.
type AtomicCheck struct {
	ForEachMaster      HostCheck
	ForEachAgent       HostCheck
	ForEachPublicAgent HostCheck
	ProblemMessage     string
	OKMessage          string
}

// Check is an impementation of the bun.Check function.
func (a *AtomicCheck) Check(ctx context.Context,
	b Bundle,
	p chan<- Progress) (fact Fact, err error) {
	progress := Progress{}
	okResults := make([]HostCheckResult, 0)
	problemResults := make([]HostCheckResult, 0)
	errResults := make([]HostCheckResult, 0)
	if a.ProblemMessage == "" {
		a.ProblemMessage = "Problems were found."
	}
	if a.OKMessage == "" {
		a.OKMessage = "No problems were found."
	}
	check := func(hosts map[string]Host, hc HostCheck) error {
		for ip, host := range hosts {
			progress.Stage = ip
			select {
			case p <- progress:
			default:
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			r := HostCheckResult{}
			r.IP = host.IP
			r.OK, r.Details, r.Err = hc(host)
			if r.Err != nil {
				errResults = append(errResults, r)
			} else if r.OK {
				okResults = append(okResults, r)
			} else {
				problemResults = append(problemResults, r)
			}
			progress.Step++
		}
		return nil
	}
	queue := make([]func() error, 0, 3)
	if a.ForEachMaster != nil {
		progress.Count += len(b.Masters)
		queue = append(queue, func() error { return check(b.Masters, a.ForEachMaster) })
	}
	if a.ForEachAgent != nil {
		progress.Count += len(b.Agents)
		queue = append(queue, func() error { return check(b.Agents, a.ForEachAgent) })
	}
	if a.ForEachPublicAgent != nil {
		progress.Count += len(b.PublicAgents)
		queue = append(queue, func() error { return check(b.PublicAgents, a.ForEachPublicAgent) })
	}
	for _, c := range queue {
		if err = c(); err != nil {
			return
		}
	}
	short := make([]string, 0)
	if len(errResults) > 0 {
		short = append(short, "Error(s) occured when performing the check.")
	}
	if len(problemResults) > 0 {
		fact.Status = SProblem
		short = append(short, a.ProblemMessage)
	} else if len(errResults) == 0 {
		fact.Status = SOK
		short = append(short, a.OKMessage)
	} else {
		fact.Status = SUndefined
	}
	fact.Short = strings.Join(short, " ")
	var long strings.Builder
	if len(problemResults) > 0 {
		long.WriteString("Problems:\n")
		for _, res := range problemResults {
			long.WriteString(fmt.Sprintf("%v: %v\n", res.IP, res.Details))
		}
	}
	if len(okResults) > 0 {
		long.WriteString("Successfull checks:\n")
		for _, res := range okResults {
			long.WriteString(fmt.Sprintf("%v: %v\n", res.IP, res.Details))
		}
	}
	if len(errResults) > 0 {
		for _, res := range errResults {
			fact.Errors = append(fact.Errors, res.Err.Error())
		}
	}
	fact.Long = long.String()
	return
}
