package cmd

import (
	_ "github.com/adyatlov/bun/check/dcosversion"
	_ "github.com/adyatlov/bun/check/health"
	_ "github.com/adyatlov/bun/check/marathon/deployments"
	_ "github.com/adyatlov/bun/check/mesos/actormailboxes"
	_ "github.com/adyatlov/bun/check/mesos/unmountvolume"
	_ "github.com/adyatlov/bun/check/nodecount"
	_ "github.com/adyatlov/bun/file/dcosversionfile"
	_ "github.com/adyatlov/bun/file/healthfile"
	_ "github.com/adyatlov/bun/file/marathon/deploymentsfile"
	_ "github.com/adyatlov/bun/file/mesos/actormailboxesfile"
	_ "github.com/adyatlov/bun/file/mesos/agentlogfile"
)
