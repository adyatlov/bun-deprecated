package filetypes

import (
	"fmt"
	"github.com/adyatlov/bun"
	"gopkg.in/yaml.v2"
)

type yamlFile struct {
	Name string `yaml:"name"`
	ContentType string `yaml:"contentType"`
	Paths []string `yaml:"paths"`
	Description string `yaml:"description"`
	DirTypes []string `yaml:"dirTypes"`
}

func init() {
	var files []yamlFile
	err := yaml.Unmarshal([]byte(filesYAML), &files)
	if err !=nil {
		panic(err)
	}
	for _, file := range files {
		fileType, err := convert(file)
		if err != nil {
			panic(err)
		}
		bun.RegisterFileType(fileType)
	}
}

func convert(y yamlFile) (fileType bun.FileType, err error) {
	fileType.Name = y.Name
	fileType.Description = y.Description
	switch y.ContentType {
	case string(bun.CTJson): fileType.ContentType = bun.CTJson
	case "journal": fileType.ContentType = bun.CTJournal
	case "dmesg": fileType.ContentType = bun.CTDmesg
	case "output": fileType.ContentType = bun.CTOutput
	case "other": fileType.ContentType = bun.CTOther
	default:
		err = fmt.Errorf("FileType '%v' has unknown ContentType '%v'", fileType.Name, y.ContentType)
	}
	for _, s := range y.DirTypes {
		var dirType bun.DirType
		dirType, err = convertDirType(s)
		if err != nil {
			return
		}
		fileType.DirTypes = append(fileType.DirTypes, dirType)
	}
	fileType.Paths = y.Paths
	return
}

func convertDirType(s string) (d bun.DirType, err error) {
	switch s {
	case "root":
		d = bun.DTRoot
	case "master":
		d = bun.DTMaster
	case "agent":
		d = bun.DTAgent
	case "public agent":
		d = bun.DTPublicAgent
	default:
		err = fmt.Errorf("unknown DirType: %v", s)
	}
	return
}

