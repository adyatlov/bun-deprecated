package dcosversionfile

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "dcos-version",
		Description: "contains DC/OS version, DC/OS image commit and bootstrap ID.",
		ContentType: bun.JSON,
		Paths:       []string{"opt/mesosphere/etc/dcos-version.json"},
		DirTypes:    []bun.DirType{bun.Master, bun.Agent, bun.PublicAgent},
	}
	bun.RegisterFileType(f)
}

// Version represents the dcos-version JSON file
type Version struct {
	Version string
}
