package govortex

type FullQuoteData struct {
	Exchange          string      `json:"exchange"`
	Token             int         `json:"token"`
	LastTradeTime     int         `json:"last_trade_time"`
	LastUpdateTime    int         `json:"last_update_time"`
	LastTradePrice    float64     `json:"last_trade_price"`
	Volume            int         `json:"volume"`
	AverageTradePrice float64     `json:"average_trade_price"`
	TotalBuyQuantity  int64       `json:"total_buy_quantity"`
	TotalSellQuantity int64       `json:"total_sell_quantity"`
	OpenInterest      int         `json:"open_interest"`
	OpenPrice         float64     `json:"open_price"`
	HighPrice         float64     `json:"high_price"`
	LowPrice          float64     `json:"low_price"`
	ClosePrice        float64     `json:"close_price"`
	LastTradeQuantity int         `json:"last_trade_quantity"`
	Depth             *QuoteDepth `json:"depth"`
	DPRHigh           float64     `json:"dpr_high"`
	DPRLow            float64     `json:"dpr_low"`
}

type OhlcvQuoteData struct {
	Exchange       string  `json:"exchange"`
	Token          int     `json:"token"`
	LastTradePrice float64 `json:"last_trade_price"`
	LastTradeTime  int     `json:"last_trade_time"`
	Volume         int     `json:"volume"`
	OpenPrice      float64 `json:"open_price"`
	HighPrice      float64 `json:"high_price"`
	LowPrice       float64 `json:"low_price"`
	ClosePrice     float64 `json:"close_price"`
}

type LtpQuoteData struct {
	Exchange       string  `json:"exchange"`
	Token          int     `json:"token"`
	LastTradePrice float64 `json:"last_trade_price"`
}

type QuoteResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type MarginResponse struct {
	Status    string  `json:"status"`
	Required  float64 `json:"required"`
	Available float64 `json:"available"`
}

type FundsResponse struct {
	NSE FundDetails `json:"nse"`
	MCX FundDetails `json:"mcx"`
}

type HistoricalResponse struct {
	S string    `json:"s"`
	T []int     `json:"t"`
	O []float64 `json:"o"`
	H []float64 `json:"h"`
	L []float64 `json:"l"`
	C []float64 `json:"c"`
	V []float64 `json:"v"`
}

type FundDetails struct {
	Deposit             float64 `json:"deposit"`
	FundsTransferred    float64 `json:"funds_transferred"`
	Collateral          float64 `json:"collateral"`
	CreditForSale       float64 `json:"credit_for_sale"`
	OptionCreditForSale float64 `json:"option_credit_for_sale"`
	LimitUtilization    float64 `json:"limit_utilization"`
	FundsWithdrawn      float64 `json:"funds_withdrawn"`
	MtmAndBookedLoss    float64 `json:"mtm_and_booked_loss"`
	BookedProfit        float64 `json:"booked_profit"`
	TotalTradingPower   float64 `json:"total_trading_power"`
	TotalUtilization    float64 `json:"total_utilization"`
	NetAvailable        float64 `json:"net_available"`
}

type Holding struct {
	ISIN               string          `json:"isin"`
	NSE                ExchangeDetails `json:"nse"`
	BSE                ExchangeDetails `json:"bse"`
	TotalFree          int             `json:"total_free"`
	DPFree             int             `json:"dp_free"`
	PoolFree           int             `json:"pool_free"`
	T1Quantity         int             `json:"t1_quantity"`
	AveragePrice       float64         `json:"average_price"`
	CollateralQuantity int             `json:"collateral_quantity"`
	CollateralValue    float64         `json:"collateral_value"`
}

type ExchangeDetails struct {
	Token    int    `json:"token"`
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
}

type HoldingsResponse struct {
	Status string    `json:"status"`
	Data   []Holding `json:"data"`
}

type PositionResponse struct {
	Status string       `json:"status"`
	Data   PositionData `json:"data"`
}

type PositionData struct {
	Net []Position `json:"net"`
	Day []Position `json:"day"`
}

