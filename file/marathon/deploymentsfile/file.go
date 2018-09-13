package deployments

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "marathon-deployments",
		Description: "Marathon application deployments",
		ContentType: bun.JSON,
		Paths: []string{
			"8443-v2_deployments.json",
			"8443:v2_deployments.json",
		},
		DirTypes: []bun.DirType{bun.Master},
	}
	bun.RegisterFileType(f)
}
