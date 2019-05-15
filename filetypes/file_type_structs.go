package filetypes

// Version represents the dcos-version JSON file
type Version struct {
	Version string
}

// Health represents the health JSON file
type Health struct {
	Hosts []Host
}

// Host represents the "host" object in the health JSON file
type Host struct {
	Units []Unit
}

// Unit represents the "unit" object in the health JSON file
type Unit struct {
	ID     string `json:"id"`
	Health int
}

// MesosActor represents the structure of the __processess__ file.
type MesosActor struct {
	ID     string `json:"id"`
	Events []struct{}
}

// MsgFailedToUnmouint message appears in the Mesos agent logs when agent cannot
// unmount local persisten colume.
const MsgFailedToUnmouint = "Failed to remove rootfs mount point"
