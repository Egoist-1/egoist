package _api

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"webook/article/_internal/repository/dao"
)

func TestViper(t *testing.T) {

	viper.SetConfigName("dev")  // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	viper.SetDefault("database.password", "456")
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	getString := viper.GetString("database.password")
	log.Print(getString)
	sub := viper.Sub("database")
	s := sub.GetString("password")
	log.Print(s)
}

// 在这里试一下 开启事务 连续连个 for update 是否会死锁
// ce shi jian xi suo
func TestName(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic("failed to connect database")
	}
	//return db
}
