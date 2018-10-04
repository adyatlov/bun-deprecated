package health

import (
	"fmt"
	"strings"

	"github.com/adyatlov/bun"
	"github.com/adyatlov/bun/file/healthfile"
)

func init() {
	builder := bun.CheckBuilder{
		Name:               "health",
		Description:        "Check if all DC/OS components are healthy",
		ForEachMaster:      check,
		ForEachAgent:       check,
		ForEachPublicAgent: check,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}

func check(host bun.Host) (ok bool, details interface{}, err error) {
	h := healthfile.Host{}
	if err = host.ReadJSON("health", &h); err != nil {
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
