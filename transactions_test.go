package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestGetPositions(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	positions, err := ts.VortexApiClient.Positions(ctx)
	if err != nil {
		t.Errorf("Error while fetching positions. %v", err)
	}
	if len(positions.Data.Day) == 0 {
		t.Errorf("Error while fetching day positions. %v", err)
	}
}

func (ts *TestSuite) TestConvertPosition(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.ConvertPosition(ctx, ConvertPositionRequest{
		Exchange:        ExchangeTypesNSEEQUITY,
		Token:           22,
		TransactionType: TransactionTypesBUY,
		OldProductType:  ProductTypesIntraday,
		NewProductType:  ProductTypesDelivery,
	})
	if err != nil {
		t.Errorf("Error while converting position. %v", err)
	}
}

func (ts *TestSuite) TestGetHoldings(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	holdings, err := ts.VortexApiClient.Holdings(ctx)
	if err != nil {
		t.Errorf("Error while fetching holdings. %v", err)
		return
	}
	if len(holdings.Data) == 0 {
		t.Errorf("Error while fetching holdings. %v", err)
	}
}

func (ts *TestSuite) TestTrades(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	trades, err := ts.VortexApiClient.Trades(ctx, 1, 20)
	if err != nil {
		t.Errorf("Error while fetching trades. %v", err)
		return
	}
	if len(trades.Trades) == 0 {
		t.Errorf("Error while fetching trades. %v", "empty trade book")
	}
}

func (ts *ErrorTestSuite) TestGetPositions(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Positions(ctx)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestConvertPosition(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.ConvertPosition(ctx, ConvertPositionRequest{
		Exchange:        ExchangeTypesNSEEQUITY,
		Token:           22,
		TransactionType: TransactionTypesBUY,
		OldProductType:  ProductTypesIntraday,
		NewProductType:  ProductTypesDelivery,
	})
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestGetHoldings(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Holdings(ctx)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestTrades(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Trades(ctx, 1, 20)
	checkError429(t, err)
}
