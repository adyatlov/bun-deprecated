package actormailboxesfile

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "processes",
		Description: "contains mailbox contents for all actors in the Mesos process on the host.",
		ContentType: bun.JSON,
		Paths: []string{"5050-__processes__.json",
			"5051-__processes__.json",
			"5050:__processes__.json",
			"5051:__processes__.json"},
		HostTypes: []bun.HostType{bun.Master, bun.Agent, bun.PublicAgent},
	}
	bun.RegisterFileType(f)
}

// MesosActor represents the structure of the __processess__ file.
type MesosActor struct {
	ID     string `json:"id"`
	Events []struct{}
}
