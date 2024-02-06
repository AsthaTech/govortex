package govortex

import (
	"time"

	"github.com/lib/pq"
)

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

type ExchangeDetails struct {
	Token    int    `json:"token"`
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
}

type ExchangeDetail struct {
	Token    int           `json:"token"`
	Exchange ExchangeTypes `json:"exchange"`
	Symbol   string        `json:"symbol"`
}

type Holding struct {
	ISIN               string          `json:"isin"`
	NSE                *ExchangeDetail `json:"nse,omitempty"`
	BSE                *ExchangeDetail `json:"bse,omitempty"`
	TotalFree          int             `json:"total_free"`
	DPFree             int             `json:"dp_free"`
	PoolFree           int             `json:"pool_free"`
	T1Quantity         int             `json:"t1_quantity"`
	AveragePrice       float64         `json:"average_price"`
	LastPrice          float64         `json:"last_price"`
	Product            string          `json:"product"`
	CollateralQuantity int             `json:"collateral_quantity"`
	CollateralValue    float64         `json:"collateral_value"`
}

type HoldingsResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message,omitempty"`
	Data    []Holding `json:"data"`
}

type PositionResponse struct {
	Status string          `json:"status"`
	Data   NetDayPositions `json:"data"`
}

type NetDayPositions struct {
	Net []PositionItem `json:"net"`
	Day []PositionItem `json:"day"`
}

type PositionItem struct {
	Exchange              ExchangeTypes `json:"exchange"`
	Symbol                string        `json:"symbol"`
	ExpiryDate            string        `json:"expiry_date"`
	OptionType            OptionType    `json:"option_type"`
	StrikePrice           float64       `json:"strike_price"`
	Token                 int           `json:"token"`
	Product               ProductTypes  `json:"product"`
	Quantity              int           `json:"quantity"`
	OvernightBuyValue     float64       `json:"overnight_buy_value"`
	OvernightSellValue    float64       `json:"overnight_sell_value"`
	OvernightAveragePrice float64       `json:"overnight_average_price"`
	LotSize               int           `json:"lot_size"`
	Multiplier            float64       `json:"multiplier"`
	AveragePrice          float64       `json:"average_price"`
	Value                 float64       `json:"value"`
	BuyValue              float64       `json:"buy_value"`
	SellValue             float64       `json:"sell_value"`
	BuyQuantity           int           `json:"buy_quantity"`
	SellQuantity          int           `json:"sell_quantity"`
	BuyPrice              float64       `json:"buy_price"`
	SellPrice             float64       `json:"sell_price"`
}

type ConvertPositionResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
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

type OrderBookResponse struct {
	Status   string   `json:"status"`
	Orders   []Order  `json:"orders"`
	Metadata Metadata `json:"metadata"`
}

type Order struct {
	OrderID                    string                `json:"order_id"`
	Exchange                   ExchangeTypes         `json:"exchange"`
	Token                      int                   `json:"token"`
	OrderNumber                string                `json:"order_number"`
	Status                     string                `json:"status"`
	ErrorReason                string                `json:"error_reason"`
	TransactionType            TransactionTypes      `json:"transaction_type"`
	Product                    ProductTypes          `json:"product"`
	Variety                    VarietyTypes          `json:"variety"`
	TotalQuantity              int                   `json:"total_quantity"`
	PendingQuantity            int                   `json:"pending_quantity"`
	TradedQuantity             int                   `json:"traded_quantity"`
	DisclosedQuantity          int                   `json:"disclosed_quantity"`
	DisclosedQuantityRemaining int                   `json:"disclosed_quantity_remaining"`
	OrderPrice                 float64               `json:"order_price"`
	TriggerPrice               float64               `json:"trigger_price"`
	TradedPrice                float64               `json:"traded_price,omitempty"`
	Validity                   ValidityTypes         `json:"validity"`
	Symbol                     string                `json:"symbol"`
	Series                     string                `json:"series"`
	InstrumentName             InstrumentName        `json:"instrument_name"`
	ExpiryDate                 string                `json:"expiry_date"`
	StrikePrice                float64               `json:"strike_price"`
	OptionType                 OptionType            `json:"option_type"`
	LotSize                    int                   `json:"lot_size"`
	OrderCreatedAt             string                `json:"order_created_at"`
	InitiatedBy                string                `json:"initiated_by"`
	ModifiedBy                 string                `json:"modified_by"`
	IsAMO                      bool                  `json:"is_amo"`
	OrderIdentifier            string                `json:"order_identifier"`
	TagsIds                    pq.Int32Array         `json:"tags_ids"`
	MiddlewareOrderId          uint                  `json:"middleware_order_id"`
	Iceberg                    *OrderBookIcebergInfo `json:"iceberg,omitempty"`
	Gtt                        *OrderBookGttInfo     `json:"gtt,omitempty"`
}
type OrderBookIcebergInfo struct {
	IcebergOrderId  string `json:"iceberg_order_id"`
	IcebergSequence int    `json:"iceberg_sequence"`
	Legs            int    `json:"legs"`
}

