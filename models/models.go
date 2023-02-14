package models

import (
	"fmt"
	"gin_templete/conf"
	"time"

	"github.com/jinzhu/gorm"
)

//创建GORM对象
var DB *gorm.DB

//数据库配置
func SetUp() {
	//拼接数据库连接URI
	connURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?chaset=%s&parseTime=True&loc=Local",
		conf.DataBase.User,
		conf.DataBase.Password,
		conf.DataBase.Host,
		conf.DataBase.Port,
		conf.DataBase.DB,
		conf.DataBase.CharSet)
	//连接数据库
	db, err := gorm.Open(conf.DataBase.Type, connURI)
	if err != nil {
		panic(err)
	}
	DB = db

	//设置默认数据库前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DataBase.Prefix + defaultTableName
	}
}

type BaseModel struct {
	ID          uint64    `gorm:"primary_key" json:"id"`
	IsDel       bool      `json:"-"`
	CreatedTime time.Time `json:"-"`
	UpdatedTime time.Time `json:"-"`
}
