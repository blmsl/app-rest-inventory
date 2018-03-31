package auth0

import (
	"fmt"
)

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
