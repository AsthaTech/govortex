package govortex

import (
	"context"
	"fmt"
)

func (v *VortexApi) PlaceIcebergOrder(ctx context.Context, request PlaceIcebergOrderRequest) (*PlaceIcebergOrderResponse, error) {
	var resp PlaceIcebergOrderResponse
	_, err := v.doJson(ctx, "POST", fmt.Sprintf(URIPlaceOrder, "iceberg"), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) ModifyIcebergOrder(ctx context.Context, request ModifyIcebergOrderRequest) (*PlaceIcebergOrderResponse, error) {
	var resp PlaceIcebergOrderResponse
	_, err := v.doJson(ctx, "POST", fmt.Sprintf(URIPlaceOrder, "iceberg"), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
