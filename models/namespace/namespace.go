package namespace

import (
	"database/sql"
	"errors"
)

type Namespace struct {
	ID        string
	Label     string
	UserID    string
	Created   string
	Active    bool
	Removed   bool
	KubeExist bool
}

var db *sql.DB

func checkErr(err error) bool {
	if err != nil {
		return false
	}
	return true
}

// Проверяет колличество обработаных записей, если не было обработано ни одной - возвращает ошибку noRowsProcessedError, иначе nil.
func rowNumbersHandler(row sql.Result) error {
	noRowsProcessedError := errors.New("Failed to update the namespace. Maybe there is no namespace with such ID in the database.")
	rowsNumber, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsNumber < 1 {
		return noRowsProcessedError
	}
	return err
}
