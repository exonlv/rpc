package user

import (
	"bitbucket.org/exonch/ch-store/models/utils"
	"encoding/base64"
	"fmt"
	"log"
)

//Create добавляет нового пользователя в БД
func (_ *User) Create(user User, ok *bool) error {
	var errs []error
	*ok, errs = user.checkFields("password", "login", "email", "name")
	if !*ok {
		return errs[0]
	}
	*ok = false
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	user.debugShow()
	salt := genSalt(*user.Login, *user.Email, *user.Name)
	byteKey := getByteKey(*user.Password, salt)
	key := base64.StdEncoding.EncodeToString(byteKey)
	fmt.Printf("salt: %s\n key: %s\n", salt, key) //Debug
	//INSERT
	result, err := db.Exec("INSERT INTO users (pwd_key, salt, login, email, name, last_name) VALUES($1, $2, $3, $4, $5, $6)",
			       key, salt, user.Login, user.Email, user.Name, user.LastName)
	if err != nil {
		return err
	}
	fmt.Println(result)

	*ok = true
	return nil
}

//use UserID or Login and Password fields from User struct
func (_ *User) ChangePassword(user User, ok *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	isPwd, errs := user.checkFields("password")
	if !isPwd {
		return errs[0]
	}
	isUserID, _ := user.checkFields("user_id")
	isLogin, _ := user.checkFields("login")
	var where string
	if isUserID {
		where = fmt.Sprintf("WHERE UserID=%d", *user.UserID)
	} else if isLogin {
		where = fmt.Sprintf("WHERE login='%s'", *user.Login)
	} else {
		return fmt.Errorf("Set UserID or login")
	}
	selectQuery := fmt.Sprintf("SELECT user_id, login, email, name, last_name FROM users %s", where)
	fmt.Println(selectQuery)
	//Check exist
	err = db.QueryRow("SELECT user_id, login, email, name, last_name FROM users "+where).Scan(
		&user.UserID, &user.Login, &user.Email, &user.Name, &user.LastName)
	if err != nil {
		return err
	}
	salt := genSalt(*user.Login, *user.Email, *user.Name)
	key := getKey(*user.Password, salt)
	_, err = db.Exec("UPDATE users SET salt=$1, pwd_key=$2 WHERE user_id=$3", salt, key, user.UserID)
	if err != nil {
		return err
	}
	*ok = true
	return nil
}

func (_ *User) CheckPassword(user User, ok *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	isPwd, errs := user.checkFields("password")
	if !isPwd {
		return errs[0]
	}
	isUserID, _ := user.checkFields("user_id")
	isLogin, _ := user.checkFields("login")
	var where string
	if isUserID {
		where = fmt.Sprintf("WHERE UserID=%d", *user.UserID)
	} else if isLogin {
		where = fmt.Sprintf("WHERE login='%s'", *user.Login)
	} else {
		return fmt.Errorf("Set UserID or login")
	}
	var key, salt string
	err = db.QueryRow(fmt.Sprintf("SELECT pwd_key, salt FROM users %s", where)).Scan(&key, &salt)
	if err != nil {
		return err
	}
	if !checkPassword(*user.Password, salt, key) {
		return fmt.Errorf("Wrong password!")
	}
	*ok = true
	return nil
}

//Replace пытается заменить данные пользователя
func (u *User) Replace(newUser User, resp *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	oldUser := new(User)
	if isLogin, _ := newUser.checkFields("login"); isLogin {
		if err := u.FindByLogin(*newUser.Login, oldUser); err != nil {
			return err
		}
		_, err = db.Exec("UPDATE users SET name=$1, last_name=$2, email=$3, notes=$4 WHERE login=$5",
			&newUser.Name, &newUser.LastName, &newUser.Email, &newUser.Notes, &newUser.Login)
	} else if isUserID, _ := newUser.checkFields("user_id"); isUserID {
		if err := u.FindByID(*newUser.UserID, oldUser); err != nil {
			return err
		}
		_, err = db.Exec("UPDATE users SET name=$1, last_name=$2, email=$3, notes=$4 WHERE user_id=$5",
			&newUser.Name, &newUser.LastName, &newUser.Email, &newUser.Notes, &newUser.Login)
	} else {
		return fmt.Errorf("Set userId or login")
	}
	if err != nil {
		log.Println(err)
		return err
	}
	*resp = true
	return nil
}

//********FIND METHODS***********
//FindByID поиск пользователя по его ID
func (_ *User) FindByID(id string, resp *User) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		log.Println(err)
		return err
	}
	resp.allocateMem()
	row := db.QueryRow("SELECT user_id, login, name, last_name, email, active, notes, register_date FROM users WHERE user_id=$1", id)
	uRow := utils.Row{row}
	err = uRow.ScanNill(resp.UserID, resp.Login, resp.Name, resp.LastName, resp.Email, resp.Active, resp.Notes, resp.Register)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.debugShow()
	return nil
}

//FindByLogin поиск пользователя по его логину
func (u *User) FindByLogin(login string, resp *User) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		log.Println(err)
		return err
	}
	resp.allocateMem()
	row := db.QueryRow("SELECT user_id, login, name, last_name, email, active, notes, register_date FROM users WHERE login=$1", login)
	uRow := utils.Row{row}
	err = uRow.ScanNill(resp.UserID, resp.Login, resp.Name, resp.LastName, resp.Email, resp.Active, resp.Notes, resp.Register)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.debugShow()
	return nil
}

//********ACTIVATE METHODS***********
func (u *User) ActivateByLogin(login string, resp *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		log.Println(err)
		return err
	}
	user := new(User)
	if err := u.FindByLogin(login, user); err != nil {
		return err
	}
	if !*user.Active {
		_, err := db.Exec("UPDATE users SET active=$1 WHERE login=$2", true, login)
		if err != nil {
			return err
		}
	}
	*resp = true
	return nil
}

func (u *User) ActivateByID(userID string, resp *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		log.Println(err)
		return err
	}
	user := new(User)
	if err := u.FindByID(userID, user); err != nil {
		return err
	}
	if !*user.Active {
		_, err := db.Exec("UPDATE users SET active=$1 WHERE user_id=$2", true, userID)
		if err != nil {
			return err
		}
	}
	*resp = true
	return nil
}

func (u *User) DeactivateByLogin(login string, resp *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		log.Println(err)
		return err
	}
	user := new(User)
	if err := u.FindByLogin(login, user); err != nil {
		return err
	}
	if *user.Active {
		_, err := db.Exec("UPDATE users SET active=$1 WHERE login=$2", false, login)
		if err != nil {
			return err
		}
	}
	*resp = true
	return nil
}

func (u *User) DeactivateByID(userID string, resp *bool) error {
	db, err := mapper.GetDB("default")
	if err != nil {
		log.Println(err)
		return err
	}
	user := new(User)
	if err := u.FindByID(userID, user); err != nil {
		return err
	}
	if *user.Active {
		_, err := db.Exec("UPDATE users SET active=$1 WHERE user_id=$2", false, userID)
		if err != nil {
			return err
		}
	}
	*resp = true
	return nil
}
