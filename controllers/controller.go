package controllers

import (
	"app-rest-inventory/util/apierror"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// serveError Serves an error response.
// @Param status HTTP response status.
// @Param message Error message.
func (c *BaseController) serveError(status int, message string) {
	c.Data["json"] = apierror.ApiError{StatusCode: status, ErrorMessage: message}
	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()
	c.StopRun()
}
