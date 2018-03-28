package auth0

import (
	"github.com/astaxie/beego"
)

// This is how the wrapper must be used. In case you want to uncouple the wrapper
// you need to put this file in the main project side.

var AUTH0 Auth0

func init() {

	timeout, err := beego.AppConfig.Int("auth0::timeout")
	if err != nil {
		panic(err)
	}

	// Create builder.
	a0Builder := NewAuth0Builder().
		ManagementApi(
			beego.AppConfig.String("auth0::domain"),
			beego.AppConfig.String("auth0::clientid"),
			beego.AppConfig.String("auth0::clientsecret"),
			beego.AppConfig.String("auth0::managementapiaudience"),
			beego.AppConfig.String("auth0::managementapiurl"),
			timeout,
			timeout).
		AuthorizationExtensionApi(beego.AppConfig.String("auth0::domain"),
			beego.AppConfig.String("auth0::clientid"),
			beego.AppConfig.String("auth0::clientsecret"),
			beego.AppConfig.String("auth0::authorizationextensionapiaudience"),
			beego.AppConfig.String("auth0::authorizationextensionapiurl"),
			timeout,
			timeout)

	// Build.
	AUTH0 = a0Builder.Build()
}
