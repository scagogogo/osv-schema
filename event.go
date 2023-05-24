package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

//	"events": [
//	 {
//	   "introduced": "2.3.0"
//	 },
//	 {
//	   "fixed": "2.3.18"
//	 }
//
// ]
type Event struct {
	Introduced   string `json:"introduced" yaml:"introduced" db:"introduced" bson:"introduced"`
	Fixed        string `json:"fixed" yaml:"fixed" db:"fixed" bson:"fixed"`
	LastAffected string `json:"last_affected" yaml:"last_affected" db:"last_affected" bson:"last_affected"`
	Limit        string `json:"limit" yaml:"limit" db:"limit" bson:"limit"`
}

var _ sql.Scanner = &Event{}
var _ driver.Valuer = &Event{}

func (x *Event) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Event) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
