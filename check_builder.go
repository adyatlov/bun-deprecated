package bun

import (
	"fmt"
)

const errName = "bun.CheckBuilder.Build: Check Name should not be empty"
const errForEach = "bun.CheckBuilder.Build: At least one of the ForEach functions" +
	"shoul be specified"

// MsgErr is a standard message used in the check summary when errors
// occures during the check.
const MsgErr = "Error(s) occurred while performing the check."

// CheckHost checks an individual host. It returns status, details, and error if
// the function cannot perform a check. If the returned error is not nil, then
// the status is ignored.
type CheckHost func(Host) (bool, interface{}, error)

// Interpret check reults.
type Interpret func(*Check, CheckBuilder)

// Result hols results of the CheckHost function.
type Result struct {
	Host    Host
	OK      bool
	Details interface{}
	Err     error
}

// CheckBuilder helps to create checks.
type CheckBuilder struct {
	Name               string    // Required
	Description        string    // Optional
	ForEachMaster      CheckHost // At least one of
	ForEachAgent       CheckHost // the ForEach... functions
	ForEachPublicAgent CheckHost // are required
	ProblemSummary     string    // Optional
	OKSummary          string    // Optional
	Interpret          Interpret // Implement if the default is not sufficient
	Problems           []Result  // Do not set
	OKs                []Result  // Do not set
}

// Build returns a Check
func (b *CheckBuilder) Build() Check {
	if b.Name == "" {
		panic(errName)
	}
	if b.ForEachMaster == nil && b.ForEachAgent == nil &&
		b.ForEachPublicAgent == nil {
		panic(errForEach)
	}
	if b.ProblemSummary == "" {
		b.ProblemSummary = "Problems were found."
	}
	if b.OKSummary == "" {
		b.OKSummary = "No problems were found."
	}
	if b.Interpret == nil {
		b.Interpret = interpret
	}
	return Check{
		Name:        b.Name,
		Description: b.Description,
		CheckFunc:   b.checkFunc,
	}
}

// BuildAndRegister calls CheckBuilder.Build and register the resulted check.
func (b *CheckBuilder) BuildAndRegister() {
	check := b.Build()
	RegisterCheck(check)
}

func formatMsg(h Host, msg string) string {
	return fmt.Sprintf("%v %v: %v", h.Type, h.IP, msg)
}

func (b *CheckBuilder) checkHosts(c *Check, h map[string]Host, ch CheckHost) {
	for _, host := range h {
		r := Result{}
		r.Host = host
		r.OK, r.Details, r.Err = ch(host)
		if r.Err != nil {
			c.Errors = append(c.Errors, formatMsg(r.Host, r.Err.Error()))
		} else if r.OK {
			b.OKs = append(b.OKs, r)
		} else {
			b.Problems = append(b.Problems, r)
		}
	}
}

// Default inplementation of the Interpret function.
// It assumes that the implementations of the CheckHost function return
// Result Details as a string or nil.
func interpret(c *Check, b CheckBuilder) {
	for _, r := range b.Problems {
		if r.Details != nil {
			c.Problems = append(c.Problems, formatMsg(r.Host, r.Details.(string)))
		}
	}
	for _, r := range b.OKs {
		if r.Details != nil {
			c.OKs = append(c.OKs, formatMsg(r.Host, r.Details.(string)))
		}
	}
}

// Implementation of the Check.CheckFunc
func (b *CheckBuilder) checkFunc(c *Check, bundle Bundle) {
	if b.ForEachMaster != nil {
		b.checkHosts(c, bundle.Masters, b.ForEachMaster)
	}
	if b.ForEachAgent != nil {
		b.checkHosts(c, bundle.Agents, b.ForEachAgent)
	}
	if b.ForEachPublicAgent != nil {
		b.checkHosts(c, bundle.PublicAgents, b.ForEachPublicAgent)
	}
	b.Interpret(c, *b)
	if len(c.Problems) > 0 {
		c.Status = SProblem
		if c.Summary == "" {
			c.Summary = b.ProblemSummary
		}
		if len(c.Errors) > 0 {
			c.Summary += " " + MsgErr
		}
	} else if len(c.Errors) == 0 {
		c.Status = SOK
		if c.Summary == "" {
			c.Summary = b.OKSummary
		}
	} else {
		c.Status = SUndefined
		if c.Summary == "" {
			c.Summary = MsgErr
		}
	}
}
