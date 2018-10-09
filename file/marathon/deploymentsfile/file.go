package deployments

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "marathon-deployments",
		Description: "Marathon application deployments",
		ContentType: bun.CTJson,
		Paths: []string{
			"8443-v2_deployments.json",
			"8443:v2_deployments.json",
		},
		DirTypes: []bun.DirType{bun.DTMaster},
	}
	bun.RegisterFileType(f)
}
