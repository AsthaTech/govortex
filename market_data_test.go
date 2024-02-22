package govortex

import (
	"context"
	"testing"
	"time"
)

func (ts *TestSuite) TestQuotes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.Quotes(ctx, []string{"NSE_EQ-22"}, QuoteModesFULL)
	if err != nil {
		t.Errorf("Error while fetching quotes. %v", err)
		return
	}
	if len(resp.Data) == 0 {
		t.Errorf("Errorwhile fetching quotes. %s", "quotes are empty")
	}
}

func (ts *TestSuite) TestHistoricalCandles(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.HistoricalCandles(ctx, ExchangeTypesNSEEQUITY, 22, time.Unix(1494505756, 0), time.Unix(1494505756, 0), ResolutionsMin1)
	if err != nil {
		t.Errorf("Error while fetching historical candles. %v", err)
		return
	}
	if resp.S != "ok" {
		t.Errorf("Errorwhile fetching historical candles. %s", "historical candles are empty")
	}
}

func (ts *ErrorTestSuite) TestQuotes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Quotes(ctx, []string{"NSE_EQ-22"}, QuoteModesFULL)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestHistoricalCandles(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.HistoricalCandles(ctx, ExchangeTypesNSEEQUITY, 22, time.Unix(1494505756, 0), time.Unix(1494505756, 0), ResolutionsMin1)
	checkError429(t, err)
}
