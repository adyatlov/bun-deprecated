package deployments

import (
	"fmt"

	"github.com/adyatlov/bun"
)

// max number of Marathon deployments considered healthy
const maxDeployments = 10

func init() {
	builder := bun.CheckBuilder{
		Name:          "marathon-deployments",
		Description:   "Check for too many running Marathon app deployments",
		ForEachMaster: check,
	}
	builder.BuildAndRegister()
}

func check(host bun.Host) (ok bool, details interface{}, err error) {
	deployments := []struct{}{}
	if err = host.ReadJSON("marathon-deployments", &deployments); err != nil {
		return
	}
	if len(deployments) > maxDeployments {
		details = fmt.Sprintf("Too many deployments: %v", len(deployments))
		return
	}
	ok = true
	return
}
