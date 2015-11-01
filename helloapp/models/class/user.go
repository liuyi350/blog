package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `orm:"pk"`
	Nick  string
	Email string
}

func TestORM() {
	o := orm.NewOrm()

	u := User{"liuyi", "test", "34534532@qq.com"}
	o.Insert(&u)

	u1 := User
	o.Read(&u1)

	fmt.Println(u1)
}