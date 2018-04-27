package main

import (
	"app-rest-inventory/controllers"
	_ "app-rest-inventory/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	// Setup logs.
	setupLogs()

	// Setup CORS.
	setupCORS()

	// Setup static routes.
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	// Setup error handler.
	setupErrorHandler()

	// Run and serve.
	logs.Info("The app.rest is set up correctly.")
	logs.Info("Listen and serve at %s", beego.AppConfig.String("httpport"))
	beego.Run()
}

/** Setup application logger. */
func setupLogs() {
	// Async logs to improve performance.
	logs.Async()

	// Show method and line number configuration.
	// Set to false to improve performance.
	logs.EnableFuncCallDepth(true)

	// Setup logger.
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"app-rest-inventory.log", 
	"daily":true, "maxdays":7, "separate":["emergency", "alert", "critical", "error", 
	"warning", "notice", "info", "debug"]}`)
}

/** Setup CORS. */
func setupCORS() {
	// Allow CORS.
	// Allowed methods.
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

func setupErrorHandler() {
	beego.ErrorController(&controllers.ErrorController{})
}
