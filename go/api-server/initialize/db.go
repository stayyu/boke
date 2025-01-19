package initialize

import (
	"boke/go/api-server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {

	dsn := "root:root@tcp(192.168.1.130:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	//全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//
	//// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	//global.DB.AutoMigrate(&model.Admin{})
	//global.DB.AutoMigrate(&model.Category{})
	//global.DB.AutoMigrate(&model.Blog{})

}
