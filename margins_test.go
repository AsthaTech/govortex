package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestOrderMargin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := OrderMarginRequest{}
	resp, err := ts.VortexApiClient.OrderMargin(ctx, &request)
	if err != nil {
		t.Errorf("Error while getting order margin. %v", err)
		return
	}
	if resp.Available == 0 {
		t.Errorf("Error while getting order margin. %s", "available margin 0")
	}
}

func (ts *TestSuite) TestBasketMargin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := BasketMarginRequest{}
	resp, err := ts.VortexApiClient.BasketMargin(ctx, &request)
	if err != nil {
		t.Errorf("Error while getting basket margin. %v", err)
		return
	}
	if resp.InitialMargin == 0 {
		t.Errorf("Error while getting basket margin. %s", "initial margin 0")
	}
}

func (ts *ErrorTestSuite) TestOrderMargin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := OrderMarginRequest{}
	_, err := ts.VortexApiClient.OrderMargin(ctx, &request)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestBasketMargin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := BasketMarginRequest{}
	_, err := ts.VortexApiClient.BasketMargin(ctx, &request)
	checkError429(t, err)
}
