package qradar

import (
	"context"
	"fmt"
	"net/http"
)

// SearchResultsWindow is a default window for scrolling results of the query.
var SearchResultsWindow = 50

// SearchResultsScroller represents a scroller for the results of the query.
type SearchResultsScroller struct {
	count             int
	client            *Client
	searchID          string
	startIdx, currIdx int
	window            int
	events            []Event
}

// NewSearchResultsScroller initializes struct to scroll the records.
func (a *ArielService) NewSearchResultsScroller(ctx context.Context, searchID string) (*SearchResultsScroller, error) {
	_, num, err := a.SearchStatus(ctx, searchID)
	if err != nil {
		return nil, err
	}

	srs := &SearchResultsScroller{
		count:    num,
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

// Next returns true if an event is still available to be consumed by the
// Result() method.
func (s *SearchResultsScroller) Next(ctx context.Context) bool {
	if len(s.events) < s.window && s.currIdx-s.startIdx == len(s.events) {
		return false
	}

	if s.currIdx-s.startIdx == len(s.events) && len(s.events) == s.window {
		s.startIdx += s.window
		err := s.getEvents(ctx)
		if err != nil {
			return false
		}
	}

	return true
}

func (s *SearchResultsScroller) getEvents(ctx context.Context) error {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/results", arielSearchAPIPrefix, s.searchID), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", fmt.Sprintf("items=%d-%d", s.startIdx, s.startIdx+s.window))

	var r SearchResult
	_, err = s.client.Do(ctx, req, &r)
	if err != nil {
		return err
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

// Length returns the overall events count.
func (s *SearchResultsScroller) Length() int {
	return s.count
}

// ScrollByQuery events in the QRadar API.
// Recommended way to retrieve large amount of events.
func (a *ArielService) ScrollByQuery(ctx context.Context, sqlQuery string) (*SearchResultsScroller, *SearchMetadata, error) {
	s, err := a.SearchByQuery(ctx, sqlQuery)
	if err != nil {
		return nil, nil, err
	}

	_, err = a.WaitForSearchID(ctx, *s.SearchID, StatusCompleted, 2)
	if err != nil {
		return nil, nil, err
	}

	meta, err := a.SearchMetadata(ctx, *s.SearchID)
	if err != nil {
		return nil, nil, err
	}

	srs, err := a.NewSearchResultsScroller(ctx, *s.SearchID)
	if err != nil {
		return nil, meta, err
	}

	return srs, meta, nil
}
