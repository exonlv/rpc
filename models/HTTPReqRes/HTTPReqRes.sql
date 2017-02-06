--- SQL HTTPReqRes ---

CREATE TABLE HTTPReqRes(
    Path String,
    Params String,
    Req_data String,
    Created_date Date,
    Created_time String,
    Res_data String,
    Status String,
    Method String

)
ENGINE=MergeTree(Created_date, (Path, Created_date), 8192);
