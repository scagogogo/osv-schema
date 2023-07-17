package osv_schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type CreditsType string

const (
	// CreditsTypeFinder FINDER: identified the vulnerability.
	CreditsTypeFinder CreditsType = "FINDER"

	// CreditsTypeReporter REPORTER: notified the vendor of the vulnerability to a CNA.
	CreditsTypeReporter CreditsType = "REPORTER"

	// CreditsTypeAnalyst ANALYST: validated the vulnerability to ensure accuracy or severity.
	CreditsTypeAnalyst CreditsType = "ANALYST"

	// CreditsTypeCoordinator COORDINATOR: facilitated the coordinated response process.
	CreditsTypeCoordinator CreditsType = "COORDINATOR"

	// CreditsTypeRemediationDeveloper REMEDIATION_DEVELOPER: prepared a code change or other remediation plans.
	CreditsTypeRemediationDeveloper CreditsType = "REMEDIATION_DEVELOPER"

	// CreditsTypeRemediationReviewer REMEDIATION_REVIEWER: reviewed vulnerability remediation plans or code changes for effectiveness and completeness.
	CreditsTypeRemediationReviewer CreditsType = "REMEDIATION_REVIEWER"

	// CreditsTypeRemediationVerifier REMEDIATION_VERIFIER: tested and verified the vulnerability or its remediation.
	CreditsTypeRemediationVerifier CreditsType = "REMEDIATION_VERIFIER"

	// CreditsTypeTool TOOL: names of tools used in vulnerability discovery or identification.
	CreditsTypeTool CreditsType = "TOOL"

	// CreditsTypeSponsor SPONSOR: supported the vulnerability identification or remediation activities.
	CreditsTypeSponsor CreditsType = "SPONSOR"

	// CreditsTypeOther OTHER: any other type or role that does not fall under the categories described above.
	CreditsTypeOther CreditsType = "OTHER"
)

type Credits struct {
	Name    string   `mapstructure:"name" json:"name" yaml:"name" db:"name" bson:"name" gorm:"column:name"`
	Contact []string `mapstructure:"contact" json:"contact" yaml:"contact" db:"contact" bson:"contact" gorm:"column:contact;serializer:json"`
	Type    string   `mapstructure:"type" json:"type" yaml:"type" db:"type" bson:"type" gorm:"column:type"`
}

var _ sql.Scanner = &Credits{}
var _ driver.Valuer = &Credits{}

func (x *Credits) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Credits) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return wrapScanError(src, x)
	}
	return json.Unmarshal(bytes, &x)
}
