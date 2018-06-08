package nodecount

import (
	"context"
	"fmt"

	"github.com/adyatlov/bun"
)

const (
	name        = "node-count"
	description = "Count nodes of each type, checks if the amount of master nodes is odd"
)

func init() {
	bun.RegisterCheck(bun.CheckInfo{name, description}, countNodes)
}

func countNodes(ctx context.Context, b bun.Bundle,
	p chan<- bun.Progress) (bun.Fact, error) {
	fact := bun.Fact{}
	nMasters := 0
	nAgents := 0
	nPublic := 0
	step := 0
	for _, h := range b.Hosts {
		select {
		case <-ctx.Done():
			return fact, ctx.Err()
		default:
		}
		step++
		bun.SendProg(p, "Detecting type of "+h.IP, step, len(b.Hosts))
		switch h.Type {
		case bun.Master:
			nMasters++
		case bun.Agent:
			nAgents++
		case bun.PublicAgent:
			nPublic++
		}
		fact.Long += fmt.Sprintf("%v: %v\n", h.IP, h.Type)
	}
	fact.Long += fmt.Sprintf("total: %v", len(b.Hosts))
	if nMasters%2 != 0 {
		fact.Status = bun.SOK
		fact.Short = fmt.Sprintf(
			"Masters: %v, Agents: %v, Public Agents: %v, Total: %v",
			nMasters, nAgents, nPublic, len(b.Hosts))
	} else {
		fact.Status = bun.SProblem
		fact.Short = fmt.Sprintf(
			"Number of masters is not valid: %v, Agents: %v, Public Agents: %v, Total: %v",
			nMasters, nAgents, nPublic, len(b.Hosts))
	}
	return fact, nil
}
