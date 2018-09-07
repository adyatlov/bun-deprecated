package health

import (
	"fmt"
	"strings"

	"github.com/adyatlov/bun"
	"github.com/adyatlov/bun/file/health"
)

const (
	name        = "health"
	description = "Check if all DC/OS components are healthy"
)

func init() {
	bun.RegisterCheck(
		bun.CheckInfo{
			Name:        name,
			Description: description,
		},
		healthCheck.Check)
}

var healthCheck bun.AtomicCheck = bun.AtomicCheck{
	ForEachMaster:      check,
	ForEachAgent:       check,
	ForEachPublicAgent: check,
}

func check(host bun.Host) (ok bool, msg string, err error) {
	h := health.Host{Units: make([]health.Unit, 0)}
	if err = host.ReadJSON("health", &h); err != nil {
		return
	}
	if h.IP != host.IP {
		err = fmt.Errorf("IP specified in the health JSON file %v != host IP %v",
			h.IP, host.IP)
	}
	unhealthy := make([]string, 0)
	for _, u := range h.Units {
		if u.Health != 0 {
			unhealthy = append(unhealthy,
				fmt.Sprintf("%v: health = %v", u.ID, u.Health))
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
