package main

import (
	"context"

	govortex "github.com/AsthaTech/govortex"
)

const (
	applicationId string = "testApplicationId"
	apiKey        string = "testApiKey"
)

func main() {
	var client govortex.VortexApi
	govortex.InitializeVortexApi(applicationId, apiKey, &client)
	ctx := context.Background()
	client.Login(ctx, "clientCode", "password", "totp")
	// Access token is automatically set upon successful login call

	client.Orders(ctx, 1, 20) //orders need an offset and limit
	client.Positions(ctx)     //positions need an offset and limit

	client.PlaceOrder(ctx, govortex.PlaceOrderRequest{
		Exchange:          govortex.ExchangeTypesNSEEQUITY,
		Token:             22,
		TransactionType:   govortex.TransactionTypesBUY,
		Product:           govortex.ProductTypesDelivery,
		Variety:           govortex.VarietyTypesRegularMarketOrder,
		Quantity:          1,
		Price:             1800,
		TriggerPrice:      0,
		DisclosedQuantity: 0,
		Validity:          govortex.ValidityTypesFullDay,
		ValidityDays:      1,
		IsAMO:             false,
	})

	client.Funds(ctx)

}
