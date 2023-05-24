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
	SchemaVersion string    `json:"schema_version" yaml:"schema_version" db:"schema_version" bson:"schema_version"`
	ID            string    `json:"id" yaml:"id" db:"id" bson:"id"`
	Modified      time.Time `json:"modified" yaml:"modified" db:"modified" bson:"modified"`
	Published     time.Time `json:"published" yaml:"published" db:"published" bson:"published"`

	// TODO 2023-5-23 19:10:45 草这个字段啥意思...
	Withdrawn string `json:"withdrawn" yaml:"withdrawn" db:"withdrawn" bson:"withdrawn"`

	Aliases          Aliases                                            `json:"aliases" yaml:"aliases" db:"aliases" bson:"aliases"`
	Related          Related                                            `json:"related" yaml:"related" db:"related" bson:"related"`
	Summary          string                                             `json:"summary" yaml:"summary" db:"summary" bson:"summary"`
	Details          string                                             `json:"details" yaml:"details" db:"details" bson:"details"`
	Severity         SeveritySlice                                      `json:"severity" yaml:"severity" db:"severity" bson:"severity"`
	Affected         AffectedSlice[EcosystemSpecific, DatabaseSpecific] `json:"affected" yaml:"affected" db:"affected" bson:"affected"`
	References       References                                         `json:"references" yaml:"references" db:"references" bson:"references"`
	DatabaseSpecific DatabaseSpecific                                   `json:"database_specific" yaml:"database_specific" db:"database_specific" bson:"database_specific"`
	Credits          *Credits                                           `json:"credits" yaml:"credits" db:"credits" bson:"credits"`
}

var _ sql.Scanner = &OsvSchema[any, any]{}
var _ driver.Valuer = &OsvSchema[any, any]{}

// AffectedHasEcosystem 判断被影响到的包是否有包含给定的包管理器的，一般用于过滤
func (x *OsvSchema[EcosystemSpecific, DatabaseSpecific]) AffectedHasEcosystem(ecosystem string) bool {
	for _, item := range x.Affected {
		if item.Package != nil && item.Package.Ecosystem == ecosystem {
			return true
		}
	}
	return false
}

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
