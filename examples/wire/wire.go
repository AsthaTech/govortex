package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AsthaTech/govortex"
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
	var wire govortex.Wire
	govortex.InitializeWire(accessToken, &wire)

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
