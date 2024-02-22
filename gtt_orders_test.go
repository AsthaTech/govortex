package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestPlaceGttOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := PlaceGttRequest{}
	resp, err := ts.VortexApiClient.PlaceGttOrder(ctx, request)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
		return
	}
	if resp.Data.OrderID != "99823d7b-fd37-4d75-af7f-f21ec4671852" {
		t.Errorf("Error while placing order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestGetGttOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.GttOrders(ctx)
	if err != nil {
		t.Errorf("Error while fetching order book. %v", err)
		return
	}
	if len(resp.Data) == 0 {
		t.Errorf("Errorwhile fetching order book. %s", "order book is empty")
	}
}

func (ts *TestSuite) TestModifyGttOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.ModifyGttOrder(ctx, "99823d7b-fd37-4d75-af7f-f21ec4671852", ModifyGttRequest{})
	if err != nil {
		t.Errorf("Error while modifying order. %v", err)
		return
	}
	if resp.Data.OrderID == "99823d7b-fd37-4d75-af7f-f21ec4671852" {
		t.Errorf("Error while modifying order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestDeleteGttOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.CancelGttOrder(ctx, "99823d7b-fd37-4d75-af7f-f21ec4671852")
	if err != nil {
		t.Errorf("Error while cancelling order. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while modifying order. %s", "status is not success")
	}
}
