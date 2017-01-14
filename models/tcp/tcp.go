package tcp

import (
	"bitbucket.org/exonch/ch-store/mappers"
	"log"
	"time"
)

type Tcp struct {
	id      string
	user_id string
	channel string
	active  time.Time
	opened  bool
	ip      string
}

var mapper mappers.DBMapper

func SetDBMapper(conn *mappers.DBMapper) {
	if conn == nil {
		log.Fatalf("Mapper is nil")
	}
	mapper = *conn
}
