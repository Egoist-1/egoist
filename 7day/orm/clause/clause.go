package clause

import "strings"



//拼接各个独立的 子句

// Clause 内存储的就是 拼接后的sql 语句 sqlVars 是每个语句对应的参数
type Clause struct {
	sql     map[Type]string
	//每个sql语句类型对应的 参数
	sqlVars map[Type][]interface{}
}

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

//根据 type 调用对应的generator 生成该子句对应的SQL语句
func (c *Clause)Set(name Type,vars ...interface{})  {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql ,vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 方法根据传入的type的顺序,构造最终的sql语句
func (c *Clause) Builde(orders ...Type)(string,[]interface{}){
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql,ok := c.sql[order];ok {
			sqls = append(sqls,sql)
		vars = append(vars,c.sqlVars[order]...)
		}
		
	}
	return strings.Join(sqls," "),vars
	
}