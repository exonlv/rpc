package utils

import (
	"database/sql"
	"reflect"
	"fmt"
)

type Row struct {
	*sql.Row
}

//TODO: support other types
//Raw.scan wrapper allows scan *bool, *string, *int64 types
func (row Row) ScanNill(dest ...interface{}) error {
	replaceDest := make([]interface{}, 0)
	for _, val := range dest {
		switch reflect.TypeOf(val) {
		case reflect.PtrTo(reflect.TypeOf(true)):
			replaceDest = append(replaceDest, new(sql.NullBool))
		case reflect.PtrTo(reflect.TypeOf("")):
			replaceDest = append(replaceDest, new(sql.NullString))
		case reflect.PtrTo(reflect.TypeOf(int64(1))):
			replaceDest = append(replaceDest, new(sql.NullInt64))
		default:
			replaceDest = append(replaceDest, val)
		}
	}
	if err := row.Scan(replaceDest...); err != nil {
		return err
	}

	for i, val := range replaceDest {
		fmt.Println(reflect.TypeOf(val))
		switch reflect.TypeOf(val) {
		case reflect.TypeOf(new(sql.NullInt64)):
			t := val.(*sql.NullInt64)
			if t.Valid {
				*(dest[i].(*int64)) = t.Int64
			}
		case reflect.TypeOf(new(sql.NullBool)):
			t := val.(*sql.NullBool)
			if t.Valid {
				*(dest[i].(*bool)) = t.Bool
			}
		case reflect.TypeOf(new(sql.NullString)):
			t := val.(*sql.NullString)
			if t.Valid {
				*(dest[i].(*string)) = t.String
			}
		}
	}

	return nil
}