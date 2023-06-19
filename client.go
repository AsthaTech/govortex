package govortex

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Constants for the package
const (
	name           string        = "govortex"
	version        string        = "1.0.0"
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://vortex.restapi.asthatrade.com"
	flowBaseURI    string        = "https://flow.asthatrade.com"
)

// Constants for the API endpoints
const (
	URILogin        string = "/user/login"
	URIInstruments  string = "/data/instruments"
	URIPlaceOrder   string = "/orders/%s"       // "/orders/regular"
	URIModifyOrder  string = "/orders/%s/%s/%s" //"/orders/regular/{exchange}/{order_id}"
	URIDeleterOrder string = "/orders/%s/%s/%s" //"/orders/regular/{exchange}/{order_id}"
	URIOrderBook    string = "/orders"          //"/orders?limit={limit}&offset={offset}"
	URIOrderHistory string = "/orders/%s"       //"/orders/{order_id}"
	URITrades       string = "/trades"

	URIPositions string = "/portfolio/positions"
	URIHoldings  string = "/portfolio/holdings"

	URIFunds       string = "/user/funds"
	URIOrderMargin string = "/margins/order"
	URIQuotes      string = "/data/quote"
	URIHistory     string = "/data/history"
)

// VortexApi is a struct representing the Vortex API client
type VortexApi struct {
	applicationId string
	apiKey        string
	AccessToken   string
	baseURL       string
	enableLogging bool
	htt           HTTPClient
}

// Function to get a new instance of VortexApi Client.
func NewVortexApi(applicationId string, apiKey string) *VortexApi {
	v := VortexApi{}
	v.initialize(applicationId, apiKey)
	return &v
}

// Initialize sets the application ID and API key for the Vortex API client
func (v *VortexApi) initialize(applicationId string, apiKey string) {
	v.applicationId = applicationId
	v.apiKey = apiKey
	v.SetHTTPClient(&http.Client{
		Timeout: requestTimeout,
	})
	v.baseURL = baseURI
}

// SetLogging sets the HTTP client with logging enabled
func (v *VortexApi) SetLogging(h *http.Client) {
	v.htt = NewHTTPClient(h, nil, v.enableLogging)
}

// SetHTTPClient sets the HTTP client for the Vortex API client
func (v *VortexApi) SetHTTPClient(h *http.Client) {
	v.htt = NewHTTPClient(h, nil, v.enableLogging)
}
func (v *VortexApi) SetAccessToken(accessToken string) {
	v.AccessToken = accessToken
}

// GetLoginUrl returns the login URL for the Vortex API
func (v *VortexApi) GetLoginUrl() string {
	return fmt.Sprintf("%s/?applicationId=%s", flowBaseURI, v.applicationId)
}

// doJson is a helper function that performs a JSON request to the Vortex API
func (v *VortexApi) doJson(ctx context.Context, method, uri string, body interface{}, params url.Values, headers http.Header, obj interface{}) (HTTPResponse, error) {
	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("User-Agent", name+"/"+version)

	if v.AccessToken != "" {
		headers.Add("Authorization", "Bearer "+v.AccessToken)
	}
	return v.htt.doJSON(ctx, method, v.baseURL+uri, body, params, headers, obj)
}
