package govortex

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// HTTPClient defines the interface for performing HTTP requests.
type HTTPClient interface {
	do(ctx context.Context, method, rURL string, reqBody interface{}, params url.Values, headers http.Header) (HTTPResponse, error)
	doRaw(ctx context.Context, method, rURL string, reqBody interface{}, params url.Values, headers http.Header) (HTTPResponse, error)
	doJSON(ctx context.Context, method, rURL string, reqBody interface{}, params url.Values, headers http.Header, obj interface{}) (HTTPResponse, error)
	GetClient() *httpClient
}

// httpClient is an implementation of the HTTPClient interface.
type httpClient struct {
	client *http.Client
	hLog   *log.Logger
	debug  bool
}

// HTTPResponse contains the response body and the HTTP response object.
type HTTPResponse struct {
	Body     []byte
	Response *http.Response
}

// NewHTTPClient creates a new instance of the httpClient with the given HTTP client, logger, and debug flag.
// If the logger is nil, it uses a default logger that writes to os.Stdout.
// If the HTTP client is nil, it uses a default client with a 5-second timeout and default transport settings.
func NewHTTPClient(h *http.Client, hLog *log.Logger, debug bool) HTTPClient {
	if hLog == nil {
		hLog = log.New(os.Stdout, "base.HTTP: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if h == nil {
		h = &http.Client{
			Timeout: time.Duration(5) * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second * time.Duration(5),
			},
		}
	}
	return &httpClient{
		hLog:   hLog,
		client: h,
		debug:  debug,
	}
}

// do performs an HTTP request with the given method, URL, request body, URL parameters, and headers.
// It returns the HTTPResponse and an error if any.
func (h *httpClient) do(ctx context.Context, method, rURL string, reqBody interface{}, params url.Values, headers http.Header) (HTTPResponse, error) {
	if params == nil {
		params = url.Values{}
	}

	return h.doRaw(ctx, method, rURL, reqBody, params, headers)
}

// doRaw performs an HTTP request without JSON marshalling and unmarshalling.
// It handles request encoding, URL parameters, and headers.
// It returns the HTTPResponse and an error if any.
func (h *httpClient) doRaw(ctx context.Context, method, rURL string, reqBody interface{}, params url.Values, headers http.Header) (HTTPResponse, error) {
	var (
		resp     = HTTPResponse{}
		err      error
		postBody io.Reader
	)

	// Encode POST / PUT params.
	if (method == http.MethodPost || method == http.MethodPut) && reqBody != nil {
		aa, _ := json.Marshal(reqBody)
		postBody = bytes.NewReader(aa)
	}

	req, err := http.NewRequestWithContext(ctx, method, rURL, postBody)
	if err != nil {
		if h.debug {
			h.hLog.Printf("Request preparation failed: %v", err)
		}

		return resp, NewError(NetworkError, "Request preparation failed.", nil)
	}

	if headers != nil {
		req.Header = headers
	}

	// If a content-type isn't set, set the default one.
	if req.Header.Get("Content-Type") == "" {
		if method == http.MethodPost || method == http.MethodPut {
			req.Header.Add("Content-Type", "application/json")
		}
	}

	// If the request method is GET or DELETE, add the params as QueryString.
	if method == http.MethodGet || method == http.MethodDelete {
		req.URL.RawQuery = params.Encode()
	}

	r, err := h.client.Do(req)
	if err != nil {
		if h.debug {
			h.hLog.Printf("Request failed: %v", err)
		}

		return resp, NewError(NetworkError, "Request failed.", nil)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if h.debug {
			h.hLog.Printf("Unable to read response: %v", err)
		}

		return resp, NewError(DataError, "Error reading response.", nil)
	}

	resp.Response = r
	resp.Body = body
	if h.debug {
		h.hLog.Printf("%s %s -- %d %v", method, req.URL.RequestURI(), resp.Response.StatusCode, req.Header)
	}

	return resp, nil
}

// doJSON performs an HTTP request with JSON marshalling and unmarshalling.
// It calls the doRaw method to handle the request.
// It expects a valid JSON response and unmarshals it into the provided obj interface.
// It returns the HTTPResponse and an error if any.
func (h *httpClient) doJSON(ctx context.Context, method, rURL string, reqBody interface{}, params url.Values, headers http.Header, obj interface{}) (HTTPResponse, error) {
	resp, err := h.do(ctx, method, rURL, reqBody, params, headers)
	switch resp.Response.StatusCode {
	case http.StatusUnauthorized:
		return resp, NewError(PermissionError, "Unauthorized access", nil)
	}
	if err != nil {
		return resp, err
	}

	// We now unmarshal the body.
	if err := json.Unmarshal(resp.Body, obj); err != nil {
		if h.debug {
			h.hLog.Printf(string(resp.Body))
			h.hLog.Printf("Error parsing JSON response: %v | %s", err, resp.Body)
		}
		return resp, NewError(DataError, "Error parsing response.", nil)
	}
	if h.debug {
		h.hLog.Printf(string(resp.Body))
	}

	return resp, nil
}

// GetClient returns the underlying net/http client.
func (h *httpClient) GetClient() *httpClient {
	return h
}
