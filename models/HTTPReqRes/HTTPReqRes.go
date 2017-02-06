package HTTPReqRes

import (
	"fmt"

	clickhouse "github.com/roistat/go-clickhouse"
)

type HTTPReqRes struct {
	Path        *string
	Params      *string
	ReqData     *string
	CreatedDate *string //format "1970-01-01"
	CreatedTime *string //format "20:12:18"
	ResData     *string
	Status      *string
	Method      *string
}

// transport, connection to yandex clickhouse
var conn = clickhouse.NewConn("127.0.0.1:8123", clickhouse.NewHttpTransport())

func (httpreqres *HTTPReqRes) checkField(fields ...string) (bool, []error) {
	errs := make([]error, 0)
	for _, field := range fields {
		switch field {
		case "Path":
			if httpreqres.Path == nil {
				errs = append(errs, fmt.Errorf("Set Path field"))
			}
		case "Params":
			if httpreqres.Params == nil {
				errs = append(errs, fmt.Errorf("Set Params field"))
			}
		case "Req_data":
			if httpreqres.ReqData == nil {
				errs = append(errs, fmt.Errorf("Set Req_Data field"))
			}
		case "Created_date":
			if httpreqres.CreatedDate == nil {
				errs = append(errs, fmt.Errorf("Set Created_date field"))
			}
		case "Created_time":
			if httpreqres.CreatedTime == nil {
				errs = append(errs, fmt.Errorf("Set Created_Time field"))
			}
		case "Res_data":
			if httpreqres.ResData == nil {
				errs = append(errs, fmt.Errorf("Set Res_Data field"))
			}
		case "Status":
			if httpreqres.Status == nil {
				errs = append(errs, fmt.Errorf("Set Status field"))
			}
		case "Method":
			if httpreqres.Method == nil {
				errs = append(errs, fmt.Errorf("Set Method field"))
			}
		}
	}
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}
