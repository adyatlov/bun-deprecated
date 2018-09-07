package health

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/adyatlov/bun"
)

const (
	name        = "health"
	description = "Check if all DC/OS components are healthy"
)

func init() {
	bun.RegisterCheck(bun.CheckInfo{name, description}, healthCheck.Check)
}

type Health struct {
	Hosts []Host
}
type Host struct {
	IP    string
	Units []Unit
}
type Unit struct {
	Id     string
	Name   string
	Health int
	Output string
}

var healthCheck bun.AtomicCheck = bun.AtomicCheck{
	ForEachMaster:      check,
	ForEachAgent:       check,
	ForEachPublicAgent: check,
}

func check(host bun.Host) (ok bool, msg string, err error) {
	file, err := host.OpenFile("health")
	if err != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error when closing file health: %v", err)
		}
	}()
	h := Host{host.IP, make([]Unit, 0)}
	decoder := json.NewDecoder(file)
	// For each unit
	if err = decoder.Decode(&h); err != nil {
		return
	}
	unhealthy := make([]string, 0)
	for _, u := range h.Units {
		if u.Health != 0 {
			unhealthy = append(unhealthy,
				fmt.Sprintf("%v: health = %v", u.Id, u.Health))
		}
	}
	if len(unhealthy) > 0 {
		msg = "The following components are not healthy:\n" + strings.Join(unhealthy, "\n")
		ok = false
	} else {
		msg = "All the checked components are healthy."
		ok = true
	}
	return
}
