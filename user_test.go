package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestFunds(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	funds, err := ts.VortexApiClient.Funds(ctx)
	if err != nil {
		t.Errorf("Error while fetching funds. %v", err)
	}

	if funds.NSE.Deposit != 22449.64 {
		t.Errorf("Error while fetching funds. %s", "doesnt match value")
	}
}
