package auth0

import (
	"fmt"
)

/*{
    "email": "test@test.com",
    "username": "test",
    "email_verified": false,
    "user_id": "auth0|5abe654c982c2043a19b85ae",
    "picture": "https://s.gravatar.com/avatar/f2c97b1f2d2898cd2d6466ce95d4ba33?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fte.png",
    "identities": [
        {
            "connection": "Username-Password-Authentication",
            "user_id": "5abe654c982c2043a19b85ae",
            "provider": "auth0",
            "isSocial": false
        }
    ],
    "updated_at": "2018-03-30T16:26:52.738Z",
    "created_at": "2018-03-30T16:26:52.738Z"
}*/
type User struct {
	UserId        string                 `json:"user_id,omitempty"`
	Connection    string                 `json:"connection,omitempty"`
	Email         string                 `json:"email,omitempty"`
	Username      string                 `json:"username,omitempty"`
	Nickname      string                 `json:"nickname,omitempty"`
	Password      string                 `json:"password,omitempty"`
	PhoneNumber   string                 `json:"phone_number,omitempty"`
	Picture       string                 `json:"picture,omitempty"`
	UserMetadata  map[string]interface{} `json:"user_metadata,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	VerifyEmail   bool                   `json:"verify_email,omitempty"`
	PhoneVerified bool                   `json:"phone_verified,omitempty"`
	AppMetadata   map[string]interface{} `json:"app_metadata,omitempty"`
	Identities    []Identity             `json:"identities,omitempty"`
	UpdatedAt     string                 `json:"updated_at,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
}

type Identity struct {
	Connection string `json:"connection,omitempty"`
	UserId     string `json:"user_id,omitempty"`
	Provider   string `json:"provider,omitempty"`
	IsSocial   bool   `json:"isSocial,omitempty"`
}

// Connections.
const (
	UsernamePasswordAuthentication = "Username-Password-Authentication"
)

// @Param user User.
func (a0 *auth0) CreateUser(user *User) (*User, error) {
	err := a0.managementApi.postAuth0Api("users", user, user)
	if err != nil {
		return nil, err
	}
	// Make user copy.
	u := new(User)
	*u = *user

	return u, nil
}

// @Param id User id.
func (a0 *auth0) DeleteUser(id string) error {
	path := fmt.Sprintf("users/%s", id)

	return a0.managementApi.deleteAuth0Api(path, nil, nil)
}
