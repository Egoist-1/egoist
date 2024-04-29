package orm

import (
	"7day/orm/log"
	"7day/orm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}

// NewEngine 建立db连接
func NewEngine(driver, source string) (*Engine, error) {
	//建立连接
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error("初始化失败", err)
		return nil, err
	}
	//ping 测试

	if err = db.Ping(); err != nil {
		log.Error("ping test failed", err)
		return nil, err
	}
	//初始化Engine
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
	return session.New(e.db)
}