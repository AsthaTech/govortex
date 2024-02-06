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
	version        string        = "2.0.0"
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://vortex-api.rupeezy.in/v2"
	flowBaseURI    string        = "https://flow.rupeezy.in"
)

// Constants for the API endpoints
const (
	URILogin               string = "/user/login"
	URIInstruments         string = "/data/instruments"
	URIPlaceOrder          string = "/trading/orders/%s"              //"/trading/orders/regular"
	URIModifyOrder         string = "/trading/orders/%s/%s"           //"/trading/orders/{{order_type}}/{order_id}"
	URIModifyOrderTags     string = "/trading/orders/%s/%s/tags"      //"/trading/orders/{{order_type}}/{order_id}/tags"
	URIDeleteOrder         string = "/trading/orders/%s/%s"           //"/orders/{{order_type}}/{order_id}"
	URIDeleteMultipleOrder string = "/trading/orders/%s/multi_delete" //"/orders/{{order_type}}/{order_id}"
	URIGttOrderBook        string = "/trading/orders/gtt"             //"/orders/gtt"
	URIOrderBook           string = "/trading/orders"                 //"/orders"
	URIOrderHistory        string = "/trading/orders/%s"              //"/orders/{order_id}"
	URITrades              string = "/trading/trades"

	URIPositions       string = "/trading/portfolio/positions"
	URIConvertposition string = "/trading/portfolio/positions"
	URIHoldings        string = "/trading/portfolio/holdings"

	URIFunds     string = "/trading/user/funds"
	URIBanks     string = "/trading/user/banks"
	URIBrokerage string = "/trading/user/brokerage"

	URIWithdrawal  string = "/user/funds/withdrawal"
	URIOrderMargin string = "/margins/order"
	URIOrderBasket string = "/margins/basket"
	URIQuotes      string = "/data/quote"
	URIHistory     string = "/data/history"

	URIOptionChain string = "/strategies/option_chain"

	URIStrategies       string = "/strategies"
	URIBuildStrategies  string = "/strategies/build"
	URIPayoffStrategies string = "/strategies/payoff"

	URITradeReport           string = "/reports/trades/%s?from_date=%s&to_date=%s"
	URITurnoverSummaryReport string = "/reports/turnover/summary/%s?from_date=%s&to_date=%s"
	URITurnoverDetailsReport string = "/reports/turnover/details/%s?from_date=%s&to_date=%s"
	URIPnLReport             string = "/reports/pnl/%s?from_date=%s&to_date=%s"
	URIMTFInterestReport     string = "/reports/mtf_interest/%s?from_date=%s&to_date=%s"

	URITags string = "/reports/tags"
	URITag  string = "/reports/tags/%d"
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
func InitializeVortexApi(applicationId string, apiKey string, apiClient *VortexApi) error {
	apiClient.initialize(applicationId, apiKey)
	return nil
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
func (v *VortexApi) SetLogging(flag bool) {
	v.htt = NewHTTPClient(v.htt.GetClient().client, nil, flag)
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
	fmt.Println("URL:", v.baseURL+uri)
	return v.htt.doJSON(ctx, method, v.baseURL+uri, body, params, headers, obj)
}
