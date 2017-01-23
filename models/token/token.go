package token

import (
	"bitbucket.org/exonch/ch-store/mappers"
	"log"
	"regexp"
	"time"
)

type Token struct {
	token   *string
	user_id *string
	created *time.Time
	expired *time.Time
	active  *bool
}

var mapper mappers.DBMapper

func SetDBMapper(conn *mappers.DBMapper) {
	if conn == nil {
		log.Fatalf("Mapper is nil")
	}
	mapper = *conn
}

func (token *Token) checkFields(fields ...string) (bool, []error) {
	errs := make([]error, 0)
	patternUUID := "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"
	for _, field := range fields {
		switch field {
		case "token":
			if token.token == nil {
				errs = append(errs, fmt.Errorf("Set token field"))
			} else if ok, _ := regexp.Match(patternUUID, []byte(*token.token)); !ok {
				errs = append(errs, fmt.Errorf("Not valid token UUID format"))
			}
		case "user_id":
			if token.user_id == nil {
				errs = append(errs, fmt.Errorf("Set user_id field"))
			} else if ok, _ := regexp.Match(patternUUID, []byte(*token.user_id)); !ok {
				errs = append(errs, fmt.Errorf("Not valid user_id UUID format"))
			}
		}
	}
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}