type OrderBookGttInfo struct {
	TriggerType          GttTriggerType `json:"trigger_type"`
	SlTriggerPercent     *float64       `json:"sl_trigger_percent,omitempty"`
	ProfitTriggerPercent *float64       `json:"profit_trigger_percent,omitempty"`
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
		OrderID string `json:"order_id"`
	} `json:"data"`
}

type Trade struct {
	OrderID         string           `json:"order_id"`
	Exchange        ExchangeTypes    `json:"exchange"`
	Token           int              `json:"token"`
	TradeNo         string           `json:"trade_no"`
	ExchangeOrderNo string           `json:"exchange_order_no"`
	TransactionType TransactionTypes `json:"transaction_type"`
	Product         ProductTypes     `json:"product"`
	Variety         VarietyTypes     `json:"variety"`
	TradeQuantity   int              `json:"trade_quantity"`
	TradePrice      float64          `json:"trade_price"`
	Symbol          string           `json:"symbol"`
	Series          string           `json:"series"`
	InstrumentName  string           `json:"instrument_name"`
	ExpiryDate      string           `json:"expiry_date"`
	StrikePrice     float64          `json:"strike_price"`
	OptionType      OptionType       `json:"option_type"`
	TradedAt        string           `json:"traded_at"`
	InitiatedBy     string           `json:"initiated_by"`
	ModifiedBy      string           `json:"modified_by"`
	OrderIdentifier string           `json:"order_identifier"`
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

type OrderHistoryResponse struct {
	Status   string         `json:"status"`
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	Data     []OrderHistory `json:"orders"`
	Metadata Metadata       `json:"metadata"`
}

type OrderHistory struct {
	OrderID                string           `json:"order_id"`
	Exchange               ExchangeTypes    `json:"exchange"`
	Token                  int              `json:"token"`
	OrderNumber            string           `json:"order_number"`
	Status                 string           `json:"status"`
	ErrorReason            string           `json:"error_reason"`
	TransactionType        TransactionTypes `json:"transaction_type"`
	Product                ProductTypes     `json:"product"`
	Variety                VarietyTypes     `json:"variety"`
	TotalQuantity          int              `json:"total_quantity"`
	PendingQuantity        int              `json:"pending_quantity"`
	TradedQuantity         int              `json:"traded_quantity"`
	DisclosedQuantity      int              `json:"disclosed_quantity"`
	OrderPrice             float64          `json:"order_price"`
	TriggerPrice           float64          `json:"trigger_price"`
	Validity               ValidityTypes    `json:"validity"`
	ValidityDays           int              `json:"validity_days"`
	Symbol                 string           `json:"symbol"`
	Series                 string           `json:"series"`
	InstrumentName         InstrumentName   `json:"instrument_name"`
	ExpiryDate             string           `json:"expiry_date"`
	StrikePrice            float64          `json:"strike_price"`
	OptionType             OptionType       `json:"option_type"`
	OrderCreatedAt         string           `json:"order_created_at"`
	ExchangeOrderCreatedAt string           `json:"exchange_order_created_at"`
	InitiatedBy            string           `json:"initiated_by"`
	ModifiedBy             string           `json:"modified_by"`
	IsAMO                  bool             `json:"is_amo"`
	OrderIdentifier        string           `json:"order_identifier"`
}

type Metadata struct {
	TotalRecords int `json:"total_records"`
}

type InstrumentName string

const (
	InstrumentNameEqIndex             InstrumentName = "EQIDX"
	InstrumentNameCom                 InstrumentName = "COM"
	InstrumentNameEquities            InstrumentName = "EQUITIES"
	InstrumentNameCommodityFuture     InstrumentName = "FUTCOM"
	InstrumentNameCurrencyFuture      InstrumentName = "FUTCUR"
	InstrumentNameIndexFuture         InstrumentName = "FUTIDX"
	InstrumentNameInterestFuture      InstrumentName = "FUTIRC"
	InstrumentNameInterestFutureT     InstrumentName = "FUTIRT"
	InstrumentNameStockFuture         InstrumentName = "FUTSTK"
	InstrumentNameCurrencyOption      InstrumentName = "OPTCUR"
	InstrumentNameCommodityOption     InstrumentName = "OPTFUT"
	InstrumentNameIndexOption         InstrumentName = "OPTIDX"
	InstrumentNameInterestOption      InstrumentName = "OPTIRC"
	InstrumentNameStockOption         InstrumentName = "OPTSTK"
	InstrumentNameCurrentcyUnderlying InstrumentName = "UNDCUR"
)

type OptionType string

const (
	OptionTypeCall OptionType = "CE"
	OptionTypePut  OptionType = "PE"
)

type GttOrderbookResponse struct {
	Status string              `json:"status"`
	Data   []*GttOrderResponse `json:"data"`
}

type GttOrderResponse struct {
	Id              string                   `json:"id"`
	Token           int                      `json:"token" binding:"required"`
	Exchange        ExchangeTypes            `json:"exchange" binding:"required"`
	Symbol          string                   `json:"symbol" binding:"required"`
	Series          string                   `json:"series" binding:"required"`
	InstrumentName  InstrumentName           `json:"instrument_name" binding:"required"`
	ExpiryDate      string                   `json:"expiry_date" binding:"required"`
	StrikePrice     float64                  `json:"strike_price" binding:"required"`
	OptionType      OptionType               `json:"option_type" binding:"required"`
	LotSize         int                      `json:"lot_size" binding:"required"`
	TriggerType     GttTriggerType           `json:"trigger_type" binding:"required"`
	TransactionType TransactionTypes         `json:"transaction_type" binding:"required"`
	TagIds          pq.Int32Array            `json:"tag_ids"`
	Orders          []GttOrderResponseOrders `json:"orders" binding:"required"`
}

type GttOrderResponseOrders struct {
	Id              uint             `json:"id"`
	ProductType     ProductTypes     `json:"product"`
	Variety         VarietyTypes     `json:"variety"`
	TransactionType TransactionTypes `json:"transaction_type"`
	Price           float64          `json:"price"`
	TriggerPrice    float64          `json:"trigger_price"`
	Quantity        int              `json:"quantity"`
	Status          GttOrderStatus   `json:"status"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	TrigerredAt     time.Time        `json:"trigerred_at"`
}

type GttOrderStatus string

const (
	GttOrderStatusTriggered GttOrderStatus = "triggered"
	GttOrderStatusActive    GttOrderStatus = "active"
	GttOrderStatusCancelled GttOrderStatus = "cancelled"
	GttOrderStatusExpired   GttOrderStatus = "expired"
	GttOrderStatusCompleted GttOrderStatus = "completed"
)

type FundResponse struct {
	Nse ExchangeFundResponse `json:"nse"`
	Mcx ExchangeFundResponse `json:"mcx"`
}

type ExchangeFundResponse struct {
	Deposit             float64 `json:"deposit"`
	FundsTransferred    float64 `json:"funds_transferred"`
	Collateral          float64 `json:"collateral"`
	CreditForSale       float64 `json:"credit_for_sale"`
	OptionCreditForSale float64 `json:"option_credit_for_sale"`
	LimitUtilization    float64 `json:"limit_utilization"`
	MtmAndBookedLoss    float64 `json:"mtm_and_booked_loss"`
	BookedProfit        float64 `json:"booked_profit"`
	TotalTradingPower   float64 `json:"total_trading_power"`
	TotalUtilization    float64 `json:"total_utilization"`
	NetAvailable        float64 `json:"net_available"`
	FundsWithdrawal     float64 `json:"funds_withdrawal"`
}

type FundWithdrawalListResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message,omitempty"`
	Data    []FundWithdrawalItem `json:"data"`
}

type FundWithdrawalItem struct {
	TransactionId string        `json:"transaction_id"`
	Amount        float64       `json:"amount"`
	CreatedAt     time.Time     `json:"created_at"`
	Status        string        `json:"status"`
	Exchange      ExchangeTypes `json:"exchange"`
}

type TagResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    TagInfo `json:"data"`
}

type TagsResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []TagInfo `json:"data"`
}
type TagInfo struct {
	Id          int       `json:"id"`
	ClientCode  string    `json:"client_code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type IcebergOrderResponse struct {
	Message string           `json:"message"`
	Code    string           `json:"code"`
	Status  string           `json:"status"`
	Data    IcebergOrderData `json:"data"`
}

type CancelIcebergOrderResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
type IcebergOrderData struct {
	IcebergOrderId string `json:"iceberg_order_id"`
	FirstOrderId   string `json:"first_order_id"`
}
type StrategiesResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    StrategiesResult `json:"data"`
}
type StrategiesResult struct {
	Symbol     string     `json:"symbol"`
	Token      int        `json:"token"`
	Ltp        float64    `json:"ltp"`
	Strategies []Strategy `json:"strategies"`
}

type Strategy struct {
	Name                 string               `json:"strategy_name"`
	TradingOpportunities []TradingOpportunity `json:"trading_opportunities"`
}
type TradingOpportunity struct {
	Legs             []StrategyLeg `json:"legs"`
	MaxLoss          float64       `json:"max_loss"`
	MaxProfit        float64       `json:"max_profit"`
	IsInfiniteProfit bool          `json:"is_infinite_profit"`
	IsInfiniteLoss   bool          `json:"is_infinite_loss"`
	BreakEven        []float64     `json:"breakeven"`
}

type StrategyLeg struct {
	Option   StrategyStock `json:"option"`
	Action   string        `json:"action"`
	Quantity int           `json:"quantity"`
}

type StrategyStock struct {
	Token          int           `json:"token"`
	InstrumentName string        `json:"instrument_name"`
	Symbol         string        `json:"symbol"`
	StrikePrice    float64       `json:"strike_price"`
	OptionType     OptionType    `json:"option_type"`
	LotSize        int           `json:"lot_size"`
	SecurityDesc   string        `json:"security_description"`
	Exchange       ExchangeTypes `json:"exchange"`
	ExpYyyymmdd    string        `json:"expiry_date"`
	Ltp            float64       `json:"ltp"`
	Greeks         struct {
		Theta float64 `json:"theta"`
		Delta float64 `json:"delta"`
		Gamma float64 `json:"gamma"`
		Vega  float64 `json:"vega"`
		Iv    float64 `json:"iv"`
	} `json:"greeks"`
}

type OptionChainResponse struct {
	Status   string            `json:"status"`
	Message  string            `json:"message"`
	Response OptionchainResult `json:"response"`
}
type OptionchainResult struct {
	Symbol         string   `json:"symbol"`
	ExpiryDate     string   `json:"expiry_date"`
	HasParentStock bool     `json:"has_parent_stock"`
	ExpiryDates    []string `json:"expiry_dates"`
	Options        struct {
		Exchange ExchangeTypes     `json:"exchange"`
		List     []*DateOptionData `json:"list"`
	} `json:"options"`
	ParentStock struct {
		Symbol   string        `json:"symbol"`
		Exchange ExchangeTypes `json:"exchange"`
		Token    int           `json:"token"`
		ISINCode string        `json:"isin"`
		LTP      float64       `json:"ltp"`
	} `json:"parent"`
}

type DateOptionData struct {
	StrikePrice float64        `json:"strike_price"`
	IV          float64        `json:"iv"`
	Theta       float64        `json:"theta"`
	Vega        float64        `json:"vega"`
	Gamma       float64        `json:"gamma"`
	CE          StockWithGreek `json:"CE"`
	PE          StockWithGreek `json:"PE"`
}

type StockWithGreek struct {
	Token          int     `json:"token"`
	InstrumentName string  `json:"instrument_name"`
	LotSize        int     `json:"lot_size"`
	SecurityDesc   string  `json:"security_description"`
	Eligibility    int     `json:"eligibility"`
	Ltp            float64 `json:"ltp"`
	OpenInterest   int     `json:"open_interest"`
	DayFirstTickOI int     `json:"day_first_tick_oi" `
	Volume         int     `json:"volume"`
	Delta          float64 `json:"delta"`
}
type PayoffResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    PayOffData `json:"data"`
}

type PayOffData struct {
	MaxLoss          float64         `json:"max_loss"`
	MaxProfit        float64         `json:"max_profit"`
	IsInfiniteProfit bool            `json:"infinite_profit"`
	IsInfiniteLoss   bool            `json:"infinite_loss"`
	BreakEvens       []float64       `json:"breakevens"`
	PayOffs          []PayOff        `json:"payoffs"`
	CombinedGreeks   Greeks          `json:"combined_greeks"`
	LivePrice        float64         `json:"last_trade_price"`
	LegStocks        []LegStockGreek `json:"leg_greeks"`
	MinDaysToExpiry  int             `json:"min_days_to_expiry"`
}

type LegStockGreek struct {
	Token       int        `json:"token" binding:"required"`
	StrikePrice float64    `json:"strike_price"`
	OptionType  OptionType `json:"option_type"`
	ExpYYYYMMDD string     `json:"expiry_date"`
	Action      string     `json:"action"`
	LotSize     int        `json:"lot_size"`
	Quantity    int        `json:"quantity"`
	Ltp         float64    `json:"last_trade_price"` // Update
	Iv          float64    `json:"iv"`               // Update *hide
	Greeks      Greeks     `json:"greeks" gorm:"-"`
}
type Greeks struct {
	Theta float64 `json:"theta"`
	Delta float64 `json:"delta"`
	Gamma float64 `json:"gamma"`
	Vega  float64 `json:"vega"`
	Iv    float64 `json:"iv"`
}

type PayOff struct {
	IPayoff float64 `json:"intraday_pay_off"`
	EPayoff float64 `json:"expiry_pay_off"`
	At      float64 `json:"at"`
}
