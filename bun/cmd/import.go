package cmd

import (
	_ "github.com/adyatlov/bun/check/dcosversion"
	_ "github.com/adyatlov/bun/check/health"
	_ "github.com/adyatlov/bun/check/mesos/actormailboxes"
	_ "github.com/adyatlov/bun/check/nodecount"
	_ "github.com/adyatlov/bun/file/dcosversion"
	_ "github.com/adyatlov/bun/file/health"
	_ "github.com/adyatlov/bun/file/mesos/actormailboxes"
)
