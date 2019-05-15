package checks

import (
	"github.com/adyatlov/bun"
	"gopkg.in/yaml.v2"
	"log"
)

func RegisterSearchChecks() {
	var searchChecks []bun.SearchCheckBuilder
	err := yaml.Unmarshal([]byte(searchChecksYAML), &searchChecks)
	if err != nil {
		panic("Cannot read search checks YAML: " + err.Error())
	}
	log.Println(len(searchChecks))
	for _, builder := range searchChecks {
		check := builder.Build()
		bun.RegisterCheck(check)
	}
}
