package volume

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

//Add(Volume, *bool) - добавление нового Volume
func (_ *Volume) Add(volume Volume, ok *bool) error {
	replicasMin := 2
	replicasMax := 5
	if replicasMax < *volume.Replicas || *volume.Replicas < replicasMin {
		*ok = false
		return fmt.Errorf("Replicas should be in range 2 - 5")
	}
	if len(*volume.VolumeServers) != *volume.Replicas {
		*ok = false
		return fmt.Errorf("Replicas must be aqual len(volumeservers)")
	}
	var err error
	var volumeID string
	volumeID, err = newUUID()
	volume.VolumeID = &volumeID
	var errs []error
	*ok, errs = volume.checkFields("volume_id", "labels", "replicas", "volumeservers", "limit", "user_id")
	if !*ok {
		return errs[0]
	}
	query := "INSERT INTO volumes(volume_id, label, replicas, volumeservers, limits, user_id) VALUES($1, $2, $3, $4, $5, $6)"
	err = queryExecutionHandler(query, volume.VolumeID, volume.Label, volume.Replicas, "{"+strings.Join(*volume.VolumeServers, ", ")+"}", volume.Limit, volume.UserID) //TODO: need refactoring
	*ok = checkErr(err)
	return err
}

// GetByUser(user_id, *[]Volume) - поиск всех volumes пользователя
func (_ *Volume) GetByUser(user_id string, volumes *[]Volume) error {
	rows, err := db.Query("SELECT * FROM volumes WHERE user_id = $1", user_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		volume := Volume{}
		var volumeServers []string
		err := rows.Scan(
			&volume.VolumeID,
			&volume.Label,
			&volume.Replicas,
			pq.Array(&volumeServers), //TODO: need refactoring
			&volume.Limit,
			&volume.UserID,
			&volume.Cteated,
			&volume.Active,
			&volume.Exists,
		)
		volume.VolumeServers = &volumeServers
		if err != nil {
			return err
		}

		*volumes = append(*volumes, volume)
	}
	return nil
}

// Get(Volume, *Volume). Возврат Volume по label, user_id
func (_ *Volume) Get(volume Volume, resp *Volume) error {
	_, errs := volume.checkFields("label", "user_id")
	if errs != nil {
		return errs[0]
	}
	vol := Volume{}
	row := db.QueryRow("SELECT * FROM volumes WHERE label = $1 and user_id = $2", volume.Label, volume.UserID)
	var volumeServers []string
	err := row.Scan(
		&vol.VolumeID,
		&vol.Label,
		&vol.Replicas,
		pq.Array(&volumeServers), //TODO: need refactoring
		&vol.Limit,
		&vol.UserID,
		&vol.Cteated,
		&vol.Active,
		&vol.Exists,
	)
	vol.VolumeServers = &volumeServers
	if err != nil {
		return err
	}
	*resp = vol
	return nil
}

//Get(volume_id, *Volume) - получение Volume
func (_ *Volume) GetById(volume_id string, resp *Volume) error {
	volume := Volume{}
	row := db.QueryRow("SELECT * FROM volumes WHERE volume_id = $1", volume_id)
	var volumeServers []string
	err := row.Scan(
		&volume.VolumeID,
		&volume.Label,
		&volume.Replicas,
		pq.Array(&volumeServers), //TODO: need refactoring
		&volume.Limit,
		&volume.UserID,
		&volume.Cteated,
		&volume.Active,
		&volume.Exists,
	)
	volume.VolumeServers = &volumeServers
	if err != nil {
		return err
	}
	*resp = volume
	return nil
}

//Scale(Volume, *bool) - обновление replicas. Действует на replicas, volumeservers
func (_ *Volume) Scale(volume Volume, ok *bool) error {
	replicasMin := 2
	replicasMax := 5
	if replicasMax < *volume.Replicas || *volume.Replicas < replicasMin {
		*ok = false
		return fmt.Errorf("Replicas should be in range 2 - 5")
	}
	if len(*volume.VolumeServers) != *volume.Replicas {
		*ok = false
		return fmt.Errorf("Replicas must be aqual len(volumeservers)")
	}
	query := "UPDATE volumes SET replicas=$1, volumeservers=$2 WHERE volume_id=$3"
	err := queryExecutionHandler(query, volume.Replicas, "{"+strings.Join(*volume.VolumeServers, ", ")+"}", volume.VolumeID)
	*ok = checkErr(err)
	return err
}

