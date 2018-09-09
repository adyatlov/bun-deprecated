package bun

import (
	"errors"
	"fmt"
)

const errName = "bun.CheckBuilder.Build: Check Name should be specified"
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
	Name               string
	Description        string
	ForEachMaster      CheckHost
	ForEachAgent       CheckHost
	ForEachPublicAgent CheckHost
	ProblemSummary     string
	OKSummary          string
	Interpret          Interpret
	Problems           []Result
	OKs                []Result
}

// Build returns a Check
func (b *CheckBuilder) Build() (Check, error) {
	check := Check{}
	if b.Name == "" {
		return check, errors.New(errName)
	}
	if b.ForEachMaster == nil && b.ForEachAgent == nil &&
		b.ForEachPublicAgent == nil {
		return check, errors.New(errForEach)
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
	b.Problems = []Result{}
	b.OKs = []Result{}
	check = Check{
		Name:        b.Name,
		Description: b.Description,
		CheckFunc:   b.check,
		Problems:    []string{},
		Errors:      []string{},
		OKs:         []string{},
	}
	return check, nil
}

// BuildAndRegister calls CheckBuilder.Build and register the resulted check.
// This function panics when error occures during the build.
func (b *CheckBuilder) BuildAndRegister() {
	check, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("Fatal error occurred while building the \"%v\"check: %v",
			b.Name, err.Error()))
	}
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
// It assumes that the CheckHost returns Result Details as a string.
func interpret(c *Check, b CheckBuilder) {
	for _, r := range b.Problems {
		d := r.Details.(string)
		if d != "" {
			c.Problems = append(c.Problems, formatMsg(r.Host, d))
		}
	}
	for _, r := range b.OKs {
		d := r.Details.(string)
		if d != "" {
			c.OKs = append(c.OKs, formatMsg(r.Host, d))
		}
	}
}

// Check.CheckFunc
func (b *CheckBuilder) check(c *Check, bundle Bundle) {
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
