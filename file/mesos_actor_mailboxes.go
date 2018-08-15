package file

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "processes",
		ContentType: bun.JSON,
		Paths:       []string{"5050-__processes__.json", "5051-__processes__.json", "5050:__processes__.json", "5051:__processes__.json"},
		Description: "contains mailbox contents for all actors in the Mesos process on the host.",
		HostTypes: map[bun.HostType]struct{}{
			bun.Master: {}, bun.Agent: {}, bun.PublicAgent: {},
		},
	}
	bun.RegisterFileType(f)
}
