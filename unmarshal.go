package osv_schema

import (
	"encoding/json"
	"os"
)

// UnmarshalFromJson 从JSON字符串中反序列化
func UnmarshalFromJson[EcosystemSpecific, DatabaseSpecific any](jsonBytes []byte) (*OsvSchema[EcosystemSpecific, DatabaseSpecific], error) {
	r := &OsvSchema[EcosystemSpecific, DatabaseSpecific]{}
	err := json.Unmarshal(jsonBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// UnmarshalFromJsonFile UnmarshalFromJson 从JSOn文件中反序列化
func UnmarshalFromJsonFile[EcosystemSpecific, DatabaseSpecific any](jsonFilePath string) (*OsvSchema[EcosystemSpecific, DatabaseSpecific], error) {
	fileBytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}
	return UnmarshalFromJson[EcosystemSpecific, DatabaseSpecific](fileBytes)
}
