package govortex

import (
	"context"
	"fmt"
)

// PlaceGttOrder places a Good Till Trigger (GTT) order with the Vortex API.
// It takes a context and a PlaceGttRequest as input.
// It returns an OrderResponse and an error.
func (v *VortexApi) PlaceGttOrder(ctx context.Context, request PlaceGttRequest) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := v.doJson(ctx, "POST", fmt.Sprintf(URIPlaceOrder, "gtt"), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ModifyGttOrder modifies an existing Good Till Trigger (GTT) order with the Vortex API.
// It takes a context, a GTT order ID, and a ModifyGttRequest as input.
// It returns an OrderResponse and an error.
func (v *VortexApi) ModifyGttOrder(ctx context.Context, gtt_order_id string, request ModifyGttRequest) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URIModifyOrder, "gtt", gtt_order_id), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelGttOrder cancels an existing Good Till Trigger (GTT) order with the Vortex API.
// It takes a context and a GTT order ID as input.
// It returns an OrderResponse and an error.
func (v *VortexApi) CancelGttOrder(ctx context.Context, gtt_order_id string) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URIDeleteOrder, "gtt", gtt_order_id), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GttOrders retrieves the Good Till Trigger (GTT) orderbook from the Vortex API.
// It takes a context as input.
// It returns a GttOrderbookResponse and an error.
func (v *VortexApi) GttOrders(ctx context.Context) (*GttOrderbookResponse, error) {
	var resp GttOrderbookResponse
	_, err := v.doJson(ctx, "GET", URIGttOrderBook, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
