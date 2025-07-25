package govortex

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
)

// Depricated: Use SSO Login instead.
// Login performs the login operation in the Vortex API.
// It takes a context, client code, password, and TOTP (Time-Based One-Time Password) as input.
// If the login is successful, the method updates the accessToken field of the VortexApi instance.
// It returns the LoginResponse and an error.
func (v *VortexApi) Login(ctx context.Context, clientCode string, password string, totp string) (*LoginResponse, error) {
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
		return nil, err
	}
	v.AccessToken = resp.Data.AccessToken
	return &resp, nil
}

// SSOLogin performs the login operation using Single Sign-On (SSO) in the Vortex API.
// It takes a context and a callback parameter as input.
// If the applicationId is not set, it returns an error.
// It returns a URL for the SSO login process.
// The User will be redirected to the callback url configured on the API Portal.
func (v *VortexApi) SSOLogin(ctx context.Context, callback_param string) (string, error) {
	if v.applicationId == "" {
		return "", fmt.Errorf("applicationId is not set")
	}

	return fmt.Sprintf("https://flow.rupeezy.in?applicationId=%s&cb_param=%s", v.applicationId, callback_param), nil
}

// ExchangeToken exchanges auth token received for the access_token from the Vortex API.
// It takes a context and auth_token as input.
// If the login is successful, the method updates the accessToken field of the VortexApi instance.
// It returns the LoginResponse and an error.
func (v *VortexApi) ExchangeToken(ctx context.Context, auth_token string) (*LoginResponse, error) {
	request := ExchangeAuthTokenRequest{
		Token:         auth_token,
		ApplicationId: v.applicationId,
		Checksum:      getSha256(v.applicationId + auth_token + v.apiKey),
	}
	var resp LoginResponse
	header := http.Header{}
	_, err := v.doJson(ctx, "POST", URISession, request, nil, header, &resp)
	if err != nil {
		return nil, err
	}
	v.AccessToken = resp.Data.AccessToken
	return &resp, nil
}

func getSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// Logout logs the user out from the Vortex API and the access token is then no longer useful.
// It takes a context as input.
// It returns the LogoutResponse and an error.
func (v *VortexApi) Logout(ctx context.Context) (*LogoutResponse, error) {
	var resp LogoutResponse
	_, err := v.doJson(ctx, "DELETE", URISession, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	v.AccessToken = ""
	return &resp, nil
}
