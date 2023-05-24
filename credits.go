package osv_schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type Credits struct {
	Name    string   `json:"name" yaml:"name" db:"name" bson:"name"`
	Contact []string `json:"contact" yaml:"contact" db:"contact" bson:"name"`
	Type    string   `json:"type" yaml:"type" db:"type" bson:"name"`
}

var _ sql.Scanner = &Credits{}
var _ driver.Valuer = &Credits{}

func (x *Credits) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Credits) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
