package models

import (
"blog/helloapp/models/class"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", "root:jkjjyjss@tcp(localhost:3306)/helloapp?charset=utf8")
    orm.RegisterModel(new(class.User))
    orm.RunSyncdb("default", false, true)

}
