package volumeserver

import (
	"database/sql"
	"errors"
)

// Add(VolumeServer, *bool) - добавление
func (_ *VolumeServer) Add(vs VolumeServer, ok *bool) error {
	var errs []error
	*ok, errs = vs.checkFields("ip", "path", "memory")
	if !*ok {
		return errs[0]
	}
	query := "INSERT INTO volumeservers(ip, path, memory) VALUES($1, $2, $3)"
	err := queryExecutionHandler(query, vs.IP, vs.Path, vs.Memory)
	*ok = checkErr(err)
	return err
}

// Activate(volumeserver_id, *bool) - active -> true
func (_ *VolumeServer) Activate(volumeserver_id string, ok *bool) error {
	query := "UPDATE volumeservers SET active=TRUE where volumeserver_id=$1"
	err := queryExecutionHandler(query, volumeserver_id)
	*ok = checkErr(err)
	return err
}

// Deactivate(volumeserver_id, *bool) - active -> false
func (_ *VolumeServer) Deactivate(volumeserver_id string, ok *bool) error {
	query := "UPDATE volumeservers SET active=FALSE where volumeserver_id=$1"
	err := queryExecutionHandler(query, volumeserver_id)
	*ok = checkErr(err)
	return err
}

// ChangeGroup(VolumeServer, *bool) - изменение group
func (_ *VolumeServer) ChangeGroup(vs VolumeServer, ok *bool) error {
	var errs []error
	*ok, errs = vs.checkFields("groups", "volumeserver_id")
	if !*ok {
		return errs[0]
	}
	query := "UPDATE volumeservers SET groups=$1 where volumeserver_id=$2"
	err := queryExecutionHandler(query, vs.Groups, vs.VolumeserverID)
	*ok = checkErr(err)
	return err
}

// ChangeDiskType(VolumeServer, *bool) - изменение disk_type
func (_ *VolumeServer) ChangeDiskType(vs VolumeServer, ok *bool) error {
	var errs []error
	*ok, errs = vs.checkFields("disk_type", "volumeserver_id")
	if !*ok {
		return errs[0]
	}
	query := "UPDATE volumeservers SET disk_type=$1 where volumeserver_id=$2"
	err := queryExecutionHandler(query, vs.DiskType, vs.VolumeserverID)
	*ok = checkErr(err)
	return err
}

// ChangePath(VolumeServer, *bool) - изменение path
func (_ *VolumeServer) ChangePath(vs VolumeServer, ok *bool) error {
	var errs []error
	*ok, errs = vs.checkFields("path", "volumeserver_id")
	if !*ok {
		return errs[0]
	}
	query := "UPDATE volumeservers SET path=$1 where volumeserver_id=$2"
	err := queryExecutionHandler(query, vs.Path, vs.VolumeserverID)
	*ok = checkErr(err)
	return err
}

// GetByGroup(group, *[]VolumeServer) - все VolumeServer определенной группы
func (_ *VolumeServer) GetByGroup(group string, resp *[]VolumeServer) error {
	rows, err := db.Query("SELECT * FROM volumeservers WHERE groups = $1", group)
	if err != nil {
		return err
	}
	defer rows.Close()
	volumeservers := make([]VolumeServer, 0)
	for rows.Next() {
		vs := VolumeServer{}
		err := rows.Scan(&vs.VolumeserverID,
			&vs.IP,
			&vs.Path,
			&vs.Memory,
			&vs.Created,
			&vs.Active,
			&vs.Groups,
			&vs.DiskType,
		)
		if err != nil {
			return err
		}
		// fmt.Println(*vs.VolumeserverID)
		volumeservers = append(volumeservers, vs)
	}
	*resp = volumeservers
	return nil
}

// GetByActivity(bool active, *[]VolumeServer) - все VolumeServer с определнным active
func (_ *VolumeServer) GetByActivity(active bool, resp *[]VolumeServer) error {
	rows, err := db.Query("SELECT * FROM volumeservers WHERE active = $1", active)
	if err != nil {
		return err
	}
	defer rows.Close()
	volumeservers := make([]VolumeServer, 0)
	for rows.Next() {
		vs := VolumeServer{}
		err := rows.Scan(&vs.VolumeserverID,
			&vs.IP,
			&vs.Path,
			&vs.Memory,
			&vs.Created,
			&vs.Active,
			&vs.Groups,
			&vs.DiskType,
		)
		if err != nil {
			return err
		}
		volumeservers = append(volumeservers, vs)
	}
	*resp = volumeservers
	return nil
}

// GetByDiskType(string disk_type, *[]VolumeServer) - все VolumeServer с определенным диском
func (_ *VolumeServer) GetByDiskType(diskType string, resp *[]VolumeServer) error {
	rows, err := db.Query("SELECT * FROM volumeservers WHERE disk_type = $1", diskType)
	if err != nil {
		return err
	}
	defer rows.Close()
	volumeservers := make([]VolumeServer, 0)
	for rows.Next() {
		vs := VolumeServer{}
		err := rows.Scan(&vs.VolumeserverID,
			&vs.IP,
			&vs.Path,
			&vs.Memory,
			&vs.Created,
			&vs.Active,
			&vs.Groups,
			&vs.DiskType,
		)
		if err != nil {
			return err
		}
		volumeservers = append(volumeservers, vs)
	}
	*resp = volumeservers
	return nil
}

// queryExecutionHandler (query string, args ...interface{}) error
// query - sql запрос типа "UPDATE volumeservers SET label=$1 where id=$2" где $1 и $2 подставляются из args
// args - аргументы встраиваемые в sql запрос(Порядок аргументов важен!)
func queryExecutionHandler(query string, args ...interface{}) error {
	row, err := db.Exec(query, args...)
	err = rowNumbersHandler(row)
	return err
}

// Проверяет колличество обработаных записей, если не было обработано ни одной - возвращает ошибку noRowsProcessedError, иначе nil.
func rowNumbersHandler(row sql.Result) error {
	noRowsProcessedError := errors.New("Failed to update/create the volumeserver. Maybe there is no volumeserver in the database.")
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return noRowsProcessedError
	}
	return nil
}
