package govortex

import (
	"context"
	"net/http"
)

// Funds retrieves the funds information from the Vortex API.
// It returns a FundsResponse containing the funds information and an error if any.
func (v *VortexApi) Funds(ctx context.Context) (*FundsResponse, error) {
	var resp FundsResponse
	_, err := v.doJson(ctx, http.MethodGet, URIFunds, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Banks retrieves a list of banks from the Vortex API.
// It returns a BanksResponse containing the list of banks and an error if any.
func (v *VortexApi) Banks(ctx context.Context) (*BanksResponse, error) {
	var resp BanksResponse
	_, err := v.doJson(ctx, "GET", URIBanks, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// func (v *VortexApi) Brokerage(ctx context.Context) (*BrokerageResponse, error) {
// 	var resp BrokerageResponse
// 	_, err := v.doJson(ctx, "GET", URIBrokerage, nil, nil, nil, &resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &resp, nil
// }
