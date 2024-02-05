package main

import (
	"context"
	"encoding/json"
	"fmt"

	govortex "github.com/AsthaTech/govortex"
)

const (
	applicationId string = "dev_hK7M2HWR"
	apiKey        string = "TYB7g0bdJFvvQOfIXsgquaDfupLieYZlOHNMRAeO"
)

func main() {
	var client govortex.VortexApi
	govortex.InitializeVortexApi(applicationId, apiKey, &client)
	client.SetLogging(true)
	ctx := context.Background()
	client.Login(ctx, "600S1871", "ApiTesting1$", "136000")
	client.SetAccessToken("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcxODY2MDAsImlhdCI6MTcwNzEzMjQyMCwiaXNzIjoiQXN0aGFUcmFkZSIsInN1YiI6IjYwMFMxODcxIiwiZGV2aWNlSWQiOiJhcGktY2xpZW50IiwiYXBwbGljYXRpb25JZCI6ImRldl9oSzdNMkhXUiIsImlwIjoiMTMuMjMyLjY0LjE0MSIsInNjb3BlIjoicHJvZmlsZS5yZWFkLG9yZGVycy5yZWFkLG9yZGVycy53cml0ZSxob2xkaW5ncy5yZWFkLHBvc2l0aW9ucy5yZWFkLHBvc2l0aW9ucy53cml0ZSxmdW5kcy5yZWFkLHRyYWRlcy5yZWFkLG1hcmdpbnMucmVhZCxkYXRhLnJlYWQiLCJjbGllbnRfY29kZSI6IjYwMFMxODcxIiwic291cmNlIjoiYXBpX2F1dGgiLCJwb2EiOnRydWV9.w4Lzj3Gu1tyeVjrWBMODMuUz3sHBLHRtEogTPImoT5uF_C7og4pvolddaET3AkqrJ_Z6fSI9fy-czWyUxHRuzsi0OijLatbdJidWzX48HqMrropBj5ggAq7Uhkqpo7qcqympmDI9ej8p3f91bZCWIevgnEGFJI9zNIu_NnIIqTU94q0kM5xiks5HNwrE4abS6jlTJ_OpuKsXgRhIqCtdgSl_QTcSaniZ0HHsYoxwJFhO2Mr-lcD22lP8BArNGpp765r36lkN1FhDvYwt2j-tRLOjWkUe-RyoFPBWh-whQ-BiE1aN0f53FwMI_TL2nf3pgM5M_CvOdF3_SEi4CuNfgAbhStirTq-E2MTLzrxOD3p5dMhms9dOlWmwK8uLV7xA2puMETUPY0_gyKAjQv9LX-29vFO6rKJzbJDMVJsD_QXQYDfJ1j35G_dPoWKRTjObWxcWOOdCBPXtcxj1N-oYO-2-ZFyGtEm-7c-HMnQ89fS4Td0brWMD8qSLIlV7nAlZ")

	// Access token is automatically set upon successful login call

	// client.Orders(ctx)    //orders need an offset and limit
	// client.Positions(ctx) //positions need an offset and limit

	// client.PlaceOrder(ctx, govortex.PlaceOrderRequest{
	// 	Exchange:          govortex.ExchangeTypesNSEEQUITY,
	// 	Token:             22,
	// 	TransactionType:   govortex.TransactionTypesBUY,
	// 	Product:           govortex.ProductTypesDelivery,
	// 	Variety:           govortex.VarietyTypesRegularMarketOrder,
	// 	Quantity:          1,
	// 	Price:             1800,
	// 	TriggerPrice:      0,
	// 	DisclosedQuantity: 0,
	// 	Validity:          govortex.ValidityTypesFullDay,
	// 	ValidityDays:      1,
	// 	IsAMO:             false,
	// })
	// client.Funds(ctx)

	aa, _ := client.Tags(ctx)
	bb, _ := json.Marshal(aa)
	fmt.Println(string(bb))

}
