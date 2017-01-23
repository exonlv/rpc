package token

import (
	"database/sql"
	"errors"
)

//Add new token in database
func (_ *Token) Add(newToken Token, ok *bool) error {
	*ok = false
	isUser, errs := newToken.checkFields("user_id")
	if !isUser {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO token(user_id) VALUES($1)", *newToken.user_id)
	if err != nil {
		return err
	}
	*ok = true
	return nil
}

//Get token by user_id
func (_ *Token) GetByUser(token Token, respToken *Token) error {
	isUser, errs := token.checkFields("user_id")
	if !isUser {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT token, user_id, created, expired, active FROM token WHERE user_id=$1 AND active=True AND expired>=now()", *token.user_id).Scan(&respToken.token, &respToken.user_id, &respToken.created, &respToken.expired, &respToken.active)
	if err != nil {
		return err
	}
	return nil
}

//Activate token
func (_ *Token) Activate(token Token, ok *bool) error {
	*ok = false
	isToken, errs := token.checkFields("token")
	if !isToken {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	result, err := db.Exec("UPDATE token SET active=True WHERE token=$1", *token.token)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Nothing activate. Token not exist.")
	}
	*ok = true
	return nil
}

//Deactivate token
func (_ *Token) Deactivate(token Token, ok *bool) error {
	*ok = false
	isToken, errs := token.checkFields("token")
	if !isToken {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	result, err := db.Exec("UPDATE token SET active=False WHERE token=$1", *token.token)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Nothing deactivate. Token not exist.")
	}
	*ok = true
	return nil
}

//get all tokens by user_id
func (_ *Token) GetAll(token Token, respTokens *[]Token) error {
	isUser, errs := token.checkFields("user_id")
	if !isUser {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT token, user_id, created, expired, active FROM token WHERE user_id=$1", *token.user_id)
	if err != nil {
		return err
	}
	defer rows.Close()
	var rTokens []Token
	for rows.Next() {
		var token Token
		if err := rows.Scan(&token.token, &token.user_id, &token.created, &token.expired, &token.active); err != nil {
			return err
		}
		rTokens = append(rTokens, token)
	}
	*respTokens = rTokens
	return nil
}

//update expired by token
func (_ *Token) Extend(token Token, ok *bool) error {
	*ok = false
	isToken, errs := token.checkFields("token")
	if !isToken {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	result, err := db.Exec("UPDATE token SET expired=expired+interval '1 day' * 90 WHERE token=$1", *token.token)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Token not exist.")
	}
	*ok = true
	return nil
}

//проверить токен, что он активен и не вышло время пользователя возвращается id пользователя
func (_ *Token) Check(token Token, user_id *string) error {
	var userUUID string
	isToken, errs := token.checkFields("token")
	if !isToken {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT user_id FROM token WHERE token=$1 AND active=True AND expired>=now()", *token.token).Scan(&userUUID)
	if err != nil {
		return err
	}
	*user_id = userUUID
	return nil
}

//get token by id
func (_ *Token) Get(token Token, respToken *Token) error {
	isToken, errs := token.checkFields("token")
	if !isToken {
		return errs[0]
	}
	db, err := mapper.GetDB("default")
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT token, user_id, created, expired, active FROM token WHERE token=$1", *token.token).Scan(&respToken.token, &respToken.user_id, &respToken.created, &respToken.expired, &respToken.active)
	if err != nil {
		return err
	}
	return nil
}
