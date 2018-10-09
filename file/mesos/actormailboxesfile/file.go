package actormailboxesfile

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "processes",
		Description: "contains mailbox contents for all actors in the Mesos process on the host.",
		ContentType: bun.CTJson,
		Paths: []string{"5050-__processes__.json",
			"5051-__processes__.json",
			"5050:__processes__.json",
			"5051:__processes__.json"},
		DirTypes: []bun.DirType{bun.DTMaster, bun.DTAgent, bun.DTPublicAgent},
	}
	bun.RegisterFileType(f)
}

// MesosActor represents the structure of the __processess__ file.
type MesosActor struct {
	ID     string `json:"id"`
	Events []struct{}
}
