package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:@/restapicrud_go?charset=utf8&parseTime=True")

}
