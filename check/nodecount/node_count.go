package nodecount

import (
	"fmt"

	"github.com/adyatlov/bun"
)

func init() {
	check := bun.Check{
		Name: "node-count",
		Description: "Count nodes of each type, checks if the amount of " +
			"master nodes is odd",
		CheckFunc: checkFunc,
	}
	bun.RegisterCheck(check)
}

func checkFunc(c *bun.Check, b bun.Bundle) {
	stats := fmt.Sprintf(
		"Masters: %v, Agents: %v, Public Agents: %v, Total: %v",
		len(b.Masters), len(b.Agents), len(b.PublicAgents), len(b.Hosts))
	if len(b.Masters)%2 != 0 {
		c.Status = bun.SOK
		c.Summary = stats
	} else {
		c.Status = bun.SProblem
		c.Summary = fmt.Sprintf("Number of masters is not valid. %v", stats)
	}
}
