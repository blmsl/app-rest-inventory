package auth0

import (
	"bytes"
	"github.com/astaxie/beego/logs"
	"time"
)

type tokenRequest struct {
	GrantType    string `json:"grant_type,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Audience     string `json:"audience,omitempty"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	Scope       string `json:"scope,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	TookenType  string `json:"token_type,omitempty"`
}

func (a0Api *auth0Api) updateToken() error {
	// Get token.
	tokenResponse, err := a0Api.getToken()
	if err != nil {
		return err
	}

	// Update token.
	a0Api._accessToken = tokenResponse.AccessToken
	a0Api._takenAt = time.Now()
	a0Api._expiresIn = tokenResponse.ExpiresIn

	return nil
}

func (a0Api *auth0Api) getToken() (*tokenResponse, error) {
	// Build request.
	request := &tokenRequest{
		GrantType:    "client_credentials",
		ClientId:     a0Api.clientId,
		ClientSecret: a0Api.clientSecret,
		Audience:     a0Api.audience}

	// Build url.
	urlBuilder := bytes.NewBufferString(a0Api.domain)
	urlBuilder.WriteString("oauth/token")
	logs.Debug(urlBuilder.String())

	// Response DTO.
	response := &tokenResponse{}

	// Build headers.
	headers := make(map[string]string)
	headers["content-type"] = "application/json"

	// Send request.
	err := post(urlBuilder.String(), headers, client(a0Api._client, a0Api.connectTimeout, a0Api.readWriteTimeout), request, response)
	if err != nil {
		return nil, err
	}

	// Return token response.
	return response, nil
}

// True if the token expired, false otherwise.
func (a0Api *auth0Api) tokenExpired() bool {
	// Verify current acces token.
	if len(a0Api._accessToken) > 0 {
		return true
	}

	// Verify taken at time.
	zero := time.Time{}
	if a0Api._takenAt == zero {
		return true
	}

	expiresAt := a0Api._takenAt.Add(time.Duration(a0Api._expiresIn))
	return expiresAt.Before(time.Now())
}
