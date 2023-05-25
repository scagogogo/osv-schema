package osv_schema

import (
	"time"
)

// OsvSchema 表示一个OSV格式的漏洞数据
// 参考文档： https://ossf.github.io/osv-schema/
type OsvSchema[EcosystemSpecific, DatabaseSpecific any] struct {

	// OSV的版本
	SchemaVersion string    `json:"schema_version" yaml:"schema_version" db:"schema_version" bson:"schema_version" gorm:"schema_version"`
	ID            string    `json:"id" yaml:"id" db:"id" bson:"id" gorm:"id"`
	Modified      time.Time `json:"modified" yaml:"modified" db:"modified" bson:"modified" gorm:"modified"`
	Published     time.Time `json:"published" yaml:"published" db:"published" bson:"published" gorm:"published"`

	// TODO 2023-5-23 19:10:45 草这个字段啥意思...
	Withdrawn string `json:"withdrawn" yaml:"withdrawn" db:"withdrawn" bson:"withdrawn" gorm:"withdrawn"`

	Aliases          Aliases                                            `json:"aliases" yaml:"aliases" db:"aliases" bson:"aliases" gorm:"aliases;serializer:json"`
	Related          Related                                            `json:"related" yaml:"related" db:"related" bson:"related" gorm:"related;serializer:json"`
	Summary          string                                             `json:"summary" yaml:"summary" db:"summary" bson:"summary" gorm:"summary"`
	Details          string                                             `json:"details" yaml:"details" db:"details" bson:"details" gorm:"details"`
	Severity         SeveritySlice                                      `json:"severity" yaml:"severity" db:"severity" bson:"severity" gorm:"severity;serializer:json"`
	Affected         AffectedSlice[EcosystemSpecific, DatabaseSpecific] `json:"affected" yaml:"affected" db:"affected" bson:"affected" gorm:"affected;serializer:json"`
	References       References                                         `json:"references" yaml:"references" db:"references" bson:"references" gorm:"references;serializer:json"`
	DatabaseSpecific DatabaseSpecific                                   `json:"database_specific" yaml:"database_specific" db:"database_specific" bson:"database_specific" gorm:"database_specific;serializer:json"`
	Credits          *Credits                                           `json:"credits" yaml:"credits" db:"credits" bson:"credits" gorm:"credits;serializer:json"`
}

//var _ sql.Scanner = &OsvSchema[any, any]{}
//var _ driver.Valuer = &OsvSchema[any, any]{}
//
//func (x *OsvSchema[EcosystemSpecific, DatabaseSpecific]) Value() (driver.Value, error) {
//	if x == nil {
//		return nil, nil
//	}
//	return json.Marshal(x)
//}
//
//func (x *OsvSchema[EcosystemSpecific, DatabaseSpecific]) Scan(src any) error {
//	if src == nil {
//		return nil
//	}
//	bytes, ok := src.([]byte)
//	if !ok {
//		return fmt.Errorf("can not unmarshal from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
//	}
//	return json.Unmarshal(bytes, &x)
//}
