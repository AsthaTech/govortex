package govortex

import (
	"context"
)

// GetStrategies retrieves strategies from the Vortex API based on the provided StrategiesRequest.
// It takes a context and a StrategiesRequest as input.
// It returns a StrategiesResponse and an error.
func (v *VortexApi) GetStrategies(ctx context.Context, req StrategiesRequest) (*StrategiesResponse, error) {
	var resp StrategiesResponse
	_, err := v.doJson(ctx, "POST", URIStrategies, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOptionChain retrieves the option chain from the Vortex API based on the provided OptionChainRequest.
// It takes a context and an OptionChainRequest as input.
// It returns an OptionChainResponse and an error.
func (v *VortexApi) GetOptionChain(ctx context.Context, req OptionChainRequest) (*OptionChainResponse, error) {
	var resp OptionChainResponse
	_, err := v.doJson(ctx, "POST", URIOptionChain, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BuildStrategy initiates the strategy building process with the Vortex API based on the provided StrategyBuilderRequest.
// It takes a context and a StrategyBuilderRequest as input.
// It returns a StrategiesResponse and an error.
func (v *VortexApi) BuildStrategy(ctx context.Context, req StrategyBuilderRequest) (*StrategiesResponse, error) {
	var resp StrategiesResponse
	_, err := v.doJson(ctx, "POST", URIBuildStrategies, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPayoff calculates the payoff for strategies with the Vortex API based on the provided PayoffRequest.
// It takes a context and a PayoffRequest as input.
// It returns a PayoffResponse and an error.
func (v *VortexApi) GetPayoff(ctx context.Context, req PayoffRequest) (*PayoffResponse, error) {
	var resp PayoffResponse
	_, err := v.doJson(ctx, "POST", URIPayoffStrategies, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
