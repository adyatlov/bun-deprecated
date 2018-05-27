package dbundle

import (
	"context"
	"errors"
	"sync"
)

type Progress struct {
	Name    string
	Stage   string
	Percent int
}

type Check struct {
	Name        string
	Description string
	Run         func(context.Context, Bundle, chan<- Progress) (*Fact, error)
}

// Fact is a piece of knowledge about DC/OS cluster learned from a
// diagnostic bundle.
type Fact struct {
	Name       string // the same as Check
	OK         bool
	Short      string
	Details    string
	Structured interface{}
}

// TODO: delete? move to tests?
func Run(check Check, b Bundle) (*Fact, error) {
	var p chan<- Progress
	return check.Run(context.Background(), b, p)
}

var (
	checks   = make(map[string]Check)
	checksMu sync.RWMutex
)

func RegisterCheck(c Check) {
	checksMu.Lock()
	defer checksMu.Unlock()
	if _, dup := checks[c.Name]; dup {
		panic("dcosbundle.RegisterCheck: called twice for driver " + c.Name)
	}
	checks[c.Name] = c
}

func GetCheck(name string) (Check, error) {
	checksMu.RLock()
	defer checksMu.RUnlock()
	check, ok := checks[name]
	if !ok {
		return check, errors.New("No such check: " + name)
	}
	return check, nil
}
