package session

import (
	"7day/orm/log"
	"database/sql"
	"strings"
)

//数据库交互
type Session struct {
	db      *sql.DB
	sql     strings.Builder //拼接sql语句
	sqlVars []interface{}   //语句中对应的占位符
}

func New(db *sql.DB) *Session{
	return &Session{db: db}
}

func (s *Session)Clear()  {
	s.sql.Reset()
	s.sqlVars = nil
}
func (s *Session)DB() *sql.DB{
	return s.db
}
func (s *Session)Raw(sql string, values ...interface{}) *Session{
	s.sql.WriteString(sql)
	s.sql.WriteString("")
	s.sqlVars = append(s.sqlVars,values...)
	return s
}

func (s *Session)Exec() (result sql.Result,err error){
	defer s.Clear()
	log.Info(s.sql.String(),s.sqlVars)
	if result,err =s.db.Exec(s.sql.String(),s.sqlVars...);err!=nil {
		log.Error(err)
	}
	return
	
}
func (s *Session)QueryRow() (row *sql.Row){
	defer s.Clear()
	log.Info(s.sql.String(),s.sqlVars)
	return s.db.QueryRow(s.sql.String(),s.sqlVars...)
}
func (s *Session)QueryRows()(rows *sql.Rows, err error){
	defer s.Clear()
	log.Info(s.sql.String(),s.sqlVars)
	rows,err = s.db.Query(s.sql.String(),s.sqlVars...)
	if err!=nil{
		log.Error(err)
	}
	return
}


