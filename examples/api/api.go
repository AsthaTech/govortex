package main

import (
	"context"
	"fmt"

	govortex "github.com/AsthaTech/govortex/v2"
)

const (
	applicationId string = "testApplicationId"
	apiKey        string = "testApiKey"
)

func main() {
	var client govortex.VortexApi
	govortex.InitializeVortexApi(applicationId, apiKey, &client)
	ctx := context.Background()

	// Open this URL in the browser to login. The page will redirect to the callback URL configured on the API Portal.
	// The callback URL will contain the auth_token as a query parameter.
	url := client.GetLoginUrl()
	fmt.Printf("URL for login: %s\n", url)

	auth_token := "your_auth_token_here" // Replace with the auth_token received after login

	_, err := client.ExchangeToken(ctx, auth_token)
	if err != nil {
		fmt.Println("Error exchanging token:", err)
		return
	}
	fmt.Println("Access Token:", client.AccessToken)
	// Access token is automatically set upon successful login call

	client.Orders(ctx)    //orders need an offset and limit
	client.Positions(ctx) //positions need an offset and limit

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
