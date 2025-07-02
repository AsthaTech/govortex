package govortex

import (
	"github.com/lib/pq"
)

// PlaceOrderRequest represents a request to place an order.
type PlaceOrderRequest struct {
	Exchange          ExchangeTypes    `json:"exchange"`           // Required: Exchange type.
	Token             int              `json:"token"`              // Required: Token of the underlying instrument.
	TransactionType   TransactionTypes `json:"transaction_type"`   // Required: Type of transaction.
	Product           ProductTypes     `json:"product"`            // Required: Type of product.
	Variety           VarietyTypes     `json:"variety"`            // Required: Type of variety.
	Quantity          int              `json:"quantity"`           // Required: Quantity of the order.
	Price             float64          `json:"price"`              // Optional if market order. Price of the order.
	TriggerPrice      float64          `json:"trigger_price"`      // Optional if not stoploss order. Trigger price for the order.
	OrderIdentifier   string           `json:"order_identifier"`   // Optional: Your identifier for the order.
	DisclosedQuantity int              `json:"disclosed_quantity"` // Optional: Disclosed quantity for the order.
	Validity          ValidityTypes    `json:"validity"`           // Required: Validity type for the order.
	ValidityDays      int              `json:"validity_days"`      // Optional: Number of validity days.
	IsAMO             bool             `json:"is_amo"`             // Optional: Flag indicating if the order is an after-market order.
	Gtt               *GttLegs         `json:"gtt"`                // Optional: Good 'til Triggered (GTT) legs.
	TagIds            []int            `json:"tag_ids"`            // Optional: IDs of tags associated with the order.
}

// GttLegs represents legs of a Good 'til Triggered (GTT) order.
type GttLegs struct {
	SlTriggerPercent     *float64       `json:"sl_trigger_percent"`     // Optional: Stop loss trigger percentage.
	ProfitTriggerPercent *float64       `json:"profit_trigger_percent"` // Optional: Profit trigger percentage.
	SlVariety            *VarietyTypes  `json:"sl_variety"`
	ProfitVariety        *VarietyTypes  `json:"profit_variety"`
	TrailJumpPoint       *float64       `json:"trail_jump_point"`
	TrailJumpType        *TrailJumpType `json:"trail_jump_type"`
}

type ModifyOrderRequest struct {
	Variety           VarietyTypes  `json:"variety" `           // Required: Type of variety.
	Quantity          int           `json:"quantity" `          // Required: Quantity of the order.
	TradedQuantity    *int          `json:"traded_quantity" `   // Required: Quantity which has already been traded according to you. This is important..
	Price             float64       `json:"price" `             // Optional if market order. Price of the order.
	TriggerPrice      float64       `json:"trigger_price"`      // Optional if not stoploss order. Trigger price for the order.
	DisclosedQuantity int           `json:"disclosed_quantity"` // Optional: Disclosed quantity for the order.
	Validity          ValidityTypes `json:"validity" `          // Required: Validity type for the order.
	ValidityDays      int           `json:"validity_days"`      // Optional: Number of validity days.
	TagIds            []int         `json:"tag_ids"`            // Optional: IDs of tags associated with the order.
}

// OrderMarginRequest represents a request to calculate margin for an order.
type OrderMarginRequest struct {
	Exchange        ExchangeTypes    `json:"exchange"`         // Required: Exchange type.
	Token           int              `json:"token"`            // Required: Token of the underlying instrument.
	TransactionType TransactionTypes `json:"transaction_type"` // Required: Type of transaction.
	Product         ProductTypes     `json:"product"`          // Required: Type of product.
	Variety         VarietyTypes     `json:"variety"`          // Required: Type of variety.
	Quantity        int              `json:"quantity"`         // Required: Quantity of the order.
	Price           float64          `json:"price"`            // Required: Price of the order.
	OldPrice        float64          `json:"old_price"`        // Required: Old price of the order.
	OldQuantity     int              `json:"old_quantity"`     // Required: Old quantity of the order.
	Mode            MarginModes      `json:"mode"`             // Required: Mode of margin calculation.
}

type BasketMarginOrder struct {
	Exchange        ExchangeTypes    `json:"exchange"`         // Required: Exchange type.
	Token           int              `json:"token"`            // Required: Token of the underlying instrument.
	TransactionType TransactionTypes `json:"transaction_type"` // Required: Type of transaction.
	Product         ProductTypes     `json:"product"`          // Required: Type of product.
	Variety         VarietyTypes     `json:"variety"`          // Required: Type of variety.
	Quantity        int              `json:"quantity"`         // Required: Quantity of the order.
	Price           float64          `json:"price"`            // Required: Price of the order.
}
type BasketMarginRequest struct {
	Orders []BasketMarginOrder `json:"orders"`
}

