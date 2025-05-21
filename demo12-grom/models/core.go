package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局变量，用于存储数据库连接
var err error

// 在执行包语句时，会自动触发包内部 init 函数的调用，并且不能在代码中主动调用它。注意，init 函数没有参数也没有返回值。
func init() {
	// 连接 MySQL 数据库：https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("failed to connect database!!!!!!") // 如果连接失败，则终止程序运行
	}
}
