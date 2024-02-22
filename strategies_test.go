package govortex

import (
	"context"
	"testing"
)

// func (v *VortexApi) PlaceGttOrder() {
// }
func (ts *TestSuite) TestGetStrategies(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.GetStrategies(ctx, StrategiesRequest{})
	if err != nil {
		t.Errorf("Error while fetching strategies. %v", err)
		return
	}
	if len(resp.Data.Strategies) == 0 {
		t.Errorf("Error while fetching strategies. %s", "length is 0")
	}
}

func (ts *TestSuite) TestOptionChain(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.GetOptionChain(ctx, OptionChainRequest{})
	if err != nil {
		t.Errorf("Error while fetching option chain %v", err)
		return
	}
	if len(resp.Response.Options.List) == 0 {
		t.Errorf("Error while fetching option chain %s", "length is 0")
	}
}
func (ts *TestSuite) TestBuildStrategies(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.BuildStrategy(ctx, StrategyBuilderRequest{})
	if err != nil {
		t.Errorf("Error while fetching strategies. %v", err)
		return
	}
	if len(resp.Data.Strategies) == 0 {
		t.Errorf("Error while fetching strategies. %s", "length is 0")
	}
}

func (ts *TestSuite) TestPayoff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.GetPayoff(ctx, PayoffRequest{})
	if err != nil {
		t.Errorf("Error while fetching payoffs. %v", err)
		return
	}
	if len(resp.Data.PayOffs) == 0 {
		t.Errorf("Error while fetching payoffs. %s", "length is 0")
	}
}

func (ts *ErrorTestSuite) TestGetStrategies(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.GetStrategies(ctx, StrategiesRequest{})
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestOptionChain(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.GetOptionChain(ctx, OptionChainRequest{})
	checkError429(t, err)
}
func (ts *ErrorTestSuite) TestBuildStrategies(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.BuildStrategy(ctx, StrategyBuilderRequest{})
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestPayoff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.GetPayoff(ctx, PayoffRequest{})
	checkError429(t, err)
}
