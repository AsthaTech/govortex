package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestPlaceIcebergOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.PlaceIcebergOrder(ctx, PlaceIcebergOrderRequest{})
	if err != nil {
		t.Errorf("Error while creating order. %v", err)
		return
	}
	if resp.Data.IcebergOrderId != "5eaefd25-518c-4a39-b556-93fc8e78e855" {
		t.Errorf("Error while creating order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestModifyIcebergOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.ModifyIcebergOrder(ctx, "5eaefd25-518c-4a39-b556-93fc8e78e855", ModifyIcebergOrderRequest{})
	if err != nil {
		t.Errorf("Error while modifying order. %v", err)
		return
	}
	if resp.Data.IcebergOrderId == "5eaefd25-518c-4a39-b556-93fc8e78e855" {
		t.Errorf("Error while modifying order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestCancelIcebergOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.CancelIcebergOrder(ctx, "5eaefd25-518c-4a39-b556-93fc8e78e855")
	if err != nil {
		t.Errorf("Error while cancelling order. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while cancelling order. %s", "status is not success")
	}
}
