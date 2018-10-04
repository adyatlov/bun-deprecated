package unmountvolume

import (
	"fmt"

	"github.com/adyatlov/bun"
)

func init() {
	builder := bun.CheckBuilder{
		Name:                    "unmount-volume",
		Description:             "Checks if Mesos agents had problems unmounting local persistent volumes",
		CollectFromAgents:       collect,
		CollectFromPublicAgents: collect,
		Aggregate:               bun.DefaultAggregate,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}

func collect(host bun.Host) (ok bool, details interface{}, err error) {
	line, n, err := host.FindFirstLine("mesos-agent-log", "Failed to destroy nested containers")
	if err != nil {
		return
	}
	if n != 0 {
		details = fmt.Sprintf("%v: %v", n, line)
		return
	}
	ok = true
	return
}
