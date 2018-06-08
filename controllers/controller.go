package controllers

import (
	"app-rest-inventory/util/apierror"
	"github.com/astaxie/beego"
)

// serveError Serves an error response.
func serveError(c beego.Controller, status int, message string) {
	c.Data["json"] = apierror.ApiError{StatusCode: status, ErrorMessage: message}
	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()
	c.StopRun()
}
