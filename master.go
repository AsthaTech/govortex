package govortex

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

// DownloadMaster retrieves the master data from the Vortex API.
// It returns a slice of maps representing the CSV records and an error if any.
func (v *VortexApi) DownloadMaster(ctx context.Context) ([]map[string]string, error) {
	endpoint := "/data/instruments"
	bearerToken := fmt.Sprintf("Bearer %s", v.AccessToken)
	headers := make(http.Header, 0)
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", bearerToken)
	endpointUrl := v.baseURL + endpoint

	req, err := http.NewRequestWithContext(ctx, "GET", endpointUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header = headers
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
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV record: %v", err)
		}
		if len(columns) == 0 {
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
