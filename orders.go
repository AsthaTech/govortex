package govortex

import (
	"context"
	"fmt"
	"net/url"
)

// PlaceOrder places an order with the Vortex API.
// It takes a context and a PlaceOrderRequest as input.
// The request's Validity field is used to determine the ValidityDays and IsAMO values.
// It returns an OrderResponse and an error.
func (v *VortexApi) PlaceOrder(ctx context.Context, request PlaceOrderRequest) (*OrderResponse, error) {
	switch request.Validity {
	case ValidityTypesFullDay:
		request.ValidityDays = 1
		request.IsAMO = false
	case ValidityTypesImmediateOrCancel:
		request.ValidityDays = 0
		request.IsAMO = false
	default:
		request.ValidityDays = 1
		request.IsAMO = true
	}

	var resp OrderResponse
	_, err := v.doJson(ctx, "POST", fmt.Sprintf(URIPlaceOrder, "regular"), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ModifyOrder modifies an existing order with the Vortex API.
// It takes a context, a ModifyOrderRequest, an ExchangeTypes value, and an order ID as input.
// The request's Validity field is used to determine the ValidityDays value.
// It returns an OrderResponse and an error.
func (v *VortexApi) ModifyOrder(ctx context.Context, request ModifyOrderRequest, exchange ExchangeTypes, orderID string) (*OrderResponse, error) {
	// Determine validity_days based on validity type
	switch request.Validity {
	case ValidityTypesFullDay:
		request.ValidityDays = 1
	case ValidityTypesImmediateOrCancel:
		request.ValidityDays = 0
	default:
		request.ValidityDays = 1
	}
	var resp OrderResponse
	encodedOrderId := url.QueryEscape(orderID)
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URIModifyOrder, "regular", encodedOrderId), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelOrder cancels an existing order with the Vortex API.
// It takes a context, an ExchangeTypes value, and an order ID as input.
// It returns an OrderResponse and an error.
func (v *VortexApi) CancelOrder(ctx context.Context, orderID string) (*OrderResponse, error) {
	var resp OrderResponse
	encodedOrderId := url.QueryEscape(orderID)
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URIDeleteOrder, "regular", encodedOrderId), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil

}

// Orders retrieves the order book information from the Vortex API.
// It takes a context, an offset, and a limit as input.
// It returns an OrderBookResponse and an error.
func (v *VortexApi) Orders(ctx context.Context) (*OrderBookResponse, error) {
	var resp OrderBookResponse
	_, err := v.doJson(ctx, "GET", URIOrderBook, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) OrderHistory(ctx context.Context, orderId string) (*OrderHistoryResponse, error) {
	var resp OrderHistoryResponse
	encodedOrderId := url.QueryEscape(orderId)
	_, err := v.doJson(ctx, "GET", fmt.Sprintf(URIOrderHistory, encodedOrderId), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) CancelMultipleRegularOrders(ctx context.Context, req MultipleOrderCancelRequest) (*MultipleOrderResponse, error) {
	var resp MultipleOrderResponse
	_, err := v.doJson(ctx, "POST", URIMultiCancelrders, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
