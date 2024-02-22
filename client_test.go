package govortex

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"strings"
	"testing"

	httpmock "github.com/jarcoal/httpmock"
)

const (
	apiKey        = "testApiKey"
	applicationId = "testApplicationId"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	var vortexApi VortexApi
	InitializeVortexApi(applicationId, apiKey, &vortexApi)
	if vortexApi.apiKey != apiKey {
		t.Errorf("Api Key is not assigned properly.")
	}
}

const mockBaseDir = "./vortex-mocks"

var MockResponders = [][]interface{}{
	// GET REQUESTS

	{http.MethodGet, URIFunds, url.Values{}, "user/funds.json"},
	{http.MethodGet, URIHoldings, url.Values{}, "portfolio/holdings.json"},
	{http.MethodGet, URIPositions, url.Values{}, "portfolio/positions.json"},
	{http.MethodGet, URIOrderBook, url.Values{}, "portfolio/orders.json"},
	{http.MethodGet, URIQuotes, url.Values{"q": []string{"NSE_EQ-22"}, "mode": []string{string(QuoteModesFULL)}}, "data/quotes.json"},
	{http.MethodGet, URITrades, url.Values{"limit": []string{"20"}, "offset": []string{"1"}}, "portfolio/trades.json"},
	{http.MethodGet, URIHistory, url.Values{"exchange": []string{string(ExchangeTypesNSEEQUITY)}, "token": []string{"22"}, "from": []string{"1494505756"}, "to": []string{"1494505756"}, "resolution": []string{string(ResolutionsMin1)}}, "data/history.json"},
	{http.MethodGet, fmt.Sprintf(URIOrderHistory, "test"), url.Values{}, "regular_orders/order_history.json"},
	{http.MethodGet, URIGttOrderBook, url.Values{}, "gtt_orders/list.json"},
	{http.MethodGet, URITags, url.Values{}, "reports/tags/list.json"},
	{http.MethodGet, URIBanks, url.Values{}, "user/banks.json"},

	// DELETE REQUESTS
	{http.MethodDelete, fmt.Sprintf(URIDeleteOrder, "regular", "NXAAE00002K3"), url.Values{}, "regular_orders/order.json"},
	{http.MethodDelete, fmt.Sprintf(URIDeleteOrder, "gtt", "99823d7b-fd37-4d75-af7f-f21ec4671852"), url.Values{}, "gtt_orders/delete.json"},
	{http.MethodDelete, fmt.Sprintf(URITag, 1), url.Values{}, "reports/tags/delete.json"},
	{http.MethodDelete, fmt.Sprintf(URIDeleteOrder, "iceberg", "5eaefd25-518c-4a39-b556-93fc8e78e855"), url.Values{}, "iceberg_orders/delete.json"},
	{http.MethodDelete, URISession, url.Values{}, "user/logout.json"},
	// POST  REQUESTS
	{http.MethodPost, fmt.Sprintf(URIPlaceOrder, "regular"), url.Values{}, "regular_orders/order.json"},
	{http.MethodPost, fmt.Sprintf(URIPlaceOrder, "gtt"), url.Values{}, "gtt_orders/create.json"},
	{http.MethodPost, fmt.Sprintf(URIPlaceOrder, "iceberg"), url.Values{}, "iceberg_orders/create.json"},
	{http.MethodPost, URITags, url.Values{}, "reports/tags/create.json"},
	{http.MethodPost, URIStrategies, url.Values{}, "strategies/all.json"},
	{http.MethodPost, URIOptionChain, url.Values{}, "strategies/option_chain.json"},
	{http.MethodPost, URIBuildStrategies, url.Values{}, "strategies/build_strategy.json"},
	{http.MethodPost, URIPayoffStrategies, url.Values{}, "strategies/payoff.json"},
	{http.MethodPost, URIMultiCancelrders, url.Values{}, "regular_orders/multi_cancel.json"},
	{http.MethodPost, URIBasketMargin, url.Values{}, "margins/basket_margin.json"},
	{http.MethodPost, URIOrderMargin, url.Values{}, "margins/order_margin.json"},
	{http.MethodPost, URISession, url.Values{}, "user/login.json"},
	{http.MethodPost, URILogin, url.Values{}, "user/login.json"},

	// PUT REQUESTS
	{http.MethodPut, fmt.Sprintf(URIModifyOrder, "regular", "NXAAE00002K3"), url.Values{}, "regular_orders/order.json"},
	{http.MethodPut, URIConvertposition, url.Values{}, "portfolio/position_conversion.json"},
	{http.MethodPut, fmt.Sprintf(URIModifyOrder, "gtt", "99823d7b-fd37-4d75-af7f-f21ec4671852"), url.Values{}, "gtt_orders/modify.json"},
	{http.MethodPut, fmt.Sprintf(URIModifyOrder, "iceberg", "5eaefd25-518c-4a39-b556-93fc8e78e855"), url.Values{}, "iceberg_orders/modify.json"},
	{http.MethodPut, fmt.Sprintf(URITag, 1), url.Values{}, "reports/tags/update.json"},
}

