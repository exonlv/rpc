package status

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
)

var (
	client *tarantool.Connection
	err    error
	errs   []error
)

func (_ *Status) Update(updateStatus Status, ok *bool) error {
	*ok = false
	if client == nil {
		return fmt.Errorf("Connection nil")
	}
	*ok, errs = updateStatus.checkFields("command_id", "user_id", "initiator")
	if !*ok {
		return errs[0]
	}
	*ok = false
	_, err := client.Call("update", []interface{}{*updateStatus.Command_id, *updateStatus.User_id, *updateStatus.Created, *updateStatus.Status_id, *updateStatus.Initiator})
	if err != nil {
		return err
	}
	*ok = true
	return nil
}

func (_ *Status) Get(getStatus Status, respStatus *Status) error {
	if client == nil {
		return fmt.Errorf("Connection nil")
	}
	ok, errs := getStatus.checkFields("command_id")
	if !ok {
		return errs[0]
	}
	var rStat []Status
	err := client.SelectTyped("status", "command_id", 0, 1, tarantool.IterEq, []interface{}{*getStatus.Command_id}, &rStat)
	if err != nil {
		return err
	}
	if len(rStat) != 1 {
		return fmt.Errorf("Get method: nothing find with command_id=%s", *getStatus.Command_id)
	}
	*respStatus = rStat[0]
	return nil
}

func (_ *Status) GetAll(getStatus Status, respStatus *[]Status) error {
	if client == nil {
		return fmt.Errorf("Connection nil")
	}
	ok, errs := getStatus.checkFields("command_id")
	if !ok {
		return errs[0]
	}
	var rStat []Status
	err = client.SelectTyped("status", "user_id", 0, STATUS_MAXIMUM, tarantool.IterEq, []interface{}{*getStatus.User_id}, &rStat)
	if err != nil {
		return err
	}
	if len(rStat) == 0 {
		return fmt.Errorf("GetAll method: nothing find with user_id=%s", *getStatus.User_id)
	}
	*respStatus = rStat
	return nil
}
