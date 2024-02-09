package govortex

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
)

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
	return &resp, nil
}

func getSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (v *VortexApi) Logout(ctx context.Context) (*map[string]interface{}, error) {
	var resp map[string]interface{}
	_, err := v.doJson(ctx, "POST", URISession, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
