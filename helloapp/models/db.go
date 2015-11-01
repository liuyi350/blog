package models

import (
"blog/helloapp/models/class"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MYSQL)
    orm.RegisterDataBase("default", "mysql", "root:jkjjyjss(localhost:3306)/
    	jblog?charset=utf8")
    orm.RegisterModel(new(class.User))
    orm.RunSyncdb("default", false, true)

}
