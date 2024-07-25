package main

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"

	_ "github.com/udistrital/documentos_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/customerror"
	"github.com/udistrital/utils_oas/xray"
)

func main() {
	orm.RegisterDataBase("default", "postgres",
		"postgres://"+beego.AppConfig.String("PGuser")+
			":"+url.QueryEscape(beego.AppConfig.String("PGpass"))+
			"@"+beego.AppConfig.String("PGurls")+
			":"+beego.AppConfig.String("PGport")+
			"/"+beego.AppConfig.String("PGdb")+
			"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas"))
	if beego.BConfig.RunMode == "dev" {
		// orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	/*logPath := "{\"filename\":\""
	logPath += beego.AppConfig.String("logPath")
	logPath += "\"}"
	logs.SetLogger(logs.AdapterFile, logPath)*/
	xray.InitXRay()
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.ErrorController(&customerror.CustomErrorController{})
	beego.Run()
}
