package qradar

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const arielSearchAPIPrefix = "api/ariel/searches"

// ArielService handles communication with the search-related methods of
// the QRadar API.
type ArielService service

// Search represent Ariel search state.
type Search struct {
	CursorID                 *string        `json:"cursor_id,omitempty"`
	CompressedDataFileCount  *int           `json:"compressed_data_file_count,omitempty"`
	CompressedDataTotalSize  *int           `json:"compressed_data_total_size,omitempty"`
	DataFileCount            *int           `json:"data_file_count,omitempty"`
	DataTotalSize            *int           `json:"data_total_size,omitempty"`
	IndexFileCount           *int           `json:"index_file_count,omitempty"`
	IndexTotalSize           *int           `json:"index_total_size,omitempty"`
	ProcessedRecordCount     *int           `json:"processed_record_count,omitempty"`
	ErrorMessages            []ErrorMessage `json:"error_messages,omitempty"`
	DesiredRetentionTimeMsec *int           `json:"desired_retention_time_msec,omitempty"`
	Progress                 *int           `json:"progress,omitempty"`
	ProgressDetails          []int          `json:"progress_details,omitempty"`
	QueryExecutionTime       *int           `json:"query_execution_time,omitempty"`
	QueryString              *string        `json:"query_string,omitempty"`
	RecordCount              *int           `json:"record_count,omitempty"`
	SaveResults              *bool          `json:"save_results,omitempty"`
	Status                   *string        `json:"status,omitempty"`
	Snapshot                 *struct {
		Events []Event `json:"events,omitempty"`
	} `json:"snapshot,omitempty"`
	SubsearchIds []string `json:"subsearch_ids,omitempty"`
	SearchID     *string  `json:"search_id,omitempty"`
}

// Event represents generic event result.
type Event map[string]interface{}

// SearchResult represents search result.
type SearchResult struct {
	Events []Event `json:"events,omitempty"`
}

// SearchMetadata represents search metadata.
type SearchMetadata struct {
	Columns []SearchColumn `json:"columns,omitempty"`
}

// SearchColumn represents found column and it's properties.
type SearchColumn struct {
	ArgumentType    *string `json:"argument_type,omitempty"`
	Indexable       *bool   `json:"indexable,omitempty"`
	Name            *string `json:"name,omitempty"`
	Nullable        *bool   `json:"nullable,omitempty"`
	ObjectValueType *string `json:"object_value_type,omitempty"`
	ProviderName    *string `json:"provider_name,omitempty"`
}

// SearchByQuery events in the QRadar API.
// It's caller responsibility to wait for results and get the final data.
func (a *ArielService) SearchByQuery(ctx context.Context, sqlQuery string) (*Search, error) {
	req, err := a.client.NewRequest(http.MethodPost, arielSearchAPIPrefix, nil)
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
	req, err := a.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", arielSearchAPIPrefix, searchID), nil)
	if err != nil {
		return "", 0, err
	}

	var s Search
	_, err = a.client.Do(ctx, req, &s)
	if err != nil {
		return "", 0, err
	}

	return *s.Status, *s.RecordCount, nil
}

// SearchMetadata represents a metadata retriever.
func (a *ArielService) SearchMetadata(ctx context.Context, searchID string) (*SearchMetadata, error) {
	req, err := a.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/metadata", arielSearchAPIPrefix, searchID), nil)
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

// DeleteSearch returns a search status that has been deleted and the error.
func (a *ArielService) DeleteSearch(ctx context.Context, searchID string) (string, error) {
	req, err := a.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", arielSearchAPIPrefix, searchID), nil)
	if err != nil {
		return "", err
	}

	var s Search
	_, err = a.client.Do(ctx, req, &s)
	if err != nil {
		return "", err
	}

	return *s.Status, nil
}