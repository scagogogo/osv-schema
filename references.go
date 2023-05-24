package osv

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// References
// Example:
//    {
//      "type": "WEB",
//      "url": "https://github.com/tensorflow/tensorflow/security/advisories/GHSA-vxv8-r8q2-63xw"
//    }
//
type References struct {
	Type string `json:"type" yaml:"type" db:"type" bson:"type"`
	URL  string `json:"url" yaml:"url" db:"url" bson:"url"`
}

var _ sql.Scanner = &References{}
var _ driver.Valuer = &References{}

func (x *References) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *References) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
