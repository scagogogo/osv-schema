package osv

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

//	{
//	  "type": "ECOSYSTEM",
//	  "events": [
//	    {
//	      "introduced": "0"
//	    },
//	    {
//	      "last_affected": "0.7.246"
//	    }
//	  ]
//	}
type Range[DatabaseSpecific any] struct {
	Type   string   `json:"type" yaml:"type" db:"type" bson:"type"`
	Repo   string   `json:"repo" yaml:"repo" db:"repo" bson:"repo"`
	Events []*Event `json:"events" yaml:"events" db:"events" bson:"events"`

	// 由具体实现的数据库决定
	DatabaseSpecific DatabaseSpecific `json:"database_specific" yaml:"database_specific" db:"database_specific" bson:"database_specific"`
}

var _ sql.Scanner = &Range[any]{}
var _ driver.Valuer = &Range[any]{}

func (x *Range[DatabaseSpecific]) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Range[DatabaseSpecific]) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
