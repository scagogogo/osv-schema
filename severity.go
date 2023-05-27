package osv_schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// ------------------------------------------------ ---------------------------------------------------------------------

type SeveritySlice []*Severity

var _ sql.Scanner = &SeveritySlice{}
var _ driver.Valuer = &SeveritySlice{}

func (x SeveritySlice) GetCVSS3() *Severity {
	for _, s := range x {
		if s.Type == SeverityTypeCVSS3 {
			return s
		}
	}
	return nil
}

func (x SeveritySlice) GetCVSS2() *Severity {
	for _, s := range x {
		if s.Type == SeverityTypeCVSS2 {
			return s
		}
	}
	return nil
}

func (x *SeveritySlice) Scan(src any) error {
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

func (x SeveritySlice) Value() (driver.Value, error) {
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

type SeverityType string

const (
	SeverityTypeCVSS2 SeverityType = "CVSS_V2"
	SeverityTypeCVSS3 SeverityType = "CVSS_V3"
)

// Severity
// Example:
//    {
//      "type": "CVSS_V3",
//      "score": "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:H"
//    }
type Severity struct {
	Type  SeverityType `json:"type" yaml:"type" db:"type" bson:"type" gorm:"column:type"`
	Score string       `json:"score" yaml:"score" db:"score" bson:"score" gorm:"column:score"`
}

var _ sql.Scanner = &Severity{}
var _ driver.Valuer = &Severity{}

func (x *Severity) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Severity) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
