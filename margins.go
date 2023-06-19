package govortex

import (
	"context"
)

// OrderMargin places an order for margin trading in the Vortex API.
// It takes a context and an OrderMarginRequest as input.
// It returns a MarginResponse and an error.
func (v *VortexApi) OrderMargin(ctx context.Context, request OrderMarginRequest) (MarginResponse, error) {
	var resp MarginResponse
	_, err := v.doJson(ctx, "POST", URIOrderMargin, request, nil, nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
