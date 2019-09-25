# THE PROJECT MOVED TO THE MESOSPHERE GITHUB ACCOUNT, please submit issues and pull requests there. Thank you!
[New repository](github.com/mesosphere/bun)
# Bun

DC/OS [diagnostics bundle](https://docs.mesosphere.com/1.11/cli/command-reference/dcos-node/dcos-node-diagnostics-create/) analysis tool

## Installation

[Download binaries from the release page](https://github.com/adyatlov/bun/releases) or install from source:

```bash
$ go get github.com/adyatlov/bun/bun
```

## Usage

```bash
$ mkdir bundle
$ unzip bundle.zip -d bundle
$ cd bundle
bundle$ bun 
[PROBLEM] "dcos-version" - Versions are different.
---------------
Problem details
---------------
master 172.20.0.23 has DC/OS version 1.11.0
master 172.20.0.24 has DC/OS version 1.11.0
agent 172.20.0.27 has DC/OS version 1.11.0
agent 172.20.0.28 has DC/OS version 1.11.0
agent 172.20.0.29 has DC/OS version 1.11.0
agent 172.20.0.21 has DC/OS version 1.10.1
agent 172.20.0.25 has DC/OS version 1.11.0
public agent 172.20.0.26 has DC/OS version 1.11.0

[PROBLEM] "health" - Problems were found.
---------------
Problem details
---------------
agent 172.20.0.21: The following components are not healthy:
dcos-docker-gc.service: health = 1

[OK] "mesos-actor-mailboxes" - All Mesos actors are fine.
[OK] "node-count" - Masters: 3, Agents: 5, Public Agents: 1, Total: 9
```

You can use the `-p` flag if you do not want to change a current directory: 

```bash
bun -p <path-to-bundle-directory>
```

## Adding new checks

Each check uses one or more bundle files. Please, refer to the `filetypes/files_type_yaml.go`
file to find out a name of the file type by the file name and vice versa.

### Simple search check

To add a simple check which fails when a specified string is found in a
specified file of a specified type, add a YAML definition to the YAML 
document in the `checks/search_checks_yaml.go` file:

```yaml
...
- name: disk-space-exhibitor
  description: Check disk space errors in Exhibitor logs
  fileTypeName: exhibitor-log
  searchString: Failed to remove rootfs mount point
```

### Check if a certain condition is fulfilled on each node

If you would like to check if a certain condition is fulfilled on each node of a certain role 
(i.e.: master, agent or public agent), please use the `bun.CheckBuilder` with a default 
aggregate function:

```go
package health

import (
	"fmt"
	"github.com/adyatlov/bun/filetypes"
	"strings"

	"github.com/adyatlov/bun"
)

func init() {
	builder := bun.CheckBuilder{
		Name:                    "diagnostics-health",
		Description:             "Check if all DC/OS components are healthy",
		CollectFromMasters:      collect,
		CollectFromAgents:       collect,
		CollectFromPublicAgents: collect,
		Aggregate:               bun.DefaultAggregate,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}

func collect(host bun.Host) (ok bool, details interface{}, err error) {
	h := filetypes.Host{}
	if err = host.ReadJSON("diagnostics-health", &h); err != nil {
		return
	}
	unhealthy := []string{}
	for _, u := range h.Units {
		if u.Health != 0 {
			unhealthy = append(unhealthy,
				fmt.Sprintf("%v: health = %v", u.ID, u.Health))
		}
	}
	if len(unhealthy) > 0 {
		details = fmt.Sprintf("The following components are not healthy:\n%v",
			strings.Join(unhealthy, "\n"))
		ok = false
	} else {
		ok = true
	}
	return
}
```

### More complex checks

If you need a check which requires analysis of a collected data, you can use a custom
aggregate function:

```go
package dcosversion

import (
	"fmt"
	"github.com/adyatlov/bun"
	"github.com/adyatlov/bun/filetypes"
)

func init() {
	builder := bun.CheckBuilder{
		Name: "dcos-version",
		Description: "Verify that all hosts in the cluster have the " +
			"same DC/OS version installed",
		CollectFromMasters:      collect,
		CollectFromAgents:       collect,
		CollectFromPublicAgents: collect,
		Aggregate:               aggregate,
	}
	check := builder.Build()
	bun.RegisterCheck(check)
}

func collect(host bun.Host) (ok bool, details interface{}, err error) {
	v := filetypes.Version{}
	if err = host.ReadJSON("dcos-version", &v); err != nil {
		return
	}
	details = v.Version
	ok = true
	return
}

func aggregate(c *bun.Check, b bun.CheckBuilder) {
	version := ""
	// Compare versions
	details := []string{}
	ok := true
	for _, r := range b.OKs {
		v := r.Details.(string)
		if version == "" {
			version = v
		}
		if v != version {
			ok = false
		}
		details = append(details, fmt.Sprintf("%v %v has DC/OS version %v",
			r.Host.Type, r.Host.IP, v))
	}
	// No need to interpret problems, as we didn't create it in the host check.
	if ok {
		c.OKs = details
		c.Summary = fmt.Sprintf("All versions are the same: %v.", version)
	} else {
		c.Problems = details
		c.Summary = "Versions are different."
	}
}
```

## Feedback

Please, report bugs and share your ideas for new features via the [issue page](https://github.com/adyatlov/bun/issues).

## Contributing

[Pull requests](https://github.com/adyatlov/bun/pulls) are welcome.
