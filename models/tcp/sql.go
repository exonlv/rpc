package tcp

import (
	"database/sql"
	"errors"
)

var err error

func (tcp *Tcp) insert() error {
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	//записывает в таблице время в формате unix timestamp
	_, err = db.Exec("INSERT INTO tcp(user_id, channel, active, opened, ip) VALUES($1, $2, now(), True, $3)", tcp.user_id, tcp.channel, tcp.ip)
	return err
}

//Обновляет время последней активности и статус, если пользователь с таким каналом и id существует в таблице иначе вызывается метод по добавлению записи
func (_ *Tcp) Open(tcp Tcp, ok *bool) error {
	*ok = false
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	var count string
	err = db.QueryRow("SELECT channel FROM tcp WHERE user_id=$1 AND channel=$2", tcp.user_id, tcp.channel).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		err = tcp.insert()
		if err != nil {
			return err
		}
	case err != nil:
		return err
	default:
		_, err := db.Exec("UPDATE tcp SET (opened, active)=(True, now()) WHERE user_id=$1 AND channel=$2", tcp.user_id, tcp.channel)
		if err != nil {
			return err
		}
	}
	*ok = true
	return nil
}

//Меняет статус активности в False у пользователя на определнном канале
func (_ *Tcp) Close(tcp Tcp, ok *bool) error {
	*ok = false
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	result, err := db.Exec("UPDATE tcp SET opened=False WHERE user_id=$1 and channel=$2", tcp.user_id, tcp.channel)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Nothing to update")
	}
	*ok = true
	return nil
}

//Возвращает все TCP соединения у пользователя
func (_ *Tcp) GetAll(user_id string, resp *[]Tcp) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT id, user_id, channel, active, opened, ip FROM tcp WHERE user_id=$1", user_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	var rTcp []Tcp
	for rows.Next() {
		var tcp Tcp
		if err := rows.Scan(&tcp.id, &tcp.user_id, &tcp.channel, &tcp.active, &tcp.opened, &tcp.ip); err != nil {
			return err
		}
		rTcp = append(rTcp, tcp)
	}
	*resp = rTcp
	return nil
}

func (_ *Tcp) Get(id string, tcp *Tcp) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT id, user_id, channel, active, opened, ip FROM tcp WHERE id=$1", id).Scan(&tcp.id, &tcp.user_id, &tcp.channel, &tcp.active, &tcp.opened, &tcp.ip)
	if err != nil {
		return err
	}
	return nil
}
