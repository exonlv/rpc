package volume

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
)

type Volume struct {
	VolumeID      *string
	Label         *string
	Replicas      *int
	VolumeServers *[]string `pg:",array"`
	Limit         *int
	UserID        *string
	Cteated       *string
	Active        *bool
	Exists        *bool
}

var db *sql.DB

func SetDB(sdb *sql.DB) {
	db = sdb
}

func (v *Volume) checkFields(fields ...string) (bool, []error) {
	errs := make([]error, 0)
	for _, field := range fields {
		switch field {
		case "volume_id":
			if v.VolumeID == nil {
				errs = append(errs, fmt.Errorf("Set volume_id field"))
			}
		case "label":
			if v.Label == nil {
				errs = append(errs, fmt.Errorf("Set label field"))
			}
		case "replicas":
			if v.Replicas == nil {
				errs = append(errs, fmt.Errorf("Set replicas field"))
			}
		case "volumeservers":
			if v.VolumeServers == nil {
				errs = append(errs, fmt.Errorf("Set volumeservers field"))
			}
		case "limits":
			if v.Limit == nil {
				errs = append(errs, fmt.Errorf("Set limit field"))
			}
		case "user_id":
			if v.UserID == nil {
				errs = append(errs, fmt.Errorf("Set user_id field"))
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

// newUUID Генерирует uuid
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
