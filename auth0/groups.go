package auth0

import ()

type Group struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Id          string   `json:"_id,omitempty"`
	Members     []string `json:"members,omitempty"`
}

func (a0 *auth0) CreateGroup(name string, description string) (*Group, error) {
	group := &Group{Name: name,
		Description: description}

	err := a0.authorizationExtensionApi.postAuth0Api("groups", group, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}
