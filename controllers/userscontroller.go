package controllers

import (
	"app-rest-inventory/auth0"
	"encoding/json"
	"fmt"
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

	// Validate app_metadata.
	customerId := user.AppMetadata["customer_id"].(string)
	authorization := user.AppMetadata["authorization"].(map[string]interface{})
	roles := authorization["roles"].([]interface{})

	// Add connection.
	user.Connection = auth0.UsernamePasswordAuthentication

	// Create user.
	a0User, err := auth0.AUTH0.CreateUser(user)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Get nested groups of the customer.
	nestedGroups, err := auth0.AUTH0.GetNestedGroups(customerId)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	for i := range nestedGroups {
		nGroup := nestedGroups[i]

		for j := range roles {
			role := roles[j].(string)

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
	user.Nickname = a0User.Username
	user.Picture = a0User.Picture
	user.UpdatedAt = a0User.UpdatedAt
	user.CreatedAt = a0User.CreatedAt
	user.UserMetadata = a0User.UserMetadata
	user.AppMetadata = a0User.AppMetadata

	// Serve JSON.
	c.Data["json"] = user
	c.ServeJSON()
}

// @Param	user_id	path	string	false	"User id."
// @router /users/:user_id [get]
func (c *UsersController) GetUser(user_id *string) {
	// Validate the tenant getting user from the tenant group and not from the
	// management api.

	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate user ID.
	if user_id == nil {
		err := fmt.Errorf("user_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Get user.
	nestedMember, err := auth0.AUTH0.GetNestedGroupMember(customerID, *user_id)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON
	c.Data["json"] = nestedMember.User
	c.ServeJSON()

}

// @router /users [get]
func (c *UsersController) GetUsers() {
	// Validate the tenant getting user from the tenant group and not from the
	// management api.

	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Get group members.
	users, err := auth0.AUTH0.GetNestedGroupMembers(customerID)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON
	c.Data["json"] = users
	c.ServeJSON()
}

// @Param	user_id	path	string	false	"User id."
// @router /users/:user_id [delete]
func (c *UsersController) DeleteUser(user_id *string) {
	// Validate the tenant getting user from the tenant group and not from the
	// management api.

	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate user ID.
	if user_id == nil {
		err := fmt.Errorf("user_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Obtain user groups.
	userGroups, err := auth0.AUTH0.GetUserGroups(*user_id)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Delete user from groups.
	u_id := *user_id
	for _, group := range userGroups {
		err_ := auth0.AUTH0.DeleteGroupMembers(group.Id, u_id)
		if err_ != nil {
			logs.Error(err_.Error())
			c.Abort(err.Error())
		}

	}

	// Delete user from tenant.
	err = auth0.AUTH0.DeleteUser(*user_id)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON
	c.Data["json"] = ""
	c.ServeJSON()
}
