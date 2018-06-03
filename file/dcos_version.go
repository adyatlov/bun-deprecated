package file

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "dcos-version",
		ContentType: bun.JSON,
		Path:        "opt/mesosphere/etc/dcos-version.json",
		Description: "Contains DC/OS version, DC/OS image commit and bootstrap ID",
		HostTypes: map[bun.HostType]struct{}{
			bun.Master: {}, bun.Agent: {}, bun.PublicAgent: {},
		},
	}
	bun.RegisterFileType(f)
}
