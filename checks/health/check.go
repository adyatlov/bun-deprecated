package health

import (
	"fmt"
	"github.com/adyatlov/bun/filetypes"
	"strings"

	"github.com/adyatlov/bun"
)

func init() {
	builder := bun.CheckBuilder{
		Name:                    "diagnostics-health",
		Description:             "Check if all DC/OS components are healthy",
		CollectFromMasters:      collect,
		CollectFromAgents:       collect,
		CollectFromPublicAgents: collect,
		Aggregate:               bun.DefaultAggregate,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}

func collect(host bun.Host) (ok bool, details interface{}, err error) {
	h := filetypes.Host{}
	if err = host.ReadJSON("diagnostics-health", &h); err != nil {
		return
	}
	unhealthy := []string{}
	for _, u := range h.Units {
		if u.Health != 0 {
			unhealthy = append(unhealthy,
				fmt.Sprintf("%v: health = %v", u.ID, u.Health))
		}
	}
	if len(unhealthy) > 0 {
		details = fmt.Sprintf("The following components are not healthy:\n%v",
			strings.Join(unhealthy, "\n"))
		ok = false
	} else {
		ok = true
	}
	return
}
