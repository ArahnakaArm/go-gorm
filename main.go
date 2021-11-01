package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:12345678@tcp(localhost)/gorm?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)

	if err != nil {
		panic(err)
	}

	err = db.Migrator().CreateTable(Test{})

	if err != nil {
		fmt.Printf(err.Error())
	}

}

type Test struct {
	ID        uint
	Name      string
	CreatedAt string
}
