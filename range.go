package osv_schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// ------------------------------------------------- --------------------------------------------------------------------

type RangeType string

const (

	// RangeTypeSemver The versions introduced and fixed are semantic versions as defined by SemVer 2.0.0, with no leading “v” prefix.
	// The relation u < v denotes the precedence order defined in section 11 of SemVer 2.0. Ranges listed with type SEMVER
	// should not overlap: since SEMVER is a strict linear ordering, it is always possible to simplify to non-overlapping ranges.
	// Specifying one or more SEMVER ranges removes the requirement to specify an explicit enumerated versions list (see the discussion above).
	// Some ecosystems may recommend using SemVer 2.0 for versioning without explicitly enforcing it. In those cases you should use the ECOSYSTEM type instead.
	RangeTypeSemver = "SEMVER"

	// RangeTypeEcosystem The versions introduced and fixed are arbitrary, uninterpreted strings specific to the package ecosystem,
	// which does not conform to SemVer 2.0’s version ordering.
	// It is recommended that you provide an explicitly enumerated versions list when specifying one or more ECOSYSTEM ranges,
	// because ECOSYSTEM range inclusion queries may not be able to be answered without reference to the package ecosystem’s
	// own logic and therefore may not be able to be used by ecosystem-independent processors. The infrastructure and tooling
	// provided by https://osv.dev also provides automation for auto-populating the versions list based on supported ECOSYSTEM
	// ranges as part of the ingestion process.
	RangeTypeEcosystem = "ECOSYSTEM"

	// RangeTypeGit The versions introduced and fixed are full-length Git commit hashes. The repository’s commit graph is needed to evaluate
	// whether a given version is in the range. The relation u < v is true when commit u is a (perhaps distant) parent of commit v.
	RangeTypeGit = "GIT"
)

// ------------------------------------------------- --------------------------------------------------------------------

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
	Type   RangeType `json:"type" yaml:"type" db:"type" bson:"type" gorm:"column:type"`
	Repo   string    `json:"repo" yaml:"repo" db:"repo" bson:"repo" gorm:"column:repo"`
	Events Events    `json:"events" yaml:"events" db:"events" bson:"events" gorm:"column:events;serializer:json"`

	// 由具体实现的数据库决定
	DatabaseSpecific DatabaseSpecific `json:"database_specific" yaml:"database_specific" db:"database_specific" bson:"database_specific" gorm:"column:database_specific;serializer:json"`
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

// ------------------------------------------------- --------------------------------------------------------------------