// ConvertPositionRequest represents a request to convert a position.
type ConvertPositionRequest struct {
	Exchange        ExchangeTypes    `json:"exchange"`         // Required: Exchange type.
	Token           int              `json:"token"`            // Required: Token of the underlying instrument.
	TransactionType TransactionTypes `json:"transaction_type"` // Required: Type of transaction.
	Quantity        int              `json:"quantity"`         // Required: Quantity of the position to convert.
	OldProductType  ProductTypes     `json:"old_product"`      // Required: Old product type of the position.
	NewProductType  ProductTypes     `json:"new_product"`      // Required: New product type to convert the position to.
}

// ModifyGttRequest represents a request to modify a Good 'til Triggered (GTT) order.
type ModifyGttRequest struct {
	Id           uint          `json:"id"`            // Required: Identifier of the GTT order to modify.
	TriggerPrice *float64      `json:"trigger_price"` // Required: New trigger price for the GTT order.
	Price        *float64      `json:"price"`         // Required: New price for the GTT order.
	Quantity     *int          `json:"quantity"`      // Required: New quantity for the GTT order.
	Variety      ValidityTypes `json:"variety"`       // Required: Type of variety. Accepted values: [RL,RL-MKT].
	Trail        *Trail        `json:"trail"`         // Optional: Trail information for the GTT order.
}

// PlaceGttRequest represents a request to place a Good 'til Triggered (GTT) order.
type PlaceGttRequest struct {
	Exchange         ExchangeTypes       `json:"exchange"`           // Required: Exchange type.
	Token            int                 `json:"token"`              // Required: Token of the underlying instrument.
	TransactionType  TransactionTypes    `json:"transaction_type"`   // Required: Type of transaction.
	Variety          ValidityTypes       `json:"variety"`            // Optional: Type of variety. By default RL is considered. Accepted values: [RL,RL-MKT].
	Quantity         *int                `json:"quantity"`           // Required: Quantity of the order.
	TriggerPrice     *float64            `json:"trigger_price"`      // Required: Trigger price for the order.
	Price            *float64            `json:"price"`              // Required: Price of the order.
	OrderIdentifier  string              `json:"order_identifier"`   // Required: Identifier for the order.
	GttTriggerType   GttTriggerType      `json:"gtt_trigger_type"`   // Required: Type of GTT trigger.
	Product          ProductTypes        `json:"product"`            // Required: Type of product.
	Stoploss         *PlaceGttLegRequest `json:"stoploss"`           // Optional: Stop loss leg of the GTT order.
	Profit           *PlaceGttLegRequest `json:"profit"`             // Optional: Profit leg of the GTT order.
	SingleTrailingSL *Trail              `json:"single_trailing_sl"` // Optional: Use when GttTriggerType is single.
	TagIds           []int               `json:"tag_ids"`            // Required: IDs of tags associated with the order.
}

// PlaceGttLegRequest represents a leg of a Good 'til Triggered (GTT) order.
type PlaceGttLegRequest struct {
	Quantity     int           `json:"quantity" binding:"required"`
	Price        *float64      `json:"price"`
	TriggerPrice float64       `json:"trigger_price" binding:"required"`
	ProductType  ProductTypes  `json:"product"`
	Variety      *VarietyTypes `json:"variety"`
	Trail        *Trail        `json:"trail"`
}

type Trail struct {
	ID             uint           `json:"id"`
	TrailJumpPoint *float64       `json:"trail_jump_point"`
	TrailJumpType  *TrailJumpType `json:"trail_jump_type"`
	TriggerTrail   *float64       `json:"trigger_trail"` // This is only used for gtt order book not for accepting the value, we are using ltp as current trail while creating
	TrailFirstJump *float64       `json:"trail_first_jump"`
}

// GttTriggerType represents the trigger type for a Good 'til Triggered (GTT) order.
type GttTriggerType string

const (
	GttTriggerTypeSingle GttTriggerType = "single" // Single trigger type.
	GttTriggerTypeOCO    GttTriggerType = "oco"    // One Cancels the Other (OCO) trigger type.
)

// PlaceIcebergOrderRequest represents a request to place an Iceberg order.
type PlaceIcebergOrderRequest struct {
	Exchange        ExchangeTypes    `json:"exchange"`         // Required: Exchange type.
	Token           int              `json:"token"`            // Required: Token of the underlying instrument.
	TransactionType TransactionTypes `json:"transaction_type"` // Required: Type of transaction.
	Product         ProductTypes     `json:"product"`          // Required: Type of product.
	Variety         VarietyTypes     `json:"variety"`          // Required: Type of variety.
	Quantity        int              `json:"quantity"`         // Required: Quantity of the order.
	Price           *float64         `json:"price"`            // Optional if market order. Price of the order.
	TriggerPrice    float64          `json:"trigger_price"`    // Optional if not stoploss order. Trigger price for the order.
	OrderIdentifier string           `json:"order_identifier"` // Optional: Your identifier for the order.
	Validity        ValidityTypes    `json:"validity"`         // Required: Validity type for the order.
	Legs            int              `json:"legs"`             // Required: Number of legs for the order.
	TagIds          pq.Int32Array    `json:"tag_ids"`          // IDs of tags associated with the order.
}

