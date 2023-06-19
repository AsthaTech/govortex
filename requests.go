package govortex

type PlaceOrderRequest struct {
	Exchange          ExchangeTypes    `json:"exchange"`
	Token             int              `json:"token"`
	TransactionType   TransactionTypes `json:"transaction_type"`
	Product           ProductTypes     `json:"product"`
	Variety           VarietyTypes     `json:"variety"`
	Quantity          int              `json:"quantity"`
	Price             float64          `json:"price"`
	TriggerPrice      float64          `json:"trigger_price"`
	DisclosedQuantity int              `json:"disclosed_quantity"`
	Validity          ValidityTypes    `json:"validity"`
	ValidityDays      int              `json:"validity_days"`
	IsAMO             bool             `json:"is_amo"`
}

type ModifyOrderRequest struct {
	Variety           VarietyTypes  `json:"variety"`
	Quantity          int           `json:"quantity"`
	TradedQuantity    int           `json:"traded_quantity"`
	Price             float64       `json:"price"`
	TriggerPrice      float64       `json:"trigger_price"`
	DisclosedQuantity int           `json:"disclosed_quantity"`
	Validity          ValidityTypes `json:"validity"`
	ValidityDays      int           `json:"validity_days"`
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
