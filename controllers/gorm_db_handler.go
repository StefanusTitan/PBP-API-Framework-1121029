package controllers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func gormConn() *gorm.DB {
	dsn := "root:5XqTQjJ@LfWX$mf8@tcp(localhost:3308)/db_musify?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("")
	return db
}
