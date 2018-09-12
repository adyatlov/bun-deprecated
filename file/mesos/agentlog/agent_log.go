package agentlog

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "mesos-agent-log",
		ContentType: bun.Journal,
		Paths: []string{
			"dcos-mesos-slave.service",
			"dcos-mesos-slave-public.service",
		},
		Description: "Mesos agent jounrald log",
		HostTypes:   map[bun.HostType]struct{}{bun.Agent: {}, bun.PublicAgent: {}},
	}
	bun.RegisterFileType(f)
}
