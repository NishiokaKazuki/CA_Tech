package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine = XormConnect()

func GetDBConnect() *xorm.Engine {
	return engine
}

func XormConnect() *xorm.Engine {

	db, err := xorm.NewEngine("mysql", "kazuki:secret@tcp(localhost:13306)/ca_tech?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}
