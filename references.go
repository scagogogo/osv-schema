package osv

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// ------------------------------------------------ ---------------------------------------------------------------------

type References []*Reference

var _ sql.Scanner = &References{}
var _ driver.Valuer = &References{}

func (x *References) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("scan error")
	}
	if len(bytes) == 0 {
		return nil
	}
	return json.Unmarshal(bytes, &x)
}

func (x References) Value() (driver.Value, error) {
	if len(x) == 0 {
		return nil, nil
	}
	marshal, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	if len(marshal) == 0 {
		return nil, nil
	}
	return string(marshal), nil
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Reference
// Example:
//    {
//      "type": "WEB",
//      "url": "https://github.com/tensorflow/tensorflow/security/advisories/GHSA-vxv8-r8q2-63xw"
//    }
//
type Reference struct {
	Type string `json:"type" yaml:"type" db:"type" bson:"type"`
	URL  string `json:"url" yaml:"url" db:"url" bson:"url"`
}

var _ sql.Scanner = &Reference{}
var _ driver.Valuer = &Reference{}

func (x *Reference) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Reference) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
