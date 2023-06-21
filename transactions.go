package govortex

import (
	"context"
	"fmt"
	"net/url"
)

// Positions retrieves the positions information from the Vortex API.
// It returns a PositionsResponse and an error.
func (v *VortexApi) Positions(ctx context.Context) (*PositionResponse, error) {
	var resp PositionResponse
	_, err := v.doJson(ctx, "GET", URIPositions, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) ConvertPosition(ctx context.Context, req ConvertPositionObject) (*ConvertPositionResponse, error) {
	var resp ConvertPositionResponse
	_, err := v.doJson(ctx, "PUT", URIConvertposition, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Holdings retrieves the holdings information from the Vortex API.
// It returns a HoldingsResponse and an error.
func (v *VortexApi) Holdings(ctx context.Context) (*HoldingsResponse, error) {
	var resp HoldingsResponse
	_, err := v.doJson(ctx, "GET", URIHoldings, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Trades retrieves the trade book information from the Vortex API.
// It returns a TradeBookResponse and an error.
func (v *VortexApi) Trades(ctx context.Context, offset int, limit int) (*TradeBookResponse, error) {
	var resp TradeBookResponse
	params := url.Values{}
	params.Add("offset", fmt.Sprintf("%d", offset))
	params.Add("limit", fmt.Sprintf("%d", limit))
	_, err := v.doJson(ctx, "GET", URITrades, nil, params, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
