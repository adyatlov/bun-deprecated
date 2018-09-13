package agentlog

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "mesos-agent-log",
		Description: "Mesos agent jounrald log",
		ContentType: bun.Journal,
		Paths: []string{
			"dcos-mesos-slave.service",
			"dcos-mesos-slave-public.service",
		},
		HostTypes: []bun.HostType{bun.Agent, bun.PublicAgent},
	}
	bun.RegisterFileType(f)
}
