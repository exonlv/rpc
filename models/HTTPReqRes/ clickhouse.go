package HTTPReqRes

import clickhouse "github.com/roistat/go-clickhouse"

//Add (HTTPReqRes, *bool) - Добавление новой записи в таблице HTTPReqRes.
func (_ *HTTPReqRes) Add(httpreqres HTTPReqRes, ok *bool) error {
	var errs []error
	*ok, errs = httpreqres.checkField("Path", "Params", "Req_data", "Created_date", "Created_time", "Res_data", "Status", "Method")
	if !*ok {
		return errs[0]
	}
	query, err := clickhouse.BuildInsert("HTTPReqRes",
		clickhouse.Columns{"Path", "Params", "Req_data", "Created_date", "Created_time", "Res_data", "Status", "Method"},
		clickhouse.Row{*httpreqres.Path, *httpreqres.Params, *httpreqres.ReqData, *httpreqres.CreatedDate, *httpreqres.CreatedTime, *httpreqres.ResData, *httpreqres.Status, *httpreqres.Method})
	if err != nil {
		*ok = false
		return err
	}
	err = query.Exec(conn)
	if err != nil {
		*ok = false
		return err
	}
	*ok = true
	return nil
}
