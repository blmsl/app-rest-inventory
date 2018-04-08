package auth0

import (
	"fmt"
)

type Members struct {
	Total int    `json:"total,omitempty"`
	Users []User `json:"users,omitempty"`
}

type NestedMembers struct {
	Total        int            `json:"total,omitempty"`
	NestedMember []NestedMember `json:"nested,omitempty"`
}

type NestedMember struct {
	User  User  `json:"user,omitempty"`
	Group Group `json:"group,omitempty"`
}

type Group struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Id          string   `json:"_id,omitempty"`
	Members     []string `json:"members,omitempty"`
}

// @param name Group name.
// @param description Group description.
func (a0 *auth0) CreateGroup(name string, description string) (*Group, error) {
	group := &Group{Name: name,
		Description: description}

	err := a0.authorizationExtensionApi.postAuth0Api("groups", group, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// @Param id Group id.
func (a0 *auth0) GetGroup(id string) (*Group, error) {
	path := fmt.Sprintf("groups/%s", id)

	group := new(Group)

	err := a0.authorizationExtensionApi.getAuth0Api(path, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// @param id Parent group ID.
// @param ids Nested groups ID.
func (a0 *auth0) NestGroups(id string, ids ...string) error {
	path := fmt.Sprintf("groups/%s/nested", id)

	return a0.authorizationExtensionApi.patchAuth0Api(path, ids, nil)
}

// @param id Group id.
func (a0 *auth0) GetNestedGroups(id string) ([]*Group, error) {
	path := fmt.Sprintf("groups/%s/nested", id)

	nestedGroups := make([]*Group, 0)

	err := a0.authorizationExtensionApi.getAuth0Api(path, &nestedGroups)
	if err != nil {
		return nil, err
	}

	return nestedGroups, nil
}

// @param id Group ID.
// @param ids Members.
func (a0 *auth0) AddGroupMembers(id string, ids ...string) error {
	path := fmt.Sprintf("groups/%s/members", id)

	return a0.authorizationExtensionApi.patchAuth0Api(path, ids, nil)
}

// @Param id Group ID.
func (a0 *auth0) GetGroupMembers(id string) (*Members, error) {
	path := fmt.Sprintf("groups/%s/members", id)

	members := new(Members)

	err := a0.authorizationExtensionApi.getAuth0Api(path, members)
	if err != nil {
		return nil, err
	}

	return members, nil
}

// @Param id Group ID.
func (a0 *auth0) GetNestedGroupMembers(id string) (*NestedMembers, error) {
	path := fmt.Sprintf("groups/%s/members/nested", id)

	nestedMembers := new(NestedMembers)

	err := a0.authorizationExtensionApi.getAuth0Api(path, nestedMembers)
	if err != nil {
		return nil, err
	}

	return nestedMembers, nil
}

// @Param groupID Group ID.
// @Param userID User ID.
func (a0 *auth0) GetGroupMember(groupID, userID string) (*User, error) {
	// Get all the members first. Authorization Extension API does not have a
	// method to get a single user from specific group.
	// TODO: Look for better way to accomplish the task.
	members, err := a0.GetGroupMembers(groupID)
	if err != nil {
		return nil, err
	}

	// Filter members.
	for _, user := range members.Users {
		if user.UserId == userID {
			return &user, nil
		}
	}

	// User not found error.
	return nil, fmt.Errorf("User %s not found in group %s.", userID, groupID)
}

// @Param groupID Group ID.
// @Param userID User ID.
func (a0 *auth0) GetNestedGroupMember(groupID, userID string) (*NestedMember, error) {
	// Get all the members first. Authorization Extension API does not have a
	// method to get a single user from specific group.
	// TODO: Look for better way to accomplish the task.
	nestedMembers, err := a0.GetNestedGroupMembers(groupID)
	if err != nil {
		return nil, err
	}

	// Filter nested members.
	for _, nestedMember := range nestedMembers.NestedMember {
		if nestedMember.User.UserId == userID {
			/*fmt.Println(user.UserId)*/
			return &nestedMember, nil
		}
	}

	// User not found error.
	return nil, fmt.Errorf("User %s not found in group %s.", userID, groupID)
}

// @Param id Group id.
// @Param ids User IDs.
func (a0 *auth0) DeleteGroupMembers(id string, ids ...string) error {
	path := fmt.Sprintf("groups/%s/members", id)

	return a0.authorizationExtensionApi.deleteAuth0Api(path, ids, nil)
}

// @Param id User id.
func (a0 *auth0) GetUserGroups(id string) ([]*Group, error) {
	path := fmt.Sprintf("users/%s/groups", id)

	userGroups := make([]*Group, 0)

	err := a0.authorizationExtensionApi.getAuth0Api(path, &userGroups)
	if err != nil {
		return nil, err
	}

	return userGroups, nil
}
