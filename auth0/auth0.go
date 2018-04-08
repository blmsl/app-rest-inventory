package auth0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	client        = &http.Client{}
	requestMutex  = &sync.Mutex{}
	apiTokenMutex = &sync.Mutex{}
)

type auth0Error struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	ErrorCode  string `json:"errorCode"`
}

type Auth0 interface {
	// Users.
	CreateUser(*User) (*User, error)
	DeleteUser(string) error

	// Groups.
	CreateGroup(string, string) (*Group, error)
	GetGroup(string) (*Group, error)
	NestGroups(string, ...string) error
	GetNestedGroups(string) ([]*Group, error)
	AddGroupMembers(string, ...string) error
	GetGroupMembers(string) (*Members, error)
	GetNestedGroupMembers(string) (*NestedMembers, error)
	GetGroupMember(string, string) (*User, error)
	GetNestedGroupMember(string, string) (*NestedMember, error)
	DeleteGroupMembers(string, ...string) error
	GetUserGroups(string) ([]*Group, error)
}

type auth0 struct {
	managementApi             Auth0Api
	authorizationExtensionApi Auth0Api
}

type Auth0Api interface {
	postAuth0Api(string, interface{}, interface{}) error
	patchAuth0Api(string, interface{}, interface{}) error
	getAuth0Api(string, interface{}) error
	deleteAuth0Api(string, interface{}, interface{}) error
	doAuth0Api(string, string, map[string]string, interface{}, interface{}) error
	updateToken() error
	getToken() (*TokenResponse, error)
	tokenExpired() bool
}

type auth0Api struct {
	domain       string
	clientId     string
	clientSecret string
	audience     string
	url          string

	_accessToken string
	_takenAt     time.Time
	_expiresIn   int
}

type Auth0Builder interface {
	ManagementApi(string, string, string, string, string) Auth0Builder
	AuthorizationExtensionApi(string, string, string, string, string) Auth0Builder
	Build() Auth0
}

type auth0Builder struct {
	managementApi             Auth0Api
	authorizationExtensionApi Auth0Api
}

func NewAuth0Builder() Auth0Builder {
	return &auth0Builder{}
}

func (a0b *auth0Builder) ManagementApi(domain, clientId, clientSecret, audience, url string) Auth0Builder {
	a0b.managementApi = &auth0Api{
		domain:       domain,
		clientId:     clientId,
		clientSecret: clientSecret,
		audience:     audience,
		url:          url}
	return a0b
}

func (a0b *auth0Builder) AuthorizationExtensionApi(domain, clientId, clientSecret, audience, url string) Auth0Builder {
	a0b.authorizationExtensionApi = &auth0Api{
		domain:       domain,
		clientId:     clientId,
		clientSecret: clientSecret,
		audience:     audience,
		url:          url}
	return a0b
}

func (a0b *auth0Builder) Build() Auth0 {
	return &auth0{
		managementApi:             a0b.managementApi,
		authorizationExtensionApi: a0b.authorizationExtensionApi}
}

func (a0Api *auth0Api) postAuth0Api(path string, request, response interface{}) error {
	// Build headers.
	headers := make(map[string]string)
	headers["content-type"] = "application/json"

	return a0Api.doAuth0Api(http.MethodPost, path, headers, request, response)
}

func (a0Api *auth0Api) patchAuth0Api(path string, request, response interface{}) error {
	// Build headers.
	headers := make(map[string]string)
	headers["content-type"] = "application/json"

	return a0Api.doAuth0Api(http.MethodPatch, path, headers, request, response)
}

func (a0Api *auth0Api) getAuth0Api(path string, response interface{}) error {
	// Build headers.
	headers := make(map[string]string)

	return a0Api.doAuth0Api(http.MethodGet, path, headers, nil, response)
}

func (a0Api *auth0Api) deleteAuth0Api(path string, request, response interface{}) error {
	// Build headers.
	headers := make(map[string]string)
	headers["content-type"] = "application/json"

	return a0Api.doAuth0Api(http.MethodDelete, path, headers, request, response)
}

func (a0Api *auth0Api) doAuth0Api(method, path string, headers map[string]string, request, response interface{}) error {
	// Verify access token.
	if a0Api.tokenExpired() {
		apiTokenMutex.Lock()
		if a0Api.tokenExpired() {
			// Update token.
			err := a0Api.updateToken()
			if err != nil {
				return err
			}
		}
		apiTokenMutex.Unlock()
	}

	// Add authorization header.
	if headers != nil {
		headers["authorization"] = "Bearer " + a0Api._accessToken
	} else {
		return fmt.Errorf("headers can't be nil. ")
	}

	// Build url.
	urlBuilder := bytes.NewBufferString(a0Api.url)
	urlBuilder.WriteString(path)

	return do(method, urlBuilder.String(), headers, request, response)
}

func do(method, url string, headers map[string]string, request, response interface{}) error {
	// Go single-threaded so we can deal with the rate limit.
	requestMutex.Lock()
	defer requestMutex.Unlock()

	// Build request.
	var req *http.Request
	var err error
	switch method {
	case http.MethodPost:
		var b []byte
		b, err = json.Marshal(request)
		if err != nil {
			return fmt.Errorf("It was not possible marshal request. ", err.Error())
		}
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	case http.MethodPatch:
		var b []byte
		b, err = json.Marshal(request)
		if err != nil {
			return fmt.Errorf("It was not possible marshal request. ", err.Error())
		}
		req, err = http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(b))
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, url, nil)
	case http.MethodDelete:
		if request != nil {
			var b []byte
			b, err = json.Marshal(request)
			if err != nil {
				return fmt.Errorf("It was not possible marshal request. ", err.Error())
			}
			req, err = http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(b))
		} else {
			req, err = http.NewRequest(http.MethodDelete, url, nil)
		}
	default:
		return errors.New("No method was specified.")
	}
	if err != nil {
		return fmt.Errorf("It was not possible create request. ", err.Error())
	}

	// Setup headers.
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// Send request.
	res, err := client.Do(req)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()

	// HTTP connection error.
	if err != nil {
		return fmt.Errorf("It was not possible send request. ", err.Error())
	}

	// Verify status.
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errorResponse auth0Error
		err = json.NewDecoder(res.Body).Decode(&errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("Code (%d): %s", errorResponse.StatusCode, errorResponse.Message)
	}

	// No content.
	if res.StatusCode == 204 {
		return nil
	}

	// Unmarshall.
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return fmt.Errorf("It was not possible unmarshal response. ", err.Error())
	}

	// If everithing ok.
	return nil
}
