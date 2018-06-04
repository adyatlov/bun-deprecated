package bun

import (
	"context"
	"errors"
	"sync"
)

type Progress struct {
	Stage string
	Step  int
	Count int
}

type Check func(context.Context, Bundle, chan<- Progress) (Fact, error)

type CheckInfo struct {
	Name        string
	Description string
}

type Status string

const (
	SUndefined Status = "UNDEFINED"
	SOK               = "OK"
	SProblem          = "PROBLEM"
)

// Fact is a piece of knowledge about DC/OS cluster learned from a
// diagnostic bundle.
type Fact struct {
	Status     Status
	Short      string
	Long       string
	Errors     []string
	Structured interface{}
}

type Report struct {
	CheckInfo
	Fact
}

type checkRecord struct {
	CheckInfo
	Check
}

var (
	checks   = make(map[string]checkRecord)
	checksMu sync.RWMutex
)

func RegisterCheck(ci CheckInfo, c Check) {
	checksMu.Lock()
	defer checksMu.Unlock()
	if _, dup := checks[ci.Name]; dup {
		panic("dcosbundle.RegisterCheck: called twice for driver " + ci.Name)
	}
	checks[ci.Name] = checkRecord{ci, c}
}

func Checks() []CheckInfo {
	checksMu.RLock()
	defer checksMu.RUnlock()
	cc := make([]CheckInfo, 0, len(checks))
	for _, cr := range checks {
		cc = append(cc, cr.CheckInfo)
	}
	return cc
}

type NamedProgress struct {
	Name string
	Progress
}

func RunCheck(ctx context.Context, name string, b Bundle,
	np chan<- NamedProgress) (Report, error) {
	checksMu.RLock()
	cr, ok := checks[name]
	checksMu.RUnlock()
	if !ok {
		return Report{}, errors.New("No such check: " + name)
	}
	// Handle progress
	p := make(chan Progress, 100)
	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		for {
			select {
			case prog := <-p:
				np <- NamedProgress{cr.Name, prog}
			case <-localCtx.Done():
				return
			}
		}
	}()
	// Running check may take some time
	f, err := cr.Check(ctx, b, p)
	if err != nil {
		return Report{cr.CheckInfo, f}, err
	}
	return Report{cr.CheckInfo, f}, nil
}

func SendProg(p chan<- Progress, stage string, step int, count int) {
	select {
	case p <- Progress{stage, step, count}:
	default:
	}
}
