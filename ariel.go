package qradar

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const arielSearchAPIPrefix = "api/ariel/searches"

// SearchResultsWindow is a default window for scrolling results of the query.
var SearchResultsWindow = 50

// ArielService handles communication with the search-related methods of
// the QRadar API.
type ArielService service

// ErrorMessage represents generic error message by the QRadar API.
type ErrorMessage struct {
	Code     string   `json:"code"`
	Contexts []string `json:"contexts"`
	Message  string   `json:"message"`
	Severity string   `json:"severity"`
}

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

// SearchResultsScroller represents a scroller for the results of the query.
type SearchResultsScroller struct {
	Count             int
	client            *Client
	searchID          string
	startIdx, currIdx int
	window            int
	events            []Event
}

// NewSearchResultsScroller initializes struct to scroll the records.
func (a *ArielService) NewSearchResultsScroller(ctx context.Context, searchID string) (*SearchResultsScroller, error) {
	_, num, err := a.SearchStatus(ctx, searchID)

	srs := &SearchResultsScroller{
		Count:    num,
		window:   SearchResultsWindow,
		client:   a.client,
		searchID: searchID,
	}

	err = srs.getEvents(ctx)
	if err != nil {
		return nil, err
	}

	return srs, nil
}

// Next return the next result from the search.
func (s *SearchResultsScroller) Next(ctx context.Context) bool {
	if s.currIdx-s.startIdx == len(s.events) && len(s.events) < s.window {
		return false
	}

	if s.currIdx-s.startIdx > len(s.events) && len(s.events) == s.window {
		s.startIdx += s.window
		err := s.getEvents(ctx)
		if err != nil {
			return false
		}
	}

	return true
}

func (s *SearchResultsScroller) getEvents(ctx context.Context) error {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/%s/results", arielSearchAPIPrefix, s.searchID), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", fmt.Sprintf("items=%d-%d", s.startIdx, s.startIdx+s.window))

	var r SearchResult
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SearchResultScroller failed: code %s", resp.Status)
	}

	s.events = r.Events

	return nil
}

// Result returns the event iterated by the Next.
func (s *SearchResultsScroller) Result() Event {
	event := s.events[s.currIdx-s.startIdx]
	s.currIdx++
	return event

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

func (a *ArielService) searchByQuery(ctx context.Context, sqlQuery string) (*Search, error) {
	req, err := a.client.NewRequest("POST", arielSearchAPIPrefix, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/plain")
	q := req.URL.Query()
	q.Add("query_expression", sqlQuery)
	req.URL.RawQuery = q.Encode()

	var s Search
	resp, err := a.client.Do(ctx, req, &s)
	if err != nil {
		return nil, err
	}

	// TODO: handle [1:] error messages
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(
			"SearchByQuery: code %s, message: %s",
			s.ErrorMessages[0].Code,
			s.ErrorMessages[0].Message,
		)
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
	resp, err := a.client.Do(ctx, req, &s)
	if err != nil {
		return "", 0, err
	}

	// TODO: handle [1:] error messages
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusPartialContent) {
		return "", 0, fmt.Errorf(
			"SearchByQuery: code %s, message: %s",
			s.ErrorMessages[0].Code,
			s.ErrorMessages[0].Message,
		)
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
	resp, err := a.client.Do(ctx, req, &s)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SearchMetadata failed: code %s", resp.Status)
	}

	return &s, nil
}

// waitForSearchID returns amount of records and the error.
func (a *ArielService) waitForSearchID(ctx context.Context, searchID string, status JobStatus, seconds int) (int, error) {
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

// SearchByQuery searches events in the QRadar API by SQL query.
func (a *ArielService) SearchByQuery(ctx context.Context, sqlQuery string) (*SearchResultsScroller, *SearchMetadata, error) {
	s, err := a.searchByQuery(ctx, sqlQuery)
	if err != nil {
		return nil, nil, err
	}

	_, err = a.waitForSearchID(ctx, s.SearchID, StatusCompleted, 2)
	if err != nil {
		return nil, nil, err
	}

	meta, err := a.SearchMetadata(ctx, s.SearchID)
	if err != nil {
		return nil, nil, err
	}

	srs, err := a.NewSearchResultsScroller(ctx, s.SearchID)
	if err != nil {
		return nil, meta, err
	}

	return srs, meta, nil
}
