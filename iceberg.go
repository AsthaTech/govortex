package govortex

import (
	"context"
	"fmt"
)

// PlaceIcebergOrder places an iceberg order with the Vortex API.
// It takes a context and a PlaceIcebergOrderRequest as input.
// It returns an IcebergOrderResponse and an error.
func (v *VortexApi) PlaceIcebergOrder(ctx context.Context, request PlaceIcebergOrderRequest) (*IcebergOrderResponse, error) {
	var resp IcebergOrderResponse
	_, err := v.doJson(ctx, "POST", fmt.Sprintf(URIPlaceOrder, "iceberg"), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ModifyIcebergOrder modifies an existing iceberg order with the Vortex API.
// It takes a context, an iceberg order ID, and a ModifyIcebergOrderRequest as input.
// It returns an IcebergOrderResponse and an error.
func (v *VortexApi) ModifyIcebergOrder(ctx context.Context, iceberg_order_id string, request ModifyIcebergOrderRequest) (*IcebergOrderResponse, error) {
	var resp IcebergOrderResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URIModifyOrder, "iceberg", iceberg_order_id), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelIcebergOrder cancels an existing iceberg order with the Vortex API.
// It takes a context and an iceberg order ID as input.
// It returns a CancelIcebergOrderResponse and an error.
func (v *VortexApi) CancelIcebergOrder(ctx context.Context, iceberg_order_id string) (*CancelIcebergOrderResponse, error) {
	var resp CancelIcebergOrderResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URIDeleteOrder, "iceberg", iceberg_order_id), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
