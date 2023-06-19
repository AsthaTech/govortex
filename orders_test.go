package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestPlaceOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := PlaceOrderRequest{}
	resp, err := ts.VortexApiClient.PlaceOrder(ctx, request)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
		return
	}
	if resp.Data.OrderID != "NXAAE00002K3" {
		t.Errorf("Error while placing order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestModifyOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := ModifyOrderRequest{}
	resp, err := ts.VortexApiClient.ModifyOrder(ctx, request, ExchangeTypesNSEEQUITY, "NXAAE00002K3")
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
		return
	}
	if resp.Data.OrderID != "NXAAE00002K3" {
		t.Errorf("Error while modifying order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestCancelOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.CancelOrder(ctx, ExchangeTypesNSEEQUITY, "NXAAE00002K3")
	if err != nil {
		t.Errorf("Error while cancelling order. %v", err)
		return
	}
	if resp.Data.OrderID != "NXAAE00002K3" {
		t.Errorf("Error while cancelling order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestOrderBook(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.Orders(ctx, 1, 20)
	if err != nil {
		t.Errorf("Error while fetching order book. %v", err)
		return
	}
	if len(resp.Orders) == 0 {
		t.Errorf("Errorwhile fetching order book. %s", "order book is empty")
	}
}
