package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AsthaTech/govortex/v2"
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
	accessToken := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4NTcwMDAsImlhdCI6MTc0NDc4MTU5MywiaXNzIjoiQXN0aGFUcmFkZSIsInN1YiI6IjYwMFMxODcxIiwiZGV2aWNlSWQiOiIxNzQxODUxNTIzNzE3IiwiYXBwbGljYXRpb25JZCI6ImF0X0VKMDJ2V0tEd1MiLCJpcCI6IjEwNi41MS43NS42NyIsInNjb3BlIjoiKiIsImNsaWVudF9jb2RlIjoiNjAwUzE4NzEiLCJzb3VyY2UiOiJhcGlfYXV0aCIsInBvYSI6dHJ1ZSwidG9rZW5fdHlwZSI6InZvcnRleF92MyIsIm9zIjoiIiwiZGV2aWNlSW5mbyI6IiIsIm1mX2Z1bmRfdHlwZSI6InJlZ3VsYXIifQ.RhOWtKzOawBdF9-aCV-3MXndjyhp09_ld_LPxeKDy8v_j3u_kaZ_gyqe6kqkA_Fl9VJQlPLnP3wJGnQr3SFSJKPlMFgZH5OyDtrvWFsXDabV2keR247YSIH7k1Uo0QGequ5u-VbubAD1U3XCue16xkP7EF4aqLZsNNXRkYQ-TIBQLWPQFyyfBh-6WoKhhSMeuQoBrWgXerW6QC0zY_uZ8XXOcao6l3Uto_is_wRA3GWLe5qjZQerBKPBXs2utEkLO2YpQEqbCsemxz80Ic6MeRKyr034oT9H-lOFuXARSzvH2DIoC4hN6jE6d4ARsUzY3hL7OogUwDrAHPy8u08kuzt9U_ipMnN8nRH0gDvpfmNEF1wuwUnJFX8Q0CTx4K_yzvUfug1LW4YmALBO9blj3-ut-414sdwlhEonUuwJWAS8Y3sO1wnSmucdbH_EuX1iVNdkQLQOHu9r276ULExLBZuVs2g7Q8HwInbvLmIN3BCz7gJkWjw1dMbIN-UWW-og"

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
	wire.Subscribe(govortex.ExchangeTypesNSEEQUITY, 26017, govortex.QuoteModesFULL)
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