// FundWithdrawalRequest represents a request to withdraw funds.
type FundWithdrawalRequest struct {
	BankAccountNumber string        `json:"bank_account_number"` // Required: Bank account number.
	Ifsc              string        `json:"ifsc"`                // Required: IFSC code.
	Amount            float64       `json:"amount"`              // Required: Amount to withdraw.
	Exchange          ExchangeTypes `json:"exchange"`            // Required: Exchange type.
}

// FundWithdrawalCancelRequest represents a request to cancel a fund withdrawal.
type FundWithdrawalCancelRequest struct {
	TransactionId string        `json:"transaction_id"` // Required: Transaction ID.
	Exchange      ExchangeTypes `json:"exchange"`       // Required: Exchange type.
	Amount        float64       `json:"amount"`         // Required: Amount of the withdrawal.
}

// TagRequest represents a request to create a tag.
type TagRequest struct {
	Name        string `json:"name"`        // Name of the tag.
	Description string `json:"description"` // Description of the tag.
}

// ModifyIcebergOrderRequest represents a request to modify an Iceberg order.
type ModifyIcebergOrderRequest struct {
	Price          float64 `json:"price"`           // Required: Price of the order.
	TriggerPrice   float64 `json:"trigger_price"`   // Required: Trigger price for the order.
	TradedQuantity int     `json:"traded_quantity"` // Required: Quantity already traded.
}

// StrategiesRequest represents a request to retrieve strategies.
type StrategiesRequest struct {
	Token       int    `json:"token"`       // Required: Token for authentication.
	Symbol      string `json:"symbol"`      // Required: Symbol for the underlying asset.
	ExpYYYYMMDD string `json:"expiry_date"` // Required: Expiry date of the strategy in format YYYYMMDD.
}

// OptionChainRequest represents a request to retrieve an option chain.
type OptionChainRequest struct {
	Token      int           `json:"token"`       // Required: Token of the underlying instrument.
	ExpiryDate string        `json:"expiry_date"` // Optional: Expiry date of options in format YYYYMMDD. If not provided, the result will be of the closest expiry.
	Exchange   ExchangeTypes `json:"exchange"`    // Required: Exchange type.
	AddGreek   bool          `json:"greeks"`      // Optional: Include Greeks in the result. Default is false.
}

type StrategyBuilderRequest struct {
	Token      int            `json:"token"`       // Required: Token of the underlying instrument.
	Symbol     string         `json:"symbol"`      // Required: Symbol of the underlying instrument.
	MarketView PredictionType `json:"prediction"`  // Required: Prediction type. Should be one of ["ABOVE", "BELOW", "BETWEEN"]
	ExpiryDate string         `json:"expiry_date"` // Required: Expiry date of options in format: YYYYMMDD
	PriceRange []float64      `json:"price_range"` // Price range for the strategy.
}

// PredictionType represents the type of prediction for a trading strategy.
type PredictionType string
type PayoffAction string

const (
	PredictionTypeABOVE   PredictionType = "ABOVE"   // Predict that the asset price will be above a certain level.
	PredictionTypeBELOW   PredictionType = "BELOW"   // Predict that the asset price will be below a certain level.
	PredictionTypeBETWEEN PredictionType = "BETWEEN" // Predict that the asset price will be within a certain range.
)

// PayoffAction represents the action for a trading strategy.
const (
	PayoffActionBUY  PayoffAction = "BUY"
	PayoffActionSELL PayoffAction = "SELL"
)

// PayoffRequest represents a request to calculate the payoff for a trading strategy.
type PayoffRequest struct {
	Symbol          string         `json:"symbol" `        // Required: Symbol for the underlying asset.
	Exchange        ExchangeTypes  `json:"exchange" `      // Required: Exchange type.
	Legs            []PayoffOption `json:"legs"`           // Required: Legs of the trading strategy.
	InpDaysToExpiry *int           `json:"days_to_expiry"` // Optional: Number of days to expiry.
	CurrentPnl      float64        `json:"current_pnl"`    // Optional: Current profit or loss.
}

// PayoffOption represents an option within a trading strategy.
type PayoffOption struct {
	Token    int          `json:"token"`            // Required: Token of the underlying instrument.
	Action   PayoffAction `json:"action"`           // Required: Action for the option.
	Quantity int          `json:"quantity"`         // Required: Quantity of the option.
	Ltp      float64      `json:"last_trade_price"` // Optional: Last traded price for the option.
}

type ExchangeAuthTokenRequest struct {
	Checksum      string `json:"checksum"`
	ApplicationId string `json:"applicationId"`
	Token         string `json:"token"`
}

type MultipleOrderCancelRequest struct {
	OrderIds []string `json:"order_ids"`
}
