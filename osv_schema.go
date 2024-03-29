package osv_schema

import (
	"time"
)

// OsvSchema 表示一个OSV格式的漏洞数据
// 参考文档： https://ossf.github.io/osv-schema/
type OsvSchema[EcosystemSpecific, DatabaseSpecific any] struct {

	// OSV的版本
	SchemaVersion string `mapstructure:"schema_version" json:"schema_version" yaml:"schema_version" db:"schema_version" bson:"schema_version" gorm:"column:schema_version"`
	ID            string `mapstructure:"id" json:"id" yaml:"id" db:"id" bson:"id" gorm:"column:id"`

	// 修改日期
	Modified time.Time `mapstructure:"modified" json:"modified" yaml:"modified" db:"modified" bson:"modified" gorm:"column:modified"`

	// 发布日期
	Published time.Time `mapstructure:"published" json:"published" yaml:"published" db:"published" bson:"published" gorm:"column:published"`

	// TODO 2023-5-23 19:10:45 草这个字段啥意思...
	Withdrawn string `mapstructure:"withdrawn" json:"withdrawn" yaml:"withdrawn" db:"withdrawn" bson:"withdrawn" gorm:"column:withdrawn"`

	// 漏洞的编号
	Aliases Aliases `mapstructure:"aliases" json:"aliases" yaml:"aliases" db:"aliases" bson:"aliases" gorm:"column:aliases;serializer:json"`

	Related Related `mapstructure:"related" json:"related" yaml:"related" db:"related" bson:"related" gorm:"column:related;serializer:json"`

	// 可以认为是漏洞标题啥的
	Summary string `mapstructure:"summary" json:"summary" yaml:"summary" db:"summary" bson:"summary" gorm:"column:summary"`

	// 可以认为是漏洞详情啥的
	Details string `mapstructure:"details" json:"details" yaml:"details" db:"details" bson:"details" gorm:"column:details"`

	// 漏洞的严重级别
	Severity SeveritySlice `mapstructure:"severity" json:"severity" yaml:"severity" db:"severity" bson:"severity" gorm:"column:severity;serializer:json"`

	// 漏洞的影响范围
	Affected AffectedSlice[EcosystemSpecific, DatabaseSpecific] `mapstructure:"affected" json:"affected" yaml:"affected" db:"affected" bson:"affected" gorm:"column:affected;serializer:json"`

	// 参考资料
	References References `mapstructure:"references" json:"references" yaml:"references" db:"references" bson:"references" gorm:"column:references;serializer:json"`

	// 漏洞库自己的实现规范
	DatabaseSpecific DatabaseSpecific `mapstructure:"database_specific" json:"database_specific" yaml:"database_specific" db:"database_specific" bson:"database_specific" gorm:"column:database_specific;serializer:json"`

	Credits *Credits `mapstructure:"credits" json:"credits" yaml:"credits" db:"credits" bson:"credits" gorm:"column:credits;serializer:json"`
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
