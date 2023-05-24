package osv_schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

//	"package": {
//	  "ecosystem": "RubyGems",
//	  "name": "sprout"
//	},
type Package struct {
	Ecosystem string `json:"ecosystem" yaml:"ecosystem" db:"ecosystem" bson:"ecosystem"`
	Name      string `json:"name" yaml:"name" db:"name" bson:"name"`
	PUrl      string `json:"purl" yaml:"purl" db:"purl" bson:"purl"`
}

var _ sql.Scanner = &Package{}
var _ driver.Valuer = &Package{}

func (x *Package) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Package) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
