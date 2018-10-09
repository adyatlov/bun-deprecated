package healthfile

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "health",
		Description: "contains health of systemd services corresponding to DC/OS components.",
		ContentType: bun.CTJson,
		Paths:       []string{"dcos-diagnostics-health.json", "3dt-health.json"},
		DirTypes:    []bun.DirType{bun.DTMaster, bun.DTAgent, bun.DTPublicAgent},
	}
	bun.RegisterFileType(f)
}

// Health represents the health JSON file
type Health struct {
	Hosts []Host
}

// Host represents the "host" object in the health JSON file
type Host struct {
	Units []Unit
}

// Unit represents the "unit" object in the health JSON file
type Unit struct {
	ID     string `json:"id"`
	Health int
}
