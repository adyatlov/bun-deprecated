package bun

import (
	"fmt"
	"sync"
)

var (
	checkRegistry   = make(map[string]Check)
	checkRegistryMu sync.RWMutex
)

// RegisterCheck registers a new check to make it discoverable for consumers.
func RegisterCheck(c Check) {
	checkRegistryMu.Lock()
	defer checkRegistryMu.Unlock()
	if _, exists := checkRegistry[c.Name]; exists {
		panic(fmt.Sprintf("bun.RegisterCheck: called twice for check %v", c.Name))
	}
	checkRegistry[c.Name] = c
}

// Checks Returns all registered checks.
func Checks() []Check {
	checkRegistryMu.RLock()
	defer checkRegistryMu.RUnlock()
	checks := make([]Check, 0, len(checkRegistry))
	for _, c := range checkRegistry {
		checks = append(checks, c)
	}
	return checks
}

// GetCheck returns check by name.
func GetCheck(name string) Check {
	checkRegistryMu.RLock()
	check, ok := checkRegistry[name]
	checkRegistryMu.RUnlock()
	if !ok {
		panic(fmt.Sprintf("bun.GetCheck: don't have check %v", name))
	}
	return check
}
