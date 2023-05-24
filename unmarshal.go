package osv_schema

import "encoding/json"

// UnmarshalFromJson 从JSON字符串中反序列化
func UnmarshalFromJson[EcosystemSpecific, DatabaseSpecific any](jsonBytes []byte) (*OsvSchema[EcosystemSpecific, DatabaseSpecific], error) {
	r := &OsvSchema[EcosystemSpecific, DatabaseSpecific]{}
	err := json.Unmarshal(jsonBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
