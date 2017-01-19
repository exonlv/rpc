package volumeserver

import (
	"database/sql"
	"fmt"
)

type VolumeServer struct {
	VolumeserverID *string
	IP             *string
	Path           *string
	Memory         *int
	Created        *string
	Active         *bool
	Groups         *string
	DiskType       *string
}

var db *sql.DB

func (vs *VolumeServer) checkFields(fields ...string) (bool, []error) {
	errs := make([]error, 0)
	for _, field := range fields {
		switch field {
		case "volumeserver_id":
			if vs.VolumeserverID == nil {
				errs = append(errs, fmt.Errorf("Set volumeserver_id field"))
			}
		case "ip":
			if vs.IP == nil {
				errs = append(errs, fmt.Errorf("Set ip field"))
			}
		case "path":
			if vs.Path == nil {
				errs = append(errs, fmt.Errorf("Set path field"))
			}
		case "memory":
			if vs.Memory == nil {
				errs = append(errs, fmt.Errorf("Set memory field"))
			}
		case "disk_type":
			if vs.DiskType == nil {
				errs = append(errs, fmt.Errorf("Set disk_type field"))
			}
		}
	}
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}

func checkErr(err error) bool {
	if err != nil {
		return false
	}
	return true
}
