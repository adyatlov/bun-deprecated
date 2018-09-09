package bun

// Status defines possible check outcomes.
type Status string

const (
	// SUndefined means that the check wasn't performed successfully.
	SUndefined Status = "UNDEFINED"
	// SOK means that the bundle passed the check.
	SOK = "OK"
	// SProblem means that the bundle failed to pass the check.
	SProblem = "PROBLEM"
)

// cluster analyzing its diagnostics bundle. Check supposed to populate fields
// of the CheckResult.
// Each check should implement this interface.

// Check cheks some aspect of the DC/OS cluster analyzing its diagnostics
// bundle.
// Checks can be registered in the check registry with the egisterCheck function.
type Check struct {
	Name        string
	Description string
	Status      Status
	CheckFunc   func(*Check, Bundle)
	Summary     string
	Problems    []string
	Errors      []string
	OKs         []string
}

func (c *Check) Run(b Bundle) {
	c.CheckFunc(c, b)
}
