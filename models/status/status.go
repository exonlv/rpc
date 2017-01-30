package status

import (
	"gopkg.in/vmihailenco/msgpack.v2"
	"reflect"
	"time"
)

type Status struct {
	Command_id *string
	User_id    *string
	Created    *time.Time
	Status_id  *int8
	Initiator  *string
}

const (
	STATUS_MAXIMUM = ^uint32(0)
)

func init() {
	msgpack.Register(reflect.TypeOf(Status{}), encodeStatus, decodeStatus)
}

func (status *Status) allocateMomory() {
	status.Command_id = new(string)
	status.User_id = new(string)
	status.Created = new(time.Time)
	status.Status_id = new(int64)
	*status.Status_id = 0
	status.Initiator = new(string)
}

//encode structure Status to slice of interfaces
func encodeStatus(enc *msgpack.Encoder, val reflect.Value) error {
	status := val.Interface().(Status)
	if err := enc.EncodeSliceLen(5); err != nil {
		return err
	}
	if err := enc.EncodeString(*status.Command_id); err != nil {
		return err
	}
	if err := enc.EncodeString(*status.User_id); err != nil {
		return err
	}
	//Time generating in tarantool this just filler
	if err := enc.EncodeInt64(0); err != nil {
		return err
	}
	if err := enc.EncodeInt64(*status.Status_id); err != nil {
		return err
	}
	if err := enc.EncodeString(*status.Initiator); err != nil {
		return err
	}
	return nil
}

//decode slice of interfaces to structure Status
func decodeStatus(dec *msgpack.Decoder, val reflect.Value) error {
	var err error
	var length int
	var microseconds int64
	status := val.Addr().Interface().(*Status)
	status.allocateMomory()
	if length, err = dec.DecodeSliceLen(); err != nil {
		return err
	}
	if length != 5 {
		return fmt.Errorf("array len doesn't match: %d", length)
	}
	if *status.Command_id, err = dec.DecodeString(); err != nil {
		return err
	}
	if *status.User_id, err = dec.DecodeString(); err != nil {
		return err
	}
	if microseconds, err = dec.DecodeInt64(); err != nil {
		return err
	}
	*status.Created = time.Unix(0, microseconds*1000)
	if *status.Status_id, err = dec.DecodeInt64(); err != nil {
		return err
	}
	if *status.Initiator, err = dec.DecodeString(); err != nil {
		return err
	}
	return nil
}

// Check field on nil and blank value
func (status *Status) checkFields(fields ...string) (bool, []error) {
	errs := make([]error, 0)
	for _, field := range fields {
		switch field {
		case "command_id":
			if status.Command_id == nil || *status.Command_id == "" {
				errs = append(errs, fmt.Errorf("Set command_id field"))
			}
		case "user_id":
			if status.User_id == nil || *status.User_id == "" {
				errs = append(errs, fmt.Errorf("Set user_id field"))
			}
		case "initiator":
			if status.Initiator == nil || *status.Initiator == "" {
				errs = append(errs, fmt.Errorf("Set initiator field"))
			}
		}
	}
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}
