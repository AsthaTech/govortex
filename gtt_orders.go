package govortex

import (
	"context"
	"fmt"
)

func (v *VortexApi) PlaceGttOrder(ctx context.Context, request PlaceGttRequest) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := v.doJson(ctx, "POST", fmt.Sprintf(URIPlaceOrder, "gtt"), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) ModifyGttOrder(ctx context.Context, gtt_order_id string, request ModifyGttRequest) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URIModifyOrder, "gtt", gtt_order_id), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) CancelGttOrder(ctx context.Context, gtt_order_id string) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URIModifyOrder, "gtt", gtt_order_id), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) GttOrders(ctx context.Context) (*GttOrderbookResponse, error) {
	var resp GttOrderbookResponse
	_, err := v.doJson(ctx, "GET", URIGttOrderBook, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
