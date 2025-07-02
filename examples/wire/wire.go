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
	// Initialize the wire client
	wire = &govortex.Wire{}
	govortex.InitializeWire(client.AccessToken, wire)

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
