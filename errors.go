package osv_schema

import (
	"fmt"
	"reflect"
)

// 生成scan错误
func wrapScanError(src, dest any) error {
	return fmt.Errorf("can not scan from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(dest).Name())
}
