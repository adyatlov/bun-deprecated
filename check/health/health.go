package health

import (
	"context"
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
	bun.RegisterCheck(bun.CheckInfo{name, description}, checkHealth)
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

func checkHealth(ctx context.Context, b bun.Bundle,
	p chan<- bun.Progress) (bun.Fact, error) {
	fact := bun.Fact{Status: bun.SOK}
	fact.Errors = make([]string, 0)
	he := Health{make([]Host, 0, len(b.Hosts))}
	step := 0
	for _, host := range b.Hosts {
		// Check if canceled
		select {
		case <-ctx.Done():
			return fact, ctx.Err()
		default:
		}
		// For each host
		func() {
			step++
			file, err := host.OpenFile("health")
			if err != nil {
				fact.Errors = append(fact.Errors, err.Error())
				return
			}
			defer func() {
				if err := file.Close(); err != nil {
					log.Printf("Error when closing file health: %v", err)
				}
			}()
			h := Host{host.IP, make([]Unit, 0)}
			defer func() { he.Hosts = append(he.Hosts, h) }()
			decoder := json.NewDecoder(file)
			// For each unit
			if err = decoder.Decode(&h); err != nil {
				fact.Errors = append(fact.Errors,
					fmt.Sprintf("%v: %v", h.IP, err.Error()))
				return
			}
		}()
	}
	var long strings.Builder
	for _, h := range he.Hosts {
		for _, u := range h.Units {
			if u.Health != 0 {
				fact.Status = bun.SProblem
				long.WriteString(
					fmt.Sprintf("%v %v: health = %v\n", h.IP, u.Id, u.Health))
			}
		}
	}
	fact.Long = long.String()
	if fact.Status != bun.SProblem && len(fact.Errors) > 0 {
		fact.Status = bun.SUndefined
	}
	switch fact.Status {
	case bun.SOK:
		fact.Short = "All DC/OS systemd units are healthy."
	case bun.SProblem:
		fact.Short = "Some DC/OS systemd units are not healthy."
	case bun.SUndefined:
		fact.Short = "Errors occurred when checking systemd units health."
	}
	fact.Structured = he
	return fact, nil
}
