package auth0

import (
	"github.com/astaxie/beego"
)

// This is how the wrapper must be used. In case you want to uncouple the wrapper
// you need to put this file in the main project side.

var Auth Auth0

func init() {

	// Create builder.
	a0Builder := NewAuth0Builder().
		ManagementApi(
			beego.AppConfig.String("auth0::domain"),
			beego.AppConfig.String("auth0::managementapiclientid"),
			beego.AppConfig.String("auth0::managementapiclientsecret"),
			beego.AppConfig.String("auth0::managementapiaudience"),
			beego.AppConfig.String("auth0::managementapiurl")).
		AuthorizationExtensionApi(beego.AppConfig.String("auth0::domain"),
			beego.AppConfig.String("auth0::authorizationextensionapiclientid"),
			beego.AppConfig.String("auth0::authorizationextensionapiclientsecret"),
			beego.AppConfig.String("auth0::authorizationextensionapiaudience"),
			beego.AppConfig.String("auth0::authorizationextensionapiurl"))

	// Build.
	Auth = a0Builder.Build()
}
