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

// Schema
type Schema struct {
	Model      interface{}       //对象的Model
	Name       string            //表名Name
	Fields     []*Field          //字段
	FieldNames []string          //包含所有字段名,
	FieldMap   map[string]*Field //记录字段和Field的映射关系
}

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