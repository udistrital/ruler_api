package main

import (
	_ "github.com/apdaza/rulerApiV3/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_pass")+"@"+beego.AppConfig.String("db_url")+"/"+beego.AppConfig.String("db_name")+"?sslmode=disable&search_path="+beego.AppConfig.String("db_schemas")+")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

