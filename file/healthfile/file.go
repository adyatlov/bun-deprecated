package healthfile

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "health",
		ContentType: bun.JSON,
		Paths:       []string{"dcos-diagnostics-health.json", "3dt-health.json"},
		Description: "contains health of systemd services corresponding to DC/OS components.",
		HostTypes: map[bun.HostType]struct{}{
			bun.Master: {}, bun.Agent: {}, bun.PublicAgent: {},
		},
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
