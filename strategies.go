package govortex

import (
	"context"
)

func (v *VortexApi) GetStrategies(ctx context.Context, req StrategiesRequest) (*StrategiesResponse, error) {
	var resp StrategiesResponse
	_, err := v.doJson(ctx, "POST", URIStrategies, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) GetOptionChain(ctx context.Context, req OptionChainRequest) (*OptionChainResponse, error) {
	var resp OptionChainResponse
	_, err := v.doJson(ctx, "POST", URIOptionChain, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) BuildStrategy(ctx context.Context, req StrategyBuilderRequest) (*StrategiesResponse, error) {
	var resp StrategiesResponse
	_, err := v.doJson(ctx, "POST", URIBuildStrategies, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) GetPayoff(ctx context.Context, req PayoffRequest) (*PayoffResponse, error) {
	var resp PayoffResponse
	_, err := v.doJson(ctx, "POST", URIPayoffStrategies, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
