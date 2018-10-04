package dcosversion

import (
	"fmt"

	"github.com/adyatlov/bun"
	"github.com/adyatlov/bun/file/dcosversionfile"
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
	v := dcosversionfile.Version{}
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
			continue
		}
		if v != version {
			ok = false
		}
		details = append(details, fmt.Sprintf("%v %v has DC/OS version %v",
			r.Host.Type, r.Host.IP, v))
	}
	// No need to interpret problems, as we didn't create it in the host check.
	if ok {
		c.Summary = fmt.Sprintf("All versions are the same: %v.", version)
	} else {
		c.Problems = details
		c.Summary = "Versions are different."
	}
}
