package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"sync"
	"time"
)

const (
	// In case timeout is not configured.
	DEFAULT_TIMEOUT = 30000
)

var (
	clientMutex   = &sync.Mutex{}
	requestMutex  = &sync.Mutex{}
	apiTokenMutex = &sync.Mutex{}
)

type auth0Error struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type Auth0 interface {
	CreateGroup(string, string) (*Group, error)
}

type auth0 struct {
	managementApi             Auth0Api
	authorizationExtensionApi Auth0Api
}

type Auth0Api interface {
	postAuth0Api(string, interface{}, interface{}) error
	getAuth0Api(string, interface{}) error
	updateToken() error
	getToken() (*tokenResponse, error)
	tokenExpired() bool
}

type auth0Api struct {
	domain           string
	clientId         string
	clientSecret     string
	audience         string
	url              string
	connectTimeout   int
	readWriteTimeout int

	_accessToken string
	_takenAt     time.Time
	_expiresIn   int

	_client *http.Client
}

type Auth0Builder interface {
	ManagementApi(string, string, string, string, string, int, int) Auth0Builder
	AuthorizationExtensionApi(string, string, string, string, string, int, int) Auth0Builder
	Build() Auth0
}

type auth0Builder struct {
	managementApi             Auth0Api
	authorizationExtensionApi Auth0Api
}

func NewAuth0Builder() Auth0Builder {
	return &auth0Builder{}
}

func (a0b *auth0Builder) ManagementApi(domain, clientId, clientSecret, audience, url string, connectTimeout, readWriteTimeout int) Auth0Builder {
	a0b.managementApi = &auth0Api{
		domain:           domain,
		clientId:         clientId,
		clientSecret:     clientSecret,
		audience:         audience,
		url:              url,
		connectTimeout:   connectTimeout,
		readWriteTimeout: readWriteTimeout}
	return a0b
}

func (a0b *auth0Builder) AuthorizationExtensionApi(domain, clientId, clientSecret, audience, url string, connectTimeout, readWriteTimeout int) Auth0Builder {
	a0b.authorizationExtensionApi = &auth0Api{
		domain:           domain,
		clientId:         clientId,
		clientSecret:     clientSecret,
		audience:         audience,
		url:              url,
		connectTimeout:   connectTimeout,
		readWriteTimeout: readWriteTimeout}
	return a0b
}

func (a0b *auth0Builder) Build() Auth0 {
	return &auth0{
		managementApi:             a0b.managementApi,
		authorizationExtensionApi: a0b.authorizationExtensionApi}
}

func (a0Api *auth0Api) postAuth0Api(path string, request, response interface{}) error {
	// Build url.
	urlBuilder := bytes.NewBufferString(a0Api.url)
	urlBuilder.WriteString(path)

	if a0Api.tokenExpired() {
		// Verify access token.
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

	logs.Debug(a0Api._accessToken)

	// Build headers.
	headers := make(map[string]string)
	headers["authorization"] = "Bearer " + a0Api._accessToken
	headers["content-type"] = "application/json"

	return post(urlBuilder.String(), headers, client(a0Api._client, a0Api.connectTimeout, a0Api.readWriteTimeout), request, response)
}

func (a0Api *auth0Api) getAuth0Api(path string, response interface{}) error {
	// Build url.
	urlBuilder := bytes.NewBufferString(a0Api.url)
	urlBuilder.WriteString(path)

	if a0Api.tokenExpired() {
		// Verify access token.
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

	// Build headers.
	headers := make(map[string]string)
	headers["authorization"] = "Bearer " + a0Api._accessToken

	return get(urlBuilder.String(), headers, client(a0Api._client, a0Api.connectTimeout, a0Api.readWriteTimeout), response)
}

func post(url string, headers map[string]string, client *http.Client, request, response interface{}) error {
	// Go single-threaded so we can deal with the rate limit.
	requestMutex.Lock()
	defer requestMutex.Unlock()

	// Build request.
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
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
		return err
	}

	// Verify status.
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errorResponse auth0Error
		err = json.NewDecoder(res.Body).Decode(errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("Code (%d): %s", errorResponse.StatusCode, errorResponse.Message)
	}

	// Unmarshall.
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return err
	}

	// If everithing ok.
	return nil
}

func get(url string, headers map[string]string, client *http.Client, response interface{}) error {
	// Go single-threaded so we can deal with the rate limit.
	requestMutex.Lock()
	defer requestMutex.Unlock()

	// Build request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
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
		return err
	}

	// Verify status.
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errorResponse auth0Error
		err = json.NewDecoder(res.Body).Decode(errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("Code (%d): %s", errorResponse.StatusCode, errorResponse.Message)
	}

	// Unmarshall.
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return err
	}

	// If everithing ok.
	return nil
}

func client(client *http.Client, connectTimeout, readWriteTimeout int) *http.Client {
	if client == nil {
		clientMutex.Lock()
		if client == nil {

			// Client timeout.
			clientTimeout := DEFAULT_TIMEOUT
			if readWriteTimeout > 0 {
				clientTimeout = readWriteTimeout
			}

			// Set up client.
			client = &http.Client{Transport: http.DefaultTransport,
				Timeout: time.Duration(clientTimeout)}
		}
		clientMutex.Unlock()
	}
	return client
}
