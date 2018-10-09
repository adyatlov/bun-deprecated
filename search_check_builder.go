package bun

import (
	"fmt"
)

// SearchCheckBuilder builds a check which searches for the specified
// string in the the specified files. If the pattern
// is found, the check is considered problematic.
// The number of the found line and its content appear in the Check.Problems of the check.
// The check searches only for the first appearance of the line.
type SearchCheckBuilder struct {
	Name         string // Required
	Description  string // Optional
	FileTypeName string // Required
	SearchString string // Required
}

// Build creates a bun.Check.
func (b SearchCheckBuilder) Build() Check {
	if b.FileTypeName == "" {
		panic("FileTypeName should be specified.")
	}
	if b.SearchString == "" {
		panic("SearchString should be set.")
	}
	builder := CheckBuilder{
		Name:        b.Name,
		Description: b.Description,
		Aggregate:   DefaultAggregate,
	}
	t := GetFileType(b.FileTypeName)
	for _, dirType := range t.DirTypes {
		switch dirType {
		case DTMaster:
			builder.CollectFromMasters = b.collect
		case DTAgent:
			builder.CollectFromAgents = b.collect
		case DTPublicAgent:
			builder.CollectFromPublicAgents = b.collect
		}
	}
	return builder.Build()
}

func (b SearchCheckBuilder) collect(host Host) (ok bool, details interface{},
	err error) {
	n, line, err := host.FindLine(b.FileTypeName, b.SearchString)
	if err != nil {
		return
	}
	if n != 0 {
		details = fmt.Sprintf("%v: %v", n, line)
		return
	}
	ok = true
	return
}
