package models

import (
	"blog/helloapp/models/class"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
<<<<<<< HEAD
    orm.RegisterDataBase("default", "mysql", "root:jkjjyjss@tcp(localhost:3306)/helloapp?charset=utf8")
    orm.RegisterModel(new(class.User))
    orm.RunSyncdb("default", false, true)
=======
	orm.RegisterDataBase("default", "mysql", "root:jkjjyjss(localhost:3306)/jblog?charset=utf8")
	orm.RegisterModel(new(class.User))
	orm.RunSyncdb("default", false, true)
>>>>>>> c2911829cb88925dad3c64af5f264a5910a2922c

}
