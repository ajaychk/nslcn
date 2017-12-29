package db

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/nslcn/models"
)

func init() {

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=ajay dbname=nsdb password=Iotsc2 sslmode=disable")
	loc, _ := time.LoadLocation("Asia/Kolkata")
	orm.DefaultTimeLoc = loc
	orm.RegisterModel(new(models.Node))

	if err := orm.RunSyncdb("default", false, true); err != nil {
		beego.Error("error in creating node model:", err)
		return
	}
	//testSome()
}

func testSome() {
	o := orm.NewOrm()
	o.Using("default")

	node := new(models.Node)
	node.ID = "1xxxxxx2"
	fmt.Println(o.Insert(node))
	node.ID = "2xxxxxx2"
	fmt.Println(o.Insert(node))
	node.ID = "3xxxxxx2"
	fmt.Println(o.Insert(node))
}
