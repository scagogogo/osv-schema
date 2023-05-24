package osv

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// OsvSchema 表示一个OSV格式的漏洞数据
// 参考文档： https://ossf.github.io/osv-schema/
type OsvSchema[EcosystemSpecific, DatabaseSpecific any] struct {

	// OSV的版本
	SchemaVersion string    `json:"schema_version" yaml:"schema_version" db:"schema_version"`
	ID            string    `json:"id" yaml:"id" db:"id"`
	Modified      time.Time `json:"modified" yaml:"modified" db:"modified"`
	Published     time.Time `json:"published" yaml:"published" db:"published"`

	// TODO 2023-5-23 19:10:45 草这个字段啥意思...
	Withdrawn string `json:"withdrawn" yaml:"withdrawn" db:"withdrawn"`

	Aliases          []string                                         `json:"aliases" yaml:"aliases" db:"aliases"`
	Related          []string                                         `json:"related" yaml:"related" db:"related"`
	Summary          string                                           `json:"summary" yaml:"summary" db:"summary"`
	Details          string                                           `json:"details" yaml:"details" db:"details"`
	Severity         []*Severity                                      `json:"severity" yaml:"severity" db:"severity"`
	Affected         []*Affected[EcosystemSpecific, DatabaseSpecific] `json:"affected" yaml:"affected" db:"affected"`
	References       []*References                                    `json:"references" yaml:"references" db:"references"`
	DatabaseSpecific DatabaseSpecific                                 `json:"database_specific" yaml:"database_specific" db:"database_specific"`
	Credits          *Credits                                         `json:"credits" yaml:"credits" db:"credits"`

	// 原始的osv格式的json的哈希
	OsvHash string `json:"osv_hash" yaml:"osv_hash" db:"osv_hash"`
}

var _ sql.Scanner = &OsvSchema[any, any]{}
var _ driver.Valuer = &OsvSchema[any, any]{}

func (x *OsvSchema[EcosystemSpecific, DatabaseSpecific]) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *OsvSchema[EcosystemSpecific, DatabaseSpecific]) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
