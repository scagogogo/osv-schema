package osv

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// ------------------------------------------------ ---------------------------------------------------------------------

type AffectedSlice[EcosystemSpecific, DatabaseSpecific any] []*Affected[EcosystemSpecific, DatabaseSpecific]

var _ sql.Scanner = &AffectedSlice[any, any]{}
var _ driver.Valuer = &AffectedSlice[any, any]{}

func (x *AffectedSlice[EcosystemSpecific, DatabaseSpecific]) Scan(src any) error {
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

func (x AffectedSlice[EcosystemSpecific, DatabaseSpecific]) Value() (driver.Value, error) {
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

// Affected 此漏洞的影响范围
//Example:
// "affected": [
//
//	{
//	  "package": {
//	    "ecosystem": "RubyGems",
//	    "name": "sprout"
//	  },
//	  "ranges": [
//	    {
//	      "type": "ECOSYSTEM",
//	      "events": [
//	        {
//	          "introduced": "0"
//	        },
//	        {
//	          "last_affected": "0.7.246"
//	        }
//	      ]
//	    }
//	  ]
//	}
//
// ],
type Affected[EcosystemSpecific, DatabaseSpecific any] struct {
	Package  *Package                   `json:"package" yaml:"package" db:"package" bson:"package"`
	Ranges   []*Range[DatabaseSpecific] `json:"ranges" yaml:"ranges" db:"ranges" bson:"ranges"`
	Severity []*Severity                `json:"severity" yaml:"severity" db:"severity" bson:"severity"`
	Versions []string                   `json:"versions" yaml:"versions" db:"versions" bson:"versions"`

	// 由包管理器决定
	EcosystemSpecific EcosystemSpecific `json:"ecosystem_specific" yaml:"ecosystem_specific" db:"ecosystem_specific" bson:"ecosystem_specific"`

	// 由具体实现的数据库决定
	DatabaseSpecific DatabaseSpecific `json:"database_specific" yaml:"database_specific" db:"database_specific" bson:"database_specific"`
}

var _ sql.Scanner = &Affected[any, any]{}
var _ driver.Valuer = &Affected[any, any]{}

func (x *Affected[EcosystemSpecific, DatabaseSpecific]) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Affected[EcosystemSpecific, DatabaseSpecific]) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
