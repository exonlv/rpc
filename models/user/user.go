package user

import (
	"bitbucket.org/exonch/ch-store/mappers"
	"fmt"
	"log"
	"time"
)

type (
	//User описывает модель пользователя системы
	User struct {
		UserID   *string       `json:"user_id"`
		Login    *string    `json:"login"`
		Password *string    `json:"-"`
		Key      *string    `json:"key"` // encrypted password
		Name     *string    `json:"name"`
		LastName *string    `json:"last_name"`
		Email    *string    `json:"email"`
		Notes    *string    `json:"notes"`
		Active   *bool      `json:"active"`
		Register *time.Time `json:"register"`
	}
)

func (u *User) allocateMem() {
	u.UserID = new(string)
	u.Login = new(string)
	u.Password = new(string)
	u.Key = new(string)
	u.Name = new(string)
	u.LastName = new(string)
	u.Email = new(string)
	u.Notes = new(string)
	u.Active = new(bool)
	u.Register = new(time.Time)
}

var mapper mappers.DBMapper

func SetDBMapper(conn *mappers.DBMapper) {
	if conn == nil {
		log.Fatalf("Mapper is nil")
	}
	mapper = *conn
}

func (user *User) checkFields(fields ...string) (bool, []error) {
	errs := make([]error, 0)
	for _, field := range fields {
		switch field {
		case "password":
			if user.Password == nil {
				errs = append(errs, fmt.Errorf("Set password field"))
			}
		case "login":
			if user.Login == nil {
				errs = append(errs, fmt.Errorf("Set login field"))
			}
		case "email":
			if user.Email == nil {
				errs = append(errs, fmt.Errorf("Set email field"))
			}
		case "name":
			if user.Name == nil {
				errs = append(errs, fmt.Errorf("Set name field"))
			}
		case "last_name":
			if user.LastName == nil {
				errs = append(errs, fmt.Errorf("Set last name field"))
			}
		case "user_id":
			if user.UserID == nil {
				errs = append(errs, fmt.Errorf("Set userID field"))
			}
		}
	}
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}

func (user *User)debugShow() {
	printif := func(name string, value *string) {
		if value != nil {
			fmt.Printf("%s: %s\n", name, *value)
		} else {
			fmt.Printf("%s: <nil>\n", name)
		}
	}
	printif("Name", user.Name)
	printif("Email", user.Email)
	printif("Login", user.Login)
	printif("Pswd", user.Password)
	printif("LastName", user.LastName)
	fmt.Printf("Active %v", user.Active)
	fmt.Printf("Reg time %v", user.Register)
	printif("UserId", user.UserID)
}
