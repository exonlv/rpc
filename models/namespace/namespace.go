package namespace

import "database/sql"

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
