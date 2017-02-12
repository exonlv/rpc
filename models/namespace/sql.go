package namespace

import (
	"database/sql"
	"errors"
)

// Add (Namespace, *bool) - Добавление новой записи в таблице Namespace. Передаются только: label, user_id
func (_ *Namespace) Add(nameserver Namespace, id *string) error {
	query := "INSERT INTO namespaces(label, user_id) VALUES($1, $2)"
	err := queryExecutionHandler(query, nameserver.Label, nameserver.UserID)
	ns := Namespace{}
	row := db.QueryRow("SELECT * FROM namespaces WHERE label = $1", nameserver.Label)
	err = row.Scan(
		&ns.ID,
		&ns.Label,
		&ns.UserID,
		&ns.Created,
		&ns.Active,
		&ns.Removed,
		&ns.KubeExist,
	)
	if err != nil {
		return err
	}

	*id = ns.ID
	return nil
}

// Delete (Namespace, *bool) - изменение removed -> true
func (_ *Namespace) Delete(ns Namespace, ok *bool) error {
	query := "UPDATE namespaces SET removed=TRUE where id=$1"
	err := queryExecutionHandler(query, ns.ID)
	*ok = checkErr(err)
	return err
}

// GetAll (user_id string, *[]Namespace) - возврат всех Namespace пользователя
func (_ *Namespace) GetAll(userId string, resp *[]Namespace) error {
	rows, err := db.Query("SELECT * FROM namespaces WHERE user_id = $1", userId)
	if err != nil {
		return err
	}
	defer rows.Close()
	namespaces := make([]Namespace, 0)
	for rows.Next() {
		ns := Namespace{}
		err := rows.Scan(&ns.ID,
			&ns.Label,
			&ns.UserID,
			&ns.Created,
			&ns.Active,
			&ns.Removed,
			&ns.KubeExist,
		)
		if err != nil {
			return err
		}
		namespaces = append(namespaces, ns)
	}
	*resp = namespaces
	return nil

}

// Get (id string, *Namespace) - возврат конкретного Namespace пользователя
func (_ *Namespace) Get(id string, resp *Namespace) error {
	ns := Namespace{}
	row := db.QueryRow("SELECT * FROM namespaces WHERE id = $1", id)
	err := row.Scan(&ns.ID,
		&ns.Label,
		&ns.UserID,
		&ns.Created,
		&ns.Active,
		&ns.Removed,
		&ns.KubeExist,
	)
	if err != nil {
		return err
	}

	*resp = ns
	return nil

}

// Activate(id string, *bool) - изменение active -> true
func (_ *Namespace) Activate(id string, ok *bool) error {
	query := "UPDATE namespaces SET active=TRUE where id=$1"
	err := queryExecutionHandler(query, id)
	*ok = checkErr(err)
	return err
}

// Deactivate (id string, *bool) - изменение active -> false
func (_ *Namespace) Deactivate(id string, ok *bool) error {
	query := "UPDATE namespaces SET active=FALSE where id=$1"
	err := queryExecutionHandler(query, id)
	*ok = checkErr(err)
	return err
}

// CreatedInKube (id string, *bool) - изменение kube_exist -> true
func (_ *Namespace) CreateInKube(id string, ok *bool) error {
	query := "UPDATE namespaces SET kube_exist=TRUE where id=$1"
	err := queryExecutionHandler(query, id)
	*ok = checkErr(err)
	return err
}

// DeletedInKube (id string, *bool) - изменение kube_exist -> false
func (_ *Namespace) DeletedInKube(id string, ok *bool) error {
	query := "UPDATE namespaces SET kube_exist=FALSE where id=$1"
	err := queryExecutionHandler(query, id)
	*ok = checkErr(err)
	return err
}

// Rename (Namespace, *bool) - изменение label
func (_ *Namespace) Rename(ns Namespace, ok *bool) error {
	query := "UPDATE namespaces SET label=$1 where id=$2"
	err := queryExecutionHandler(query, ns.Label, ns.ID)
	*ok = checkErr(err)
	return err
}

// queryExecutionHandler (query string, args ...interface{}) error
// query - sql запрос типа "UPDATE namespaces SET label=$1 where id=$2" где $1 и $2 подставляются из args
// args - аргументы встраиваемые в sql запрос(Порядок аргументов важен!)
func queryExecutionHandler(query string, args ...interface{}) error {
	row, err := db.Exec(query, args...)
	err = rowNumbersHandler(row)
	return err
}

// Проверяет колличество обработаных записей, если не было обработано ни одной - возвращает ошибку noRowsProcessedError, иначе nil.
func rowNumbersHandler(row sql.Result) error {
	noRowsProcessedError := errors.New("Failed to update/create the namespace. Maybe there is no namespace with such ID in the database.")
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return noRowsProcessedError
	}
	return nil
}
