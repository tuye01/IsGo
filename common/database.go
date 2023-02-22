package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"time"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.ports")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	loc := viper.GetString("datasource.loc")
	//charset := "uft8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		url.QueryEscape(loc))
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	//db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//自动创建数据表
	db.AutoMigrate(&model.User{})
	return db
}

func GetDB() *gorm.DB {
	DB = InitDB()
	return DB
}
