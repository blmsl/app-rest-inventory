package controllers

import (
	"app-rest-inventory/auth0"
	"app-rest-inventory/util/stringutil"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

// User API
type UsersController struct {
	beego.Controller
}

func (c *UsersController) URLMapping() {
	c.Mapping("CreateCustomer", c.CreateUser)
}

// @router /users [post]
func (c *UsersController) CreateUser() {
	// Unmarshall request.
	user := new(auth0.User)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, user)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Create user.
	a0User, err := auth0.AUTH0.CreateUser(auth0.UsernamePasswordAuthentication, user.Email,
		user.Username, user.Password, stringutil.Empty, nil, nil)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate app_metadata.
	customerId := a0User.AppMetadata["customer_id"].(string)
	roles := a0User.AppMetadata["roles"].([]string)

	// Get nested groups of the customer.
	nestedGroups, err := auth0.AUTH0.GetNestedGroups(customerId)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	for i := range nestedGroups {
		nGroup := nestedGroups[i]

		for j := range roles {
			role := roles[j]

			if strings.Contains(nGroup.Name, role) {
				// Add user to the group.
				auth0.AUTH0.AddGroupMembers(nGroup.Id, a0User.UserId)
			}
		}
	}

	// Response.
	user = new(auth0.User)
	user.UserId = a0User.UserId
	user.Email = a0User.Email
	user.Username = a0User.Username
	user.Picture = a0User.Picture
	user.UpdatedAt = a0User.UpdatedAt
	user.CreatedAt = a0User.CreatedAt
	user.UserMetadata = a0User.UserMetadata
	user.AppMetadata = a0User.AppMetadata

	// Serve JSON.
	c.Data["json"] = user
	c.ServeJSON()
}
