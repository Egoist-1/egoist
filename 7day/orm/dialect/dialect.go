package dialect

import (
	"reflect"
)

var dialectMap = map[string]Dialect{}


type Dialect interface {
	DataTypeOf(typ reflect.Value) string //go类型转换为数据库类型
	TableExistSQL(tablename string)(string,[]interface{}) //返回表是否存在的sql语言
}



// RegistDialect 新增数据库支持时 可以注册到全局
func RegistDialect(name string,dialect Dialect)  {
	dialectMap[name] = dialect
	
}
// GetDialect 获取dialect 实例
func GetDialect(name string)(dialect Dialect,ok bool) {
	dialect,ok = dialectMap[name]
	return 
}
