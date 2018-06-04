package file

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "health",
		ContentType: bun.JSON,
		Path:        "dcos-diagnostics-health.json",
		Description: "contains health of systemd services corresponding to DC/OS components.",
		HostTypes: map[bun.HostType]struct{}{
			bun.Master: {}, bun.Agent: {}, bun.PublicAgent: {},
		},
	}
	bun.RegisterFileType(f)
}
