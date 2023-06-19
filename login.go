package govortex

import (
	"context"
	"net/http"
)

// Login performs the login operation in the Vortex API.
// It takes a context, client code, password, and TOTP (Time-Based One-Time Password) as input.
// If the login is successful, the method updates the accessToken field of the VortexApi instance.
// It returns the LoginResponse and an error.
func (v *VortexApi) Login(ctx context.Context, clientCode string, password string, totp string) (LoginResponse, error) {
	data := map[string]string{
		"client_code":    clientCode,
		"password":       password,
		"totp":           totp,
		"application_id": v.applicationId,
	}

	var resp LoginResponse

	header := http.Header{}
	header.Add("x-api-key", v.apiKey)
	_, err := v.doJson(ctx, "POST", URILogin, data, nil, header, &resp)
	if err != nil {
		return resp, err
	}
	v.AccessToken = resp.Data.AccessToken
	return resp, nil
}
