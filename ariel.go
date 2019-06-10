package qradar

import (
	"context"
	"fmt"
	"time"
)

const arielSearchAPIPrefix = "api/ariel/searches"

// ArielService handles communication with the search-related methods of
// the QRadar API.
type ArielService service

// Search represent Ariel search state.
type Search struct {
	CursorID                 string         `json:"cursor_id"`
	CompressedDataFileCount  int            `json:"compressed_data_file_count"`
	CompressedDataTotalSize  int            `json:"compressed_data_total_size"`
	DataFileCount            int            `json:"data_file_count"`
	DataTotalSize            int            `json:"data_total_size"`
	IndexFileCount           int            `json:"index_file_count"`
	IndexTotalSize           int            `json:"index_total_size"`
	ProcessedRecordCount     int            `json:"processed_record_count"`
	ErrorMessages            []ErrorMessage `json:"error_messages"`
	DesiredRetentionTimeMsec int            `json:"desired_retention_time_msec"`
	Progress                 int            `json:"progress"`
	ProgressDetails          []int          `json:"progress_details"`
	QueryExecutionTime       int            `json:"query_execution_time"`
	QueryString              string         `json:"query_string"`
	RecordCount              int            `json:"record_count"`
	SaveResults              bool           `json:"save_results"`
	Status                   string         `json:"status"`
	Snapshot                 struct {
		Events []Event `json:"events"`
	} `json:"snapshot"`
	SubsearchIds []string `json:"subsearch_ids"`
	SearchID     string   `json:"search_id"`
}

// Event represents generic event result.
type Event map[string]interface{}

// SearchResult represents search result.
type SearchResult struct {
	Events []Event `json:"events"`
}

// SearchMetadata represents search metadata.
type SearchMetadata struct {
	Columns []SearchColumn `json:"columns"`
}

// SearchColumn represents found column and it's properties.
type SearchColumn struct {
	ArgumentType    string `json:"argument_type"`
	Indexable       bool   `json:"indexable"`
	Name            string `json:"name"`
	Nullable        bool   `json:"nullable"`
	ObjectValueType string `json:"object_value_type"`
	ProviderName    string `json:"provider_name"`
}

// SearchByQuery events in the QRadar API.
// It's caller responsibility to wait for results and get the final data.
func (a *ArielService) SearchByQuery(ctx context.Context, sqlQuery string) (*Search, error) {
	req, err := a.client.NewRequest("POST", arielSearchAPIPrefix, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/plain")
	q := req.URL.Query()
	q.Add("query_expression", sqlQuery)
	req.URL.RawQuery = q.Encode()

	var s Search
	_, err = a.client.Do(ctx, req, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// SearchStatus returns a status and count of the records of the search.
func (a *ArielService) SearchStatus(ctx context.Context, searchID string) (string, int, error) {
	req, err := a.client.NewRequest("GET", fmt.Sprintf("%s/%s", arielSearchAPIPrefix, searchID), nil)
	if err != nil {
		return "", 0, err
	}

	var s Search
	_, err = a.client.Do(ctx, req, &s)
	if err != nil {
		return "", 0, err
	}

	return s.Status, s.RecordCount, nil
}

// SearchMetadata represents a metadata retriever.
func (a *ArielService) SearchMetadata(ctx context.Context, searchID string) (*SearchMetadata, error) {
	req, err := a.client.NewRequest("GET", fmt.Sprintf("%s/%s/metadata", arielSearchAPIPrefix, searchID), nil)
	if err != nil {
		return nil, err
	}

	var s SearchMetadata
	_, err = a.client.Do(ctx, req, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// WaitForSearchID returns amount of records and the error.
func (a *ArielService) WaitForSearchID(ctx context.Context, searchID string, status JobStatus, seconds int) (int, error) {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-ticker.C:
			s, num, err := a.SearchStatus(ctx, searchID)
			if err != nil {
				ticker.Stop()
				return 0, err
			}

			if (JobStatus)(s) == status {
				ticker.Stop()
				return num, nil
			}
		}
	}
}