type Position struct {
	Exchange              string  `json:"exchange"`
	Symbol                string  `json:"symbol"`
	ExpiryDate            string  `json:"expiry_date"`
	OptionType            string  `json:"option_type"`
	Token                 int     `json:"token"`
	Product               string  `json:"product"`
	Quantity              int     `json:"quantity"`
	OvernightBuyValue     float64 `json:"overnight_buy_value"`
	OvernightSellValue    float64 `json:"overnight_sell_value"`
	OvernightAveragePrice float64 `json:"overnight_average_price"`
	LotSize               int     `json:"lot_size"`
	Multiplier            float64 `json:"multiplier"`
	AveragePrice          float64 `json:"average_price"`
	Value                 float64 `json:"value"`
	BuyValue              float64 `json:"buy_value"`
	SellValue             float64 `json:"sell_value"`
	BuyQuantity           int     `json:"buy_quantity"`
	SellQuantity          int     `json:"sell_quantity"`
	BuyPrice              float64 `json:"buy_price"`
	SellPrice             float64 `json:"sell_price"`
}

type TradeBookResponse struct {
	Status   string  `json:"status"`
	Trades   []Trade `json:"trades"`
	Metadata struct {
		TotalRecords int `json:"total_records"`
	} `json:"metadata"`
}

type Resolutions string

const (
	ResolutionsMin1   Resolutions = "1"
	ResolutionsMin2   Resolutions = "2"
	ResolutionsMin3   Resolutions = "3"
	ResolutionsMin4   Resolutions = "4"
	ResolutionsMin5   Resolutions = "5"
	ResolutionsMin10  Resolutions = "10"
	ResolutionsMin15  Resolutions = "15"
	ResolutionsMin30  Resolutions = "30"
	ResolutionsMin45  Resolutions = "45"
	ResolutionsMin60  Resolutions = "60"
	ResolutionsMin120 Resolutions = "120"
	ResolutionsMin180 Resolutions = "180"
	ResolutionsMin240 Resolutions = "240"
	ResolutionsDay    Resolutions = "1D"
	ResolutionsWeek   Resolutions = "1W"
	ResolutionsMonth  Resolutions = "1M"
)

type MarginModes string

const (
	MarginModeNew    MarginModes = "NEW"
	MarginModeMODIFY MarginModes = "MODIFY"
)

type QuoteModes string

const (
	QuoteModesLTP   QuoteModes = "ltp"
	QuoteModesFULL  QuoteModes = "full"
	QuoteModesOHLCV QuoteModes = "ohlcv"
)

type ExchangeTypes string

const (
	ExchangeTypesNSEFO       ExchangeTypes = "NSE_FO"
	ExchangeTypesNSEEQUITY   ExchangeTypes = "NSE_EQ"
	ExchangeTypesNSECURRENCY ExchangeTypes = "NSE_CD"
	ExchangeTypesMCX         ExchangeTypes = "MCX_FO"
)

type TransactionTypes string

const (
	TransactionTypesBUY  TransactionTypes = "BUY"
	TransactionTypesSELL TransactionTypes = "SELL"
)

type ProductTypes string

const (
	ProductTypesIntraday ProductTypes = "INTRADAY"
	ProductTypesDelivery ProductTypes = "DELIVERY"
	ProductTypesMTF      ProductTypes = "MTF"
)

type VarietyTypes string

const (
	VarietyTypesRegularLimitOrder  VarietyTypes = "RL"
	VarietyTypesRegularMarketOrder VarietyTypes = "RL-MKT"
	VarietyTypesStopLimitOrder     VarietyTypes = "SL"
	VarietyTypesStopMarketOrder    VarietyTypes = "SL-MKT"
)

type ValidityTypes string

const (
	ValidityTypesFullDay           ValidityTypes = "DAY"
	ValidityTypesImmediateOrCancel ValidityTypes = "IOC"
	ValidityTypesAfterMarket       ValidityTypes = "AMO"
)

