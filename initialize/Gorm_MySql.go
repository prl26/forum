package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormMysql()(db *gorm.DB){
	dsn := "root:slxwan1314@tcp(127.0.0.1:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err !=nil{
		panic(err)
	}
	return db
}
