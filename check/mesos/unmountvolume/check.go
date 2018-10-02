package unmountvolume

import (
	"github.com/adyatlov/bun"
	"github.com/adyatlov/bun/file/mesos/agentlogfile"
)

func init() {
	builder := bun.SearchCheckBuilder{
		Name: "unmount-volume",
		Description: "Checks if Mesos agents had problems unmounting " +
			"local persistent volumes. MESOS-8830",
		FileTypeName: "mesos-agent-log",
		SearchString: agentlogfile.MsgFailedToUnmouint,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}
