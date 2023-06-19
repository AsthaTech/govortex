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

func TestNewClient(t *testing.T) {
	t.Parallel()

	apiKey := "testApiKey"
	applicationId := "testApplicationId"

	vortexApi := NewVortexApi(applicationId, apiKey)
	if vortexApi.apiKey != apiKey {
		t.Errorf("Api Key is not assigned properly.")
	}
}

const mockBaseDir = "./vortex-mocks"

var MockResponders = [][]interface{}{
	// GET REQUESTS

	{http.MethodGet, URIFunds, url.Values{}, "funds.json"},
	{http.MethodGet, URIHoldings, url.Values{}, "holdings.json"},
	{http.MethodGet, URIPositions, url.Values{}, "positions.json"},
	{http.MethodGet, URIOrderBook, url.Values{"limit": []string{"20"}, "offset": []string{"1"}}, "orders.json"},
	{http.MethodGet, URIQuotes, url.Values{"q": []string{"NSE_EQ-22"}, "mode": []string{string(QuoteModesFULL)}}, "quotes.json"},
	{http.MethodGet, URITrades, url.Values{"limit": []string{"20"}, "offset": []string{"1"}}, "trades.json"},
	{http.MethodGet, URIHistory, url.Values{"exchange": []string{string(ExchangeTypesNSEEQUITY)}, "token": []string{"22"}, "from": []string{"1494505756"}, "to": []string{"1494505756"}, "resolution": []string{string(ResolutionsMin1)}}, "history.json"},

	// DELETE REQUESTS
	{http.MethodDelete, fmt.Sprintf(URIDeleterOrder, "regular", ExchangeTypesNSEEQUITY, "NXAAE00002K3"), url.Values{}, "order.json"},
	// POST  REQUESTS
	{http.MethodPost, fmt.Sprintf(URIPlaceOrder, "regular"), url.Values{}, "order.json"},

	// PUT REQUESTS
	{http.MethodPut, fmt.Sprintf(URIModifyOrder, "regular", ExchangeTypesNSEEQUITY, "NXAAE00002K3"), url.Values{}, "order.json"},
}

const suiteTestMethodPrefix = "Test"

type TestSuite struct {
	VortexApiClient *VortexApi
}

func (ts *TestSuite) SetupAPITestSuit(t *testing.T) {

	ts.VortexApiClient = NewVortexApi("testApplicationId", "testApiSecret")
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
func (ts *TestSuite) TearDownAPITestSuit() {
	httpmock.DeactivateAndReset()
}
func (ts *TestSuite) SetupAPITest() {}

// Individual test teardown
func (ts *TestSuite) TearDownAPITest() {}

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

func TestAPIMethods(t *testing.T) {
	s := &TestSuite{}
	RunAPITests(t, s)
}
