package qradar

import (
	"context"
	"net/http"
)

// ReferenceSetService handles methods related to Reference sets of the QRadar API.
type ReferenceSetService service

const (
	referenceSetsServiceAPIPrefix = "api/reference_data/sets"
)

// ReferenceSet represents QRadar's Reference sets.
type ReferenceSet struct {
	Name             *string `json:"name,omitempty"`
	CreationTime     *int    `json:"creation_time,omitempty"`
	ElementType      *string `json:"element_type,omitempty"`
	NumberOfElements *int    `json:"number_of_elements,omitempty"`
	TimeToLive       *string `json:"time_to_live,omitempty"`
	TimeoutType      *string `json:"timeout_type,omitempty"`

	Data []ReferenceSetData `json:"data,omitempty"`
}

// ReferenceSetData represents data of Reference Set
type ReferenceSetData struct {
	FirstSeen *int    `json:"first_seen,omitempty"`
	LastSeen  *int    `json:"last_seen,omitempty"`
	Source    *string `json:"source,omitempty"`
	Value     *string `json:"value,omitempty"`
}

// Get returns Reference sets of the current QRadar installation.
func (c *ReferenceSetService) Get(ctx context.Context, fields, filter string, from, to int) ([]ReferenceSet, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceSetsServiceAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []ReferenceSet
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates Reference set in QRadar installation.
// expects pointer on a ReferenceSet
func (c *ReferenceSetService) Create(ctx context.Context, fields string, data *ReferenceSet) (*ReferenceSet, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceSetsServiceAPIPrefix, fields, "", 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()

	if data != nil {
		if data.Name != nil {
			q.Add("name", *data.Name)
		}
		if data.ElementType != nil {
			q.Add("element_type", *data.ElementType)
		}
		if data.TimeToLive != nil {
			q.Add("time_to_live", *data.TimeToLive)
		}
		if data.TimeoutType != nil {
			q.Add("timeout_type", *data.TimeoutType)
		}
	}

	req.URL.RawQuery = q.Encode()

	var result ReferenceSet
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWithData returns Reference set with data of the current QRadar installation.
func (c *ReferenceSetService) GetWithData(ctx context.Context, fields, filter, name string, from, to int) (*ReferenceSet, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceSetsServiceAPIPrefix+"/"+name, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result ReferenceSet
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkLoad uploads many values in QRadar's Reference Set
func (c *ReferenceSetService) BulkLoad(ctx context.Context, fields, name string, data interface{}) (*ReferenceSet, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceSetsServiceAPIPrefix+"/bulk_load/"+name, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result ReferenceSet
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
