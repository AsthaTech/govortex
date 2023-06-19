package govortex

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DownloadMaster retrieves the master data from the Vortex API.
// It returns a slice of maps representing the CSV records and an error if any.
func (v *VortexApi) DownloadMaster(ctx context.Context) ([]map[string]string, error) {
	endpoint := "/data/instruments"
	bearerToken := fmt.Sprintf("Bearer %s", v.AccessToken)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": bearerToken,
	}
	endpointUrl := v.baseURL + endpoint
	queryParams := url.Values{}
	for key, value := range headers {
		queryParams.Add(key, fmt.Sprintf("%v", value))
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpointUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.URL.RawQuery = queryParams.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.TrimLeadingSpace = true

	results := make([]map[string]string, 0)
	columns := make([]string, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV record: %v", err)
		}

		if columns == nil {
			columns = record
		} else {
			row := make(map[string]string)
			for i, value := range record {
				row[columns[i]] = value
			}
			results = append(results, row)
		}
	}

	return results, nil

}
