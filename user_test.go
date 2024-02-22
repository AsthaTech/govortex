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

func (ts *ErrorTestSuite) TestFunds(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Funds(ctx)
	checkError429(t, err)
}

func (ts *TestSuite) TestBanks(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	banks, err := ts.VortexApiClient.Banks(ctx)
	if err != nil {
		t.Errorf("Error while fetching banks. %v", err)
		return
	}

	if banks == nil {
		t.Errorf("Error while fetching banks. %s", "banks are empty")
		return
	}

	if len(banks.Data.Nse) == 0 {
		t.Errorf("Error while fetching funds. %s", "doesnt match value")
	}
}

func (ts *ErrorTestSuite) TestBanks_Error(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, err := ts.VortexApiClient.Banks(ctx)
	checkError429(t, err)
}
