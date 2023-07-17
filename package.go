package osv_schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"
)

// ------------------------------------------------- --------------------------------------------------------------------

// Ecosystem 表示包管理器的类型，比如 Maven
type Ecosystem string

const (

	// EcosystemGo Go	The Go ecosystem; the name field is a Go module path.
	EcosystemGo Ecosystem = "Go"

	// EcosystemNpm npm	The NPM ecosystem; the name field is an NPM package name.
	EcosystemNpm Ecosystem = "npm"

	// EcosystemOSSFuzz OSS-Fuzz	For reports from the OSS-Fuzz project that have no more appropriate ecosystem;
	// the name field is the name assigned by the OSS-Fuzz project, as recorded in the submitted fuzzing configuration.
	EcosystemOSSFuzz Ecosystem = "OSS-Fuzz"

	// EcosystemPyPI PyPI	the Python PyPI ecosystem; the name field is a normalized PyPI package name.
	EcosystemPyPI Ecosystem = "PyPI"

	// EcosystemRubyGems RubyGems	The RubyGems ecosystem; the name field is a gem name.
	EcosystemRubyGems Ecosystem = "RubyGems"

	// EcosystemCratesIo crates.io	The crates.io ecosystem for Rust; the name field is a crate name.
	EcosystemCratesIo Ecosystem = "crates.io"

	// EcosystemPackagist Packagist	The PHP package manager ecosystem; the name is a package name.
	EcosystemPackagist Ecosystem = "Packagist"

	// EcosystemMaven Maven	The Maven Java package ecosystem. The name field is a Maven package name.
	EcosystemMaven Ecosystem = "Maven"

	// EcosystemNuGet NuGet	The NuGet package ecosystem. The name field is a NuGet package name.
	EcosystemNuGet Ecosystem = "NuGet"

	// EcosystemLinux Linux	The Linux kernel. The only supported name is Kernel.
	EcosystemLinux Ecosystem = "Linux"

	// EcosystemDebian Debian	The Debian package ecosystem; the name is the name of the source package. The ecosystem
	// string might optionally have a :<RELEASE> suffix to scope the package to a particular Debian release. <RELEASE>
	// is a numeric version specified in the Debian distro-info-data. For example, the ecosystem string “Debian:7” refers
	// to the Debian 7 (wheezy) release.
	EcosystemDebian Ecosystem = "Debian"

	// EcosystemAlpine Alpine	The Alpine package ecosystem; the name is the name of the source package.
	// The ecosystem string must have a :v<RELEASE-NUMBER> suffix to scope the package to a particular Alpine release
	// branch (the v prefix is required). E.g. v3.16.
	EcosystemAlpine Ecosystem = "Alpine"

	// EcosystemHex Hex	The package manager for the Erlang ecosystem; the name is a Hex package name.
	EcosystemHex Ecosystem = "Hex"

	// EcosystemAndroid Android	The Android ecosystem; the name field is the Android component name that the patch
	// applies to, as shown in the Android Security Bulletins such as Framework, Media Framework and Kernel Component.
	// The exhaustive list of components can be found at the Appendix.
	EcosystemAndroid Ecosystem = "Android"

	// EcosystemGitHubActions GitHub Actions	The GitHub Actions ecosystem; the name field is the action’s repository
	// name with owner e.g. {owner}/{repo}.
	EcosystemGitHubActions Ecosystem = "GitHub Actions"

	// EcosystemPub Pub	The package manager for the Dart ecosystem; the name field is a Dart package name.
	EcosystemPub Ecosystem = "Pub"

	// EcosystemConanCenter ConanCenter	The ConanCenter ecosystem for C and C++; the name field is a Conan package name.
	EcosystemConanCenter Ecosystem = "ConanCenter"

	// EcosystemRocky Rocky Linux	The Rocky Linux package ecosystem; the name is the name of the source package.
	// The ecosystem string might optionally have a :<RELEASE> suffix to scope the package to a particular Rocky Linux
	// release. <RELEASE> is a numeric version.
	EcosystemRocky Ecosystem = "Rocky"

	// EcosystemAlmaLinux AlmaLinux package ecosystem; the name is the name of the source package. The ecosystem string
	// might optionally have a :<RELEASE> suffix to scope the package to a particular AlmaLinux release. <RELEASE> is a
	// numeric version.
	EcosystemAlmaLinux Ecosystem = "AlmaLinux"
)

// ------------------------------------------------- --------------------------------------------------------------------

//	"package": {
//	  "ecosystem": "RubyGems",
//	  "name": "sprout"
//	},
type Package struct {

	// 包管理器类型
	Ecosystem Ecosystem `mapstructure:"ecosystem" json:"ecosystem" yaml:"ecosystem" db:"ecosystem" bson:"ecosystem" gorm:"column:ecosystem"`

	// 包的名字
	Name string `mapstructure:"name" json:"name" yaml:"name" db:"name" bson:"name" gorm:"column:name"`

	// https://github.com/package-url/purl-spec
	PackageUrl string `mapstructure:"purl" json:"purl" yaml:"purl" db:"purl" bson:"purl" gorm:"column:purl"`
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
		return wrapScanError(src, x)
	}
	return json.Unmarshal(bytes, &x)
}

// IsMaven 判断包的类型是否是Maven的包
func (x *Package) IsMaven() bool {
	return x.Ecosystem == EcosystemMaven
}

// GetGroupID 如果ecosystem是maven的话，则name是GroupId:ArtifactID这样拼接在一起的，提供两个单独获取的API
func (x *Package) GetGroupID() string {
	if x == nil {
		return ""
	}
	split := strings.SplitN(x.Name, ":", 2)
	if len(split) != 2 {
		return ""
	} else {
		return split[0]
	}
}

// GetArtifactID @see GetGroupID
func (x *Package) GetArtifactID() string {
	if x == nil {
		return ""
	}
	split := strings.SplitN(x.Name, ":", 2)
	if len(split) != 2 {
		return ""
	} else {
		return split[1]
	}
}

// ------------------------------------------------- --------------------------------------------------------------------
