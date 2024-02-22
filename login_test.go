package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestLogin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.Login(ctx, "DEMO", "DEMOPASS", "123456")
	if err != nil {
		t.Errorf("Error while logging in. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while logging in. %s", "status is not success")
	}
}
func (ts *TestSuite) TestLogout(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.Logout(ctx)
	if err != nil {
		t.Errorf("Error while logging out. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while logging out. %s %s", "status is not success", resp.Status)
	}
}
func (ts *TestSuite) TestExchangeToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.ExchangeToken(ctx, "auth_token")
	if err != nil {
		t.Errorf("Error while exchanging token. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while exchanging token. %s", "status is not success")
	}
}

func (ts *ErrorTestSuite) TestLogin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Login(ctx, "DEMO", "DEMOPASS", "123456")
	checkError429(t, err)
}
func (ts *ErrorTestSuite) TestLogout(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Logout(ctx)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestExchangeToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.ExchangeToken(ctx, "auth_token")
	checkError429(t, err)
}
