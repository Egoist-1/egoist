package dialect

import (
	"reflect"
	"fmt"
	"time"
)
type sqlite3 struct{}

var _ Dialect =  (*sqlite3)(nil)

// 自动注册 sqlite3 
func init(){
	RegistDialect("sqlite3",&sqlite3{})
}
// DataTypeOf go语言映射 数据库 类型
func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _,ok := typ.Interface().(time.Time);ok{
			return "database"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// 判断sqlite 表是否存在
func (s *sqlite3) TableExistSQL(tablename string) (string, []interface {}) {
	args := []interface{}{tablename}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
