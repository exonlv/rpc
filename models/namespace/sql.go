package namespace

// Add (Namespace, *bool) - Добавление новой записи в таблице Namespace. Передаются только: label, user_id
func (_ *Namespace) Add(ns Namespace, ok *bool) {
	label := ns.Label
	userId := ns.UserID
	_, err := db.Exec("INSERT INTO namespaces(label, user_id) VALUES($1, $2)", label, userId)
	*ok = checkErr(err)
}

// Delete (Namespace, *bool) - изменение removed -> true
func (_ *Namespace) Delete(ns Namespace, ok *bool) {
	_, err := db.Exec("UPDATE namespaces SET removed=TRUE where id=$1", ns.ID)
	*ok = checkErr(err)

}

// GetAll (user_id string, *[]Namespace) - возврат всех Namespace пользователя
func (_ *Namespace) GetAll(userId string, ok *bool) []*Namespace {
	rows, err := db.Query("SELECT * FROM namespaces WHERE user_id = $1", userId)
	*ok = checkErr(err)
	defer rows.Close()
	namespaces := make([]*Namespace, 0)
	for rows.Next() {
		ns := new(Namespace)
		err := rows.Scan(&ns.ID,
			&ns.Label,
			&ns.UserID,
			&ns.Created,
			&ns.Active,
			&ns.Removed,
			&ns.KubeExist)
		*ok = checkErr(err)
		namespaces = append(namespaces, ns)
	}
	return namespaces

}

// Get (id string, *Namespace) - возврат конкретного Namespace пользователя
func (_ *Namespace) Get(id string, ok *bool) *Namespace {
	ns := new(Namespace)
	row := db.QueryRow("SELECT * FROM namespaces WHERE id = $1", id)
	err := row.Scan(&ns.ID,
		&ns.Label,
		&ns.UserID,
		&ns.Created,
		&ns.Active,
		&ns.Removed,
		&ns.KubeExist)
	*ok = checkErr(err)

	return ns

}

// Activate(id string, *bool) - изменение active -> true
func (_ *Namespace) Activate(id string, ok *bool) {
	_, err := db.Exec("UPDATE namespaces SET active=TRUE where id=$1", id)
	*ok = checkErr(err)
}

// Deactivate (id string, *bool) - изменение active -> false
func (_ *Namespace) Deactivate(id string, ok *bool) {
	_, err := db.Exec("UPDATE namespaces SET active=FALSE where id=$1", id)
	*ok = checkErr(err)
}

// CreatedInKube (id string, *bool) - изменение kube_exist -> true
func (_ *Namespace) CreateInKube(id string, ok *bool) {
	_, err := db.Exec("UPDATE namespaces SET kube_exist=TRUE where id=$1", id)
	*ok = checkErr(err)
}

// DeletedInKube (id string, *bool) - изменение kube_exist -> false
func (_ *Namespace) DeletedInKube(id string, ok *bool) {
	_, err := db.Exec("UPDATE namespaces SET kube_exist=FALSE where id=$1", id)
	*ok = checkErr(err)
}

// Rename (Namespace, *bool) - изменение label
func (_ *Namespace) Rename(ns Namespace, ok *bool) {
	_, err := db.Exec("UPDATE namespaces SET label=$1 where id=$2", ns.Label, ns.ID)
	*ok = checkErr(err)
}
