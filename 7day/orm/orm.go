package orm

import (
	"7day/orm/dialect"
	"7day/orm/log"
	"7day/orm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}

// NewEngine 建立db连接
func NewEngine(driver, source string) (e *Engine, err error) {
	//建立连接
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error("初始化失败", err)
		return 
	}
	//ping 测试

	if err = db.Ping(); err != nil {
		log.Error("ping test failed", err)
		return 
	}
	//初始化Engine
	dial ,ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db,dialect: dial}
	log.Info("Connect database success")
	return &Engine{db: db}, nil
}

// Close
func (e *Engine) Close(){
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("close database sucess")
}
// NewSession
func (e *Engine)NewSession() *session.Session{
	return session.New(e.db,e.dialect)
}