const suiteTestMethodPrefix = "Test"

type TestSuite struct {
	VortexApiClient *VortexApi
}

type ErrorTestSuite struct {
	VortexApiClient *VortexApi
}

func (ts *TestSuite) SetupAPITestSuit(t *testing.T) {

	var vortexApi VortexApi
	InitializeVortexApi(applicationId, apiKey, &vortexApi)
	ts.VortexApiClient = &vortexApi
	ts.VortexApiClient.SetLogging(true)
	httpmock.ActivateNonDefault(ts.VortexApiClient.htt.GetClient().client)

	for _, v := range MockResponders {
		httpMethod := v[0].(string)
		route := v[1].(string)
		params := v[2].(url.Values)
		filePath := v[3].(string)

		resp, err := ioutil.ReadFile(path.Join(mockBaseDir, filePath))
		if err != nil {
			panic("Error while reading mock response: " + filePath)
		}

		base, err := url.Parse(ts.VortexApiClient.baseURL)
		if err != nil {
			panic("Something went wrong")
		}
		// Replace all url variables with string "test"
		re := regexp.MustCompile("%s")
		formattedRoute := re.ReplaceAllString(route, "test")
		base.Path = path.Join(base.Path, formattedRoute)
		base.RawQuery = params.Encode()
		httpmock.RegisterResponder(httpMethod, base.String(), httpmock.NewBytesResponder(200, resp))
	}

}

func (ts *ErrorTestSuite) SetupAPITestSuit(t *testing.T) {

	var vortexApi VortexApi
	InitializeVortexApi(applicationId, apiKey, &vortexApi)
	ts.VortexApiClient = &vortexApi
	httpmock.ActivateNonDefault(ts.VortexApiClient.htt.GetClient().client)

	for _, v := range MockResponders {
		httpMethod := v[0].(string)
		route := v[1].(string)
		params := v[2].(url.Values)

		base, err := url.Parse(ts.VortexApiClient.baseURL)
		if err != nil {
			panic("Something went wrong")
		}
		// Replace all url variables with string "test"
		re := regexp.MustCompile("%s")
		formattedRoute := re.ReplaceAllString(route, "test")
		base.Path = path.Join(base.Path, formattedRoute)
		base.RawQuery = params.Encode()
		res, _ := httpmock.NewJsonResponder(429, map[string]interface{}{"status": "Too many requests"})
		httpmock.RegisterResponder(httpMethod, base.String(), res)
	}
	httpmock.GetCallCountInfo()

}
func (ts *TestSuite) TearDownAPITestSuit() {
	fmt.Println(httpmock.GetCallCountInfo())
	httpmock.DeactivateAndReset()
}
func (ts *ErrorTestSuite) TearDownAPITestSuit() {
	fmt.Println(httpmock.GetCallCountInfo())
	httpmock.DeactivateAndReset()
}
func (ts *TestSuite) SetupAPITest()      {}
func (ts *ErrorTestSuite) SetupAPITest() {}

// Individual test teardown
func (ts *TestSuite) TearDownAPITest()      {}
func (ts *ErrorTestSuite) TearDownAPITest() {}

func RunAPITests(t *testing.T, ts *TestSuite) {
	ts.SetupAPITestSuit(t)
	suiteType := reflect.TypeOf(ts)
	for i := 0; i < suiteType.NumMethod(); i++ {
		m := suiteType.Method(i)
		if strings.HasPrefix(m.Name, suiteTestMethodPrefix) {
			t.Run(m.Name, func(t *testing.T) {
				ts.SetupAPITest()
				defer ts.TearDownAPITest()

				in := []reflect.Value{reflect.ValueOf(ts), reflect.ValueOf(t)}
				m.Func.Call(in)
			})
		}
	}
	// ts.TearDownAPITestSuit()
}

func RunAPIErrorTests(t *testing.T, ts *ErrorTestSuite) {
	ts.SetupAPITestSuit(t)
	suiteType := reflect.TypeOf(ts)
	for i := 0; i < suiteType.NumMethod(); i++ {
		m := suiteType.Method(i)
		if strings.HasPrefix(m.Name, suiteTestMethodPrefix) {
			t.Run(m.Name, func(t *testing.T) {
				ts.SetupAPITest()
				defer ts.TearDownAPITest()

				in := []reflect.Value{reflect.ValueOf(ts), reflect.ValueOf(t)}
				m.Func.Call(in)
			})
		}
	}
	// ts.TearDownAPITestSuit()
}

func TestAPIMethods(t *testing.T) {
	s := &TestSuite{}
	RunAPITests(t, s)
}

func TestApiErrorMethods(t *testing.T) {
	e := &ErrorTestSuite{}
	RunAPIErrorTests(t, e)
}

func checkError429(t *testing.T, err error) {
	if err == nil {
		t.Error("Expected an error but got nil")
	}
	if err != nil && err.Error() != "Too many requests" {
		t.Errorf("Not correct error %v", err)
	}
}