type Order struct {
	OrderID                    string  `json:"order_id"`
	Exchange                   string  `json:"exchange"`
	Token                      int     `json:"token"`
	OrderNumber                string  `json:"order_number"`
	Status                     string  `json:"status"`
	ErrorReason                string  `json:"error_reason"`
	TransactionType            string  `json:"transaction_type"`
	Product                    string  `json:"product"`
	Variety                    string  `json:"variety"`
	TotalQuantity              int     `json:"total_quantity"`
	PendingQuantity            int     `json:"pending_quantity"`
	TradedQuantity             int     `json:"traded_quantity"`
	DisclosedQuantity          int     `json:"disclosed_quantity"`
	DisclosedQuantityRemaining int     `json:"disclosed_quantity_remaining"`
	OrderPrice                 float64 `json:"order_price"`
	TriggerPrice               float64 `json:"trigger_price"`
	TradedPrice                float64 `json:"traded_price"`
	Validity                   string  `json:"validity"`
	ValidityDays               int     `json:"validity_days"`
	Symbol                     string  `json:"symbol"`
	Series                     string  `json:"series"`
	InstrumentName             string  `json:"instrument_name"`
	ExpiryDate                 string  `json:"expiry_date"`
	StrikePrice                float64 `json:"strike_price"`
	OptionType                 string  `json:"option_type"`
	LotSize                    int     `json:"lot_size"`
	OrderCreatedAt             string  `json:"order_created_at"`
	InitiatedBy                string  `json:"initiated_by"`
	ModifiedBy                 string  `json:"modified_by"`
	IsAMO                      bool    `json:"is_amo"`
	OrderIdentifier            string  `json:"order_identifier"`
}

type OrderBookResponse struct {
	Status   string  `json:"status"`
	Orders   []Order `json:"orders"`
	Metadata struct {
		TotalRecords     int `json:"total_records"`
		AllRecords       int `json:"all_records"`
		CompletedRecords int `json:"completed_records"`
		OpenRecords      int `json:"open_records"`
	} `json:"metadata"`
}

type LoginResponse struct {
	Status string   `json:"status"`
	Data   UserData `json:"data"`
}

type UserData struct {
	AccessToken  string   `json:"access_token"`
	UserName     string   `json:"user_name"`
	LoginTime    string   `json:"login_time"`
	Email        string   `json:"email"`
	Mobile       string   `json:"mobile"`
	Exchanges    []string `json:"exchanges"`
	ProductTypes []string `json:"product_types"`
	Others       struct {
		UserCode string `json:"userCode"`
		POA      int    `json:"POA"`
	} `json:"others"`
	UserID        string `json:"user_id"`
	TradingActive bool   `json:"tradingActive"`
}

type OrderResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderID string `json:"orderId"`
	} `json:"data"`
}

type Trade struct {
	OrderID         string  `json:"order_id"`
	Exchange        string  `json:"exchange"`
	Token           int     `json:"token"`
	TradeNo         string  `json:"trade_no"`
	ExchangeOrderNo string  `json:"exchange_order_no"`
	TransactionType string  `json:"transaction_type"`
	Product         string  `json:"product"`
	Variety         string  `json:"variety"`
	TradeQuantity   int     `json:"trade_quantity"`
	TradePrice      float64 `json:"trade_price"`
	Symbol          string  `json:"symbol"`
	Series          string  `json:"series"`
	InstrumentName  string  `json:"instrument_name"`
	ExpiryDate      string  `json:"expiry_date"`
	StrikePrice     float64 `json:"strike_price"`
	OptionType      string  `json:"option_type"`
	TradedAt        string  `json:"traded_at"`
	InitiatedBy     string  `json:"initiated_by"`
	ModifiedBy      string  `json:"modified_by"`
	OrderIdentifier string  `json:"order_identifier"`
}

const (
	Min1   Resolutions = "1"
	Min2   Resolutions = "2"
	Min3   Resolutions = "3"
	Min4   Resolutions = "4"
	Min5   Resolutions = "5"
	Min10  Resolutions = "10"
	Min15  Resolutions = "15"
	Min30  Resolutions = "30"
	Min45  Resolutions = "45"
	Min60  Resolutions = "60"
	Min120 Resolutions = "120"
	Min180 Resolutions = "180"
	Min240 Resolutions = "240"
	Day1   Resolutions = "day"
	Week1  Resolutions = "week"
	Month1 Resolutions = "month"
)

type QuoteDepth struct {
	Buy  []QuoteEntry `json:"buy"`
	Sell []QuoteEntry `json:"sell"`
}

type QuoteEntry struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Orders   int     `json:"orders"`
}

type MarketDepthResponse struct {
	Status string     `json:"status"`
	Data   QuoteDepth `json:"data"`
}
