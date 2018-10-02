package bun

import (
	"fmt"
)

// SearchCheckBuilder builds a checks which searches for the specified pattern
// set by a string in the lines of the specified files. If the pattern
// is found, the check is considered problematic.
// The found lines appear in the Check.Problems of the check. By default, the
// check searches only for the first appearance. Set the FindAll to true if you
// would like to collect all the lines.
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
	}
	t := GetFileType(b.FileTypeName)
	for _, dirType := range t.DirTypes {
		switch dirType {
		case Master:
			builder.ForEachMaster = b.check
		case Agent:
			builder.ForEachAgent = b.check
		case PublicAgent:
			builder.ForEachPublicAgent = b.check
		}
	}
	return builder.Build()
}

func (b SearchCheckBuilder) check(host Host) (ok bool, details interface{},
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
