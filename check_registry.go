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
		panic("bun.RegisterCheck: called twice for check " + c.Name)
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
func GetCheck(name string) (Check, error) {
	checkRegistryMu.RLock()
	check, ok := checkRegistry[name]
	checkRegistryMu.RUnlock()
	if !ok {
		return check, fmt.Errorf("No such check: %v", name)
	}
	return check, nil
}
