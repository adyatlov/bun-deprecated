package file

import "github.com/adyatlov/dbundle"

func init() {
	f := dbundle.FileType{
		Name:        "dcos-version",
		ContentType: dbundle.JSON,
		Path:        "opt/mesosphere/etc/dcos-version.json",
		Description: "Contains DC/OS version, DC/OS image commit and bootstrap ID",
		HostTypes: map[dbundle.HostType]struct{}{
			dbundle.Master: {}, dbundle.Agent: {}, dbundle.PublicAgent: {},
		},
	}
	dbundle.RegisterFileType(f)
}
