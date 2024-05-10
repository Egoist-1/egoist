package schema

import (
	"7day/orm/dialect"
	"go/ast"
	"reflect"
)

// 对象与表结构 映射

type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema  代表一个表的结构 
type Schema struct {
	Model      interface{}       //对象的Model
	Name       string            //表名Name
	Fields     []*Field          //字段 
	FieldNames []string          //包含所有字段名,  
	FieldMap   map[string]*Field //记录字段和Field的映射关系
}
// GetField 根据表字段名 获取对应的 结构体字段
func (schema *Schema) GetField(name string) *Field {
	return schema.FieldMap[name]
}

// Parse 将任意对象解析成Schema
func Parse(dest interface{},d dialect.Dialect) *Schema{
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model: dest,
		Name: modelType.Name(),
		FieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v,ok := p.Tag.Lookup("geeorm");ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields,field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.FieldMap[p.Name] = field
		}
	}
	return schema
}
// RecordValues  根据数据库列的顺序 找到一一对应的 结构体字段 按顺序平铺
func (schema *Schema)RecordValues(dest interface{})[]interface{}{
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}