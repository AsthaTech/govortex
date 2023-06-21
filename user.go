package govortex

import (
	"context"
)

// Funds retrieves the funds information from the Vortex API.
func (v *VortexApi) Funds(ctx context.Context) (*FundsResponse, error) {
	var resp FundsResponse
	_, err := v.doJson(ctx, "GET", URIFunds, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
