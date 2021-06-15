package dao

import (
	conf "blog-go-gin/config"
	"blog-go-gin/logging"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Db *gorm.DB
)

func InitMysql() {
	config := conf.GetConf()
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.DbName,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
			//IgnoreRecordNotFoundError: true,
		},
	)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		logging.Logger.Errorf("gorm open mysql failed, err:%s", err)
		return
	}

	sqlDB, _ := Db.DB()
	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