//Rename(Volume, *bool) - обновление replicas. Действует на label
func (_ *Volume) Rename(volume Volume, ok *bool) error {
	query := "UPDATE volumes SET label=$1 WHERE volume_id=$2"
	err := queryExecutionHandler(query, volume.Label, volume.VolumeID)
	*ok = checkErr(err)
	return err
}

// Resize(Volume, *bool) - обновление. Действует на limit. Новый лимит всегда больше старого значения, иначе ошибка
func (_ *Volume) Resize(volume Volume, ok *bool) error {
	limitSize := Volume{}
	row := db.QueryRow("SELECT * FROM volumes WHERE volume_id = $1", volume.VolumeID)
	var volumeServers []string
	err := row.Scan(
		&limitSize.VolumeID,
		&limitSize.Label,
		&limitSize.Replicas,
		pq.Array(&volumeServers), //TODO: need refactoring
		&limitSize.Limit,
		&limitSize.UserID,
		&limitSize.Cteated,
		&limitSize.Active,
		&limitSize.Exists,
	)
	volume.VolumeServers = &volumeServers
	if err != nil {
		*ok = false
		return err
	}
	if *limitSize.Limit >= *volume.Limit {
		*ok = false
		return fmt.Errorf("New limit must be greater than the existing limit")
	}
	query := "UPDATE volumes SET limits=$1 WHERE volume_id=$2"
	err = queryExecutionHandler(query, volume.Limit, volume.VolumeID)
	*ok = checkErr(err)
	return err
}

//Activate(volume_id, *bool) - active -> true
func (_ *Volume) Activate(volume_id string, ok *bool) error {
	query := "UPDATE volumes SET active=TRUE where volume_id=$1"
	err := queryExecutionHandler(query, volume_id)
	*ok = checkErr(err)
	return err
}

//Deactivate(volume_id, *bool) - active -> false
func (_ *Volume) Deactivate(volume_id string, ok *bool) error {
	query := "UPDATE volumes SET active=FALSE where volume_id=$1"
	err := queryExecutionHandler(query, volume_id)
	*ok = checkErr(err)
	return err
}

//FindByVolumeserver(volumeserver_id, *[]Volume) - возврат всех Volume упомянутых в массиве volumeservers
func (_ *Volume) FindByVolumeserver(volumeserver_id string, volumes *[]Volume) error {
	rows, err := db.Query("SELECT * FROM volumes WHERE volumeservers @> ARRAY[$1]::uuid[]", volumeserver_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		volume := Volume{}
		var volumeServers []string
		err := rows.Scan(
			&volume.VolumeID,
			&volume.Label,
			&volume.Replicas,
			pq.Array(&volumeServers), //TODO: need refactoring
			&volume.Limit,
			&volume.UserID,
			&volume.Cteated,
			&volume.Active,
			&volume.Exists,
		)
		volume.VolumeServers = &volumeServers
		if err != nil {
			return err
		}
		*volumes = append(*volumes, volume)
	}

	return nil

}

//UsageByVolumeserver(volumeserver_id, integer) - сумма limit всех упомянутых в массиве volumeservers
func (_ *Volume) UsageByVolumeserver(volumeserver_id string, limits *int) error {
	rows, err := db.Query("SELECT * FROM volumes WHERE volumeservers @> ARRAY[$1]::uuid[]", volumeserver_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	*limits = 0
	for rows.Next() {
		volume := Volume{}
		var volumeServers []string
		err := rows.Scan(
			&volume.VolumeID,
			&volume.Label,
			&volume.Replicas,
			pq.Array(&volumeServers), //TODO: need refactoring
			&volume.Limit,
			&volume.UserID,
			&volume.Cteated,
			&volume.Active,
			&volume.Exists,
		)
		volume.VolumeServers = &volumeServers
		if err != nil {
			return err
		}
		*limits += *volume.Limit
	}
	return nil
}

// queryExecutionHandler (query string, args ...interface{}) error
// query - sql запрос типа "UPDATE volume SET label=$1 where id=$2" где $1 и $2 подставляются из args
// args - аргументы встраиваемые в sql запрос(Порядок аргументов важен!)
func queryExecutionHandler(query string, args ...interface{}) error {
	row, err := db.Exec(query, args...)
	err = rowNumbersHandler(row)
	return err
}

// Проверяет колличество обработаных записей, если не было обработано ни одной - возвращает ошибку noRowsProcessedError, иначе nil.
func rowNumbersHandler(row sql.Result) error {
	noRowsProcessedError := errors.New("Failed to update/create the volume. Maybe there is no volume in the database.")
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return noRowsProcessedError
	}
	return nil
}
