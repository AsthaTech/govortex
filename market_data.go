package govortex

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

// Quotes retrieves real-time quote information for the specified instruments from the Vortex API.
// It takes a context, a slice of instrument names, and a quote mode as input.
// It returns a QuoteResponse and an error.
func (v *VortexApi) Quotes(ctx context.Context, instruments []string, mode QuoteModes) (QuoteResponse, error) {
	endpoint := "/data/quote"
	params := url.Values{}
	for i := 0; i < len(instruments); i++ {
		params.Add("q", instruments[i])
	}
	params.Add("mode", string(mode))
	var resp QuoteResponse
	_, err := v.doJson(ctx, "GET", endpoint, nil, params, nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// HistoricalCandles retrieves historical candlestick data from the Vortex API.
// It takes a context, an ExchangeTypes value, a token, a start time, an end time, and a resolution as input.
// It returns a HistoricalResponse and an error.
func (v *VortexApi) HistoricalCandles(ctx context.Context, exchange ExchangeTypes, token int, from time.Time, to time.Time, resolution Resolutions) (HistoricalResponse, error) {
	params := url.Values{}
	params.Add("exchange", string(exchange))
	params.Add("token", strconv.Itoa(token))
	params.Add("from", strconv.Itoa(int(from.Unix())))
	params.Add("to", strconv.Itoa(int(to.Unix())))
	params.Add("resolution", string(resolution))
	var resp HistoricalResponse
	_, err := v.doJson(ctx, "GET", URIHistory, nil, params, nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
