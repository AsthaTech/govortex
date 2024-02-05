package govortex

import "github.com/lib/pq"

type PlaceOrderRequest struct {
	Exchange          ExchangeTypes    `json:"exchange"`
	Token             int              `json:"token"`
	TransactionType   TransactionTypes `json:"transaction_type"`
	Product           ProductTypes     `json:"product"`
	Variety           VarietyTypes     `json:"variety"`
	Quantity          int              `json:"quantity"`
	Price             float64          `json:"price"`
	TriggerPrice      float64          `json:"trigger_price"`
	OrderIdentifier   string           `json:"order_identifier"`
	DisclosedQuantity int              `json:"disclosed_quantity"`
	Validity          ValidityTypes    `json:"validity"`
	ValidityDays      int              `json:"validity_days"`
	IsAMO             bool             `json:"is_amo"`
	Gtt               *GttLegs         `json:"gtt"`
	TagIds            []int            `json:"tag_ids"`
}

type GttLegs struct {
	SlTriggerPercent     *float64 `json:"sl_trigger_percent"`
	ProfitTriggerPercent *float64 `json:"profit_trigger_percent"`
}

type ModifyOrderRequest struct {
	Variety           VarietyTypes  `json:"variety" `
	Quantity          int           `json:"quantity" `
	TradedQuantity    *int          `json:"traded_quantity" `
	Price             float64       `json:"price" `
	TriggerPrice      float64       `json:"trigger_price"`
	DisclosedQuantity int           `json:"disclosed_quantity"`
	Validity          ValidityTypes `json:"validity" `
	ValidityDays      int           `json:"validity_days"`
	TagIds            []int         `json:"tag_ids"`
}

type OrderMarginRequest struct {
	Exchange        ExchangeTypes    `json:"exchange"`
	Token           int              `json:"token"`
	TransactionType TransactionTypes `json:"transaction_type"`
	Product         ProductTypes     `json:"product"`
	Variety         VarietyTypes     `json:"variety"`
	Quantity        int              `json:"quantity"`
	Price           float64          `json:"price"`
	OldPrice        float64          `json:"old_price"`
	OldQuantity     int              `json:"old_quantity"`
	Mode            MarginModes      `json:"mode"`
}

type ConvertPositionRequest struct {
	Exchange        ExchangeTypes    `json:"exchange" `
	Token           int              `json:"token" `
	TransactionType TransactionTypes `json:"transaction_type" `
	Quantity        int              `json:"quantity"  `
	OldProductType  ProductTypes     `json:"old_product" `
	NewProductType  ProductTypes     `json:"new_product" `
}

type ModifyGttRequest struct {
	Id           uint     `json:"id" `
	TriggerPrice *float64 `json:"trigger_price"`
	Price        *float64 `json:"price"`
	Quantity     *int     `json:"quantity"`
}

type PlaceGttRequest struct {
	Exchange        ExchangeTypes       `json:"exchange"`
	Token           int                 `json:"token"`
	TransactionType TransactionTypes    `json:"transaction_type"`
	Quantity        *int                `json:"quantity"`
	TriggerPrice    *float64            `json:"trigger_price"`
	Price           *float64            `json:"price"`
	OrderIdentifier string              `json:"order_identifier"`
	GttTriggerType  GttTriggerType      `json:"gtt_trigger_type"`
	Product         ProductTypes        `json:"product"`
	Stoploss        *PlaceGttLegRequest `json:"stoploss"`
	Profit          *PlaceGttLegRequest `json:"profit"`
	TagIds          []int               `json:"tag_ids"`
}
type PlaceGttLegRequest struct {
	Quantity     int          `json:"quantity"`
	Price        float64      `json:"price"`
	TriggerPrice float64      `json:"trigger_price"`
	ProductType  ProductTypes `json:"product"`
}

type GttTriggerType string

const (
	GttTriggerTypeSingle GttTriggerType = "single"
	GttTriggerTypeOCO    GttTriggerType = "oco"
)

type PlaceIcebergOrderRequest struct {
	Exchange        ExchangeTypes    `json:"exchange" `
	Token           int              `json:"token" `
	TransactionType TransactionTypes `json:"transaction_type" `
	Product         ProductTypes     `json:"product" `
	Variety         VarietyTypes     `json:"variety" `
	Quantity        int              `json:"quantity" `
	Price           *float64         `json:"price" `
	TriggerPrice    float64          `json:"trigger_price"`
	OrderIdentifier string           `json:"order_identifier"`
	Validity        ValidityTypes    `json:"validity" `
	Legs            int              `json:"legs" `
	TagIds          pq.Int32Array    `json:"tag_ids"`
}

type FundWithdrawalRequest struct {
	BankAccountNumber string        `json:"bank_account_number"`
	Ifsc              string        `json:"ifsc"`
	Amount            float64       `json:"amount"`
	Exchange          ExchangeTypes `json:"exchange"`
}

type FundWithdrawalCancelRequest struct {
	TransactionId string        `json:"transaction_id"`
	Exchange      ExchangeTypes `json:"exchange"`
	Amount        float64       `json:"amount"`
}

type TagRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ModifyIcebergOrderRequest struct {
	Price          float64 `json:"price"`
	TriggerPrice   float64 `json:"trigger_price" `
	TradedQuantity int     `json:"traded_quantity"`
}
