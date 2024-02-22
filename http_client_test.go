package govortex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"testing"

	httpmock "github.com/jarcoal/httpmock"
)

type HttpTestSuite struct {
	VortexApiClient *VortexApi
}

var HttpMockResponders = [][]interface{}{
	{http.StatusUnauthorized, "Unauthorized access"},
	{http.StatusForbidden, "Forbidden access"},
	{http.StatusNotFound, "Resource not found"},
	{http.StatusTooManyRequests, "Too many requests"},
	{http.StatusServiceUnavailable, "Service unavailable"},
	{http.StatusGatewayTimeout, "Gateway timeout"},
	{http.StatusInternalServerError, "Internal server error"},
	{http.StatusBadRequest, "Bad request"},
	{1, ""},
	{2, "gibberish"},
}

func (ts *HttpTestSuite) SetupAPITest()    {}
func (ts *HttpTestSuite) TearDownAPITest() {}

func (ts *HttpTestSuite) SetupAPITestSuit(t *testing.T) {

	var vortexApi VortexApi
	InitializeVortexApi(applicationId, apiKey, &vortexApi)
	ts.VortexApiClient = &vortexApi
	ts.VortexApiClient.SetLogging(true)
	httpmock.ActivateNonDefault(ts.VortexApiClient.htt.GetClient().client)

	for _, v := range HttpMockResponders {
		httpMethod := "GET"
		route := fmt.Sprintf("%d", v[0])
		base, err := url.Parse(ts.VortexApiClient.baseURL)
		if err != nil {
			panic("Something went wrong")
		}
		// Replace all url variables with string "test"
		re := regexp.MustCompile("%s")
		formattedRoute := re.ReplaceAllString(route, "test")
		base.Path = path.Join(base.Path, formattedRoute)
		if v[0].(int) > 200 {
			res, _ := httpmock.NewJsonResponder(v[0].(int), map[string]interface{}{"status": v})
			httpmock.RegisterResponder(httpMethod, base.String(), res)
		} else {
			switch v[0].(int) {
			case 1:
				res := httpmock.NewErrorResponder(errors.New("error"))
				httpmock.RegisterResponder(httpMethod, base.String(), res)
			case 2:
				res := httpmock.NewStringResponder(v[0].(int), v[1].(string))
				httpmock.RegisterResponder(httpMethod, base.String(), res)
			}

		}

	}
}
func (ts *HttpTestSuite) TestNewHTTPClient(t *testing.T) {
	t.Parallel()
	client := NewHTTPClient(nil, nil, false)
	if client.GetClient() == nil {
		t.Errorf("Error while creating http client. %s", "client is nil")
	}
	if client.GetClient().hLog == nil {
		t.Errorf("Error while creating http client. %s", "log is nil")
	}
	if client.GetClient().debug != false {
		t.Errorf("Error while creating http client. %s", "debug is not false")
	}
}

func RunHttpClientTests(t *testing.T, ts *HttpTestSuite) {
	ts.SetupAPITestSuit(t)
	for _, v := range HttpMockResponders {
		t.Run(fmt.Sprintf("%d %s", v[0], v[1]), func(t *testing.T) {
			ts.SetupAPITest()
			defer ts.TearDownAPITest()
			ctx := context.Background()
			base, err := url.Parse(ts.VortexApiClient.baseURL)
			if err != nil {
				panic("Something went wrong")
			}
			// Replace all url variables with string "test"
			re := regexp.MustCompile("%s")
			route := fmt.Sprintf("%d", v[0])
			formattedRoute := re.ReplaceAllString(route, "test")
			base.Path = path.Join(base.Path, formattedRoute)
			res, err := ts.VortexApiClient.htt.doJSON(ctx, "GET", base.String(), nil, nil, nil, nil)
			if v[0].(int) > 200 {
				if res.Response.StatusCode != v[0] {
					t.Errorf("not nil %d", v[0])
					return
				}
			} else {
				if err == nil {
					t.Errorf("Error is nil")
					return
				}
			}

		})
	}
}

func TestHttpClient(t *testing.T) {
	s := &HttpTestSuite{}
	s.TestNewHTTPClient(t)
	RunHttpClientTests(t, s)
}
