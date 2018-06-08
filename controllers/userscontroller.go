package controllers

import (
	"app-rest-inventory/auth0"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	"strings"
)

// Users API
type UsersController struct {
	beego.Controller
}

func (c *UsersController) URLMapping() {
	c.Mapping("CreateUser", c.CreateUser)
	c.Mapping("GetUsers", c.GetUsers)
}

// @Title CreateUser
// @Description Create user.
// @Accept json
// @Success 200 {object} auth0.User
// @router / [post]
func (c *UsersController) CreateUser() {
	// Unmarshall request.
	user := new(auth0.User)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, user)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusBadRequest, err.Error())
	}

	// Validate app_metadata.
	customerId := user.AppMetadata["customer_id"].(string)
	authorization := user.AppMetadata["authorization"].(map[string]interface{})
	roles := authorization["roles"].([]interface{})

	// Add connection.
	user.Connection = auth0.UsernamePasswordAuthentication

	// Create user.
	u, err := auth0.Auth.CreateUser(user)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusInternalServerError, err.Error())
	}

	// Gorutine to process the user grooups.
	go func(customerID, userID string) {

		// Get nested groups of the customer.
		nestedGroups, err := auth0.Auth.GetNestedGroups(customerID)
		if err != nil {
			logs.Error(err.Error())
			c.StopRun()
		}
		for i := range nestedGroups {
			nGroup := nestedGroups[i]

			for j := range roles {
				role := roles[j].(string)

				if strings.Contains(nGroup.Name, role) {
					// Add user to the group.
					auth0.Auth.AddGroupMembers(nGroup.Id, userID)
				}
			}
		}
	}(customerId, u.UserId)

	// Serve JSON.
	c.Data["json"] = u
	c.ServeJSON()
}

// @Title GetUser
// @Description Get user.
// @Param	user_id	path	string	true	"User id."
// @Success 200 {object} auth0.User
// @router /:user_id [get]
func (c *UsersController) GetUser(user_id *string) {
	// Validate the tenant getting user from the tenant group and not from the
	// management api.

	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusUnauthorized, err.Error())
	}

	// Validate user ID.
	if user_id == nil {
		err := fmt.Errorf("user_id can not be empty.")
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusBadRequest, err.Error())
	}

	var user *auth0.User

	// Get nested groups of the customer.
	nestedGroups, err := auth0.Auth.GetNestedGroups(customerID)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusInternalServerError, err.Error())
	}

	// Search the user in the groups.
	for _, group := range nestedGroups {

		// Search in the nested group.
		user, err := auth0.Auth.GetGroupMember(group.Id, *user_id)
		if err != nil {
			logs.Error(err.Error())
			serveError(c.Controller, http.StatusInternalServerError, err.Error())
		}

		// Validate user.
		if user != nil {
			break
		}

	}

	if user == nil {
		err := fmt.Errorf("user_id is incorrect or user is not created.")
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusBadRequest, err.Error())
	}

	// Serve JSON
	c.Data["json"] = user
	c.ServeJSON()

}

// @Title GetUsers
// @Description Get users.
// @Success 200 {object} auth0.Members
// @router / [get]
func (c *UsersController) GetUsers() {
	// Validate the tenant getting user from the tenant group and not from the
	// management api.

	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusUnauthorized, err.Error())
	}

	// Get nested groups of the customer.
	nestedGroups, err := auth0.Auth.GetNestedGroups(customerID)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusInternalServerError, err.Error())
	}

	// Members var.
	members := new(auth0.Members)
	members.Total = 0
	members.Users = make([]auth0.User, 0)

	// Build response.
	for _, group := range nestedGroups {

		// Get nested group members.
		members_, err := auth0.Auth.GetGroupMembers(group.Id)
		if err != nil {
			logs.Error(err.Error())
			serveError(c.Controller, http.StatusInternalServerError, err.Error())
		}

		// If everithing ok update members.
		members.Total += members_.Total
		members.Users = append(members.Users, members_.Users...)
	}

	// Serve JSON
	c.Data["json"] = members
	c.ServeJSON()
}

// @Title UpdateUser
// @Description Update user.
// @Accept json
// @Param	user_id	path	string	true	"User id."
// @Success 200 {object} auth0.User
// @router /:user_id [patch]
func (c *UsersController) UpdateUser(user_id *string) {
	// TODO: Validate the tenant getting user from the tenant group.

	// Unmarshall request.
	user := new(auth0.User)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, user)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusBadRequest, err.Error())
	}

	// Validate user ID.
	if user_id == nil {
		err := fmt.Errorf("user_id can not be empty.")
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusBadRequest, err.Error())
	}

	user, err = auth0.Auth.UpdateUser(*user_id, user)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusInternalServerError, err.Error())
	}

	// Serve JSON
	c.Data["json"] = user
	c.ServeJSON()

}

// @Title DeleteUser
// @Description Delete user.
// @Param	user_id	path	string	true	"User id."
// @router /:user_id [delete]
func (c *UsersController) DeleteUser(user_id *string) {
	// TODO: Validate the tenant getting user from the tenant group.

	// Validate user ID.
	if user_id == nil {
		err := fmt.Errorf("user_id can not be empty.")
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusBadRequest, err.Error())
	}

	// Obtain user groups.
	userGroups, err := auth0.Auth.GetUserGroups(*user_id)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusInternalServerError, err.Error())
	}

	// Delete user from groups.
	u_id := *user_id
	for _, group := range userGroups {
		err_ := auth0.Auth.DeleteGroupMembers(group.Id, u_id)
		if err_ != nil {
			logs.Error(err_.Error())
			serveError(c.Controller, http.StatusInternalServerError, err.Error())
		}

	}

	// Delete user from tenant.
	err = auth0.Auth.DeleteUser(*user_id)
	if err != nil {
		logs.Error(err.Error())
		serveError(c.Controller, http.StatusInternalServerError, err.Error())
	}
}
