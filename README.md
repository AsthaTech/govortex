[![Go Reference](https://pkg.go.dev/badge/github.com/RupeezyTech/govortex.svg)](https://pkg.go.dev/github.com/RupeezyTech/govortex)
# Vortex API Golang Client


Official golang client for communicating with [Vortex API](https://rupeezy.in/vortex)

Vortex APIs are meant for clients who want to execute orders based on their own strategy programatically and for partners to build their own applications. These apis provide a fast and secure way to place trades, manage positions and access real time market data.


## Documentation 
- [Go Documentation](https://pkg.go.dev/github.com/RupeezyTech/govortex)
- [API Documentation](https://vortex.rupeezy.in/docs/)

## Installation 

```
    go get github.com/RupeezyTech/govortex/v2
```


## Getting Started with APIs 

```go 
package main

import (
	"context"

	govortex "github.com/RupeezyTech/govortex/v2"
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

```

## Getting started with Feed 
```go 
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/RupeezyTech/govortex/v2"
)

var wire *govortex.Wire

const (
	applicationId string = "testApplicationId"
	apiKey        string = "testApiKey"
)

func main() {
	var client govortex.VortexApi
	govortex.InitializeVortexApi(applicationId, apiKey, &client)
	ctx := context.Background()
	client.Login(ctx, "clientCode", "password", "totp")
	accessToken := client.AccessToken

	// Initialize the wire client
	wire = &govortex.Wire{}
	govortex.InitializeWire(accessToken, wire)

	// Define callbacks
	wire.OnConnect(onConnect)
	wire.OnPriceUpdate(onPriceUpdate)
	wire.OnError(onError)
	wire.OnClose(onClose)
	wire.OnReconnect(onReconnect)
	wire.OnOrderUpdate(onOrderUpdate)

	// Start websocket server
	wire.Serve()
}

func onConnect() {
	fmt.Println("connected")
	wire.Subscribe(govortex.ExchangeTypesNSEEQUITY, 22, govortex.QuoteModesFULL)
}

func onPriceUpdate(q *govortex.FullQuoteData) {
	aa, _ := json.Marshal(q)
	fmt.Println(string(aa))
}

func onError(err error) {
	fmt.Println(err.Error())
}

func onClose(code int, reason string) {
	fmt.Println(code, reason)
}

func onReconnect(attempt int, delay time.Duration) {
	fmt.Println(attempt, delay)
}

func onOrderUpdate(msg govortex.SocketMessage) {
	aa, _ := json.Marshal(msg)
	fmt.Println(string(aa))
}

```

## Run Tests 

```
go test
```

## Generate Docs 

```
go doc -all > doc.txt
```