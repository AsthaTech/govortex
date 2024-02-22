package govortex

import (
	"context"
)

// OrderMargin gets the required margin to place or modify an order in the Vortex API.
// It takes a context and an OrderMarginRequest as input.
// It returns a MarginResponse and an error.
func (v *VortexApi) OrderMargin(ctx context.Context, request *OrderMarginRequest) (*MarginResponse, error) {
	var resp MarginResponse
	_, err := v.doJson(ctx, "POST", URIOrderMargin, request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BasketMargin gets the required margin to place a collection of new orders in the Vortex API.
// It takes a context and an BasketMarginRequest as input.
// It returns a BasketMarginResponse and an error.
func (v *VortexApi) BasketMargin(ctx context.Context, request *BasketMarginRequest) (*BasketMarginResponse, error) {
	var resp BasketMarginResponse
	_, err := v.doJson(ctx, "POST", URIBasketMargin, request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
