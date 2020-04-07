package qradar

import (
	"context"
	"net/http"
)

// ReferenceMapOfSetsService handles methods related to Reference Maps of Sets of the QRadar API.
type ReferenceMapOfSetsService service

const (
	referenceMapOfSetsServiceAPIPrefix = "api/reference_data/map_of_sets"
)

// ReferenceMapOfSets represents QRadar's Reference maps of sets.
type ReferenceMapOfSets struct {
	Name             *string `json:"name,omitempty"`
	CreationTime     *int    `json:"creation_time,omitempty"`
	ElementType      *string `json:"element_type,omitempty"`
	KeyLabel         *string `json:"key_label,omitempty"`
	NumberOfElements *int    `json:"number_of_elements,omitempty"`
	TimeToLive       *string `json:"time_to_live,omitempty"`
	TimeoutType      *string `json:"timeout_type,omitempty"`
	ValueLabel       *string `json:"value_label,omitempty"`

	Data map[string][]ReferenceSetData `json:"data,omitempty"`
}

// Get returns Reference maps of sets of the current QRadar installation.
func (c *ReferenceMapOfSetsService) Get(ctx context.Context, fields, filter string, from, to int) ([]ReferenceMapOfSets, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceMapOfSetsServiceAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []ReferenceMapOfSets
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates Reference maps of sets in QRadar installation.
func (c *ReferenceMapOfSetsService) Create(ctx context.Context, fields string, data *ReferenceMapOfSets) (*ReferenceMapOfSets, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceMapOfSetsServiceAPIPrefix, fields, "", 0, 0, nil, nil)
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

	var result ReferenceMapOfSets
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWithData returns Reference Map of Sets with data of the current QRadar installation.
func (c *ReferenceMapOfSetsService) GetWithData(ctx context.Context, fields, filter, name string, from, to int) (*ReferenceMapOfSets, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceMapOfSetsServiceAPIPrefix+"/"+name, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result ReferenceMapOfSets
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkLoad uploads many values in QRadar's Reference Map o Sets
func (c *ReferenceMapOfSetsService) BulkLoad(ctx context.Context, fields, name string, data interface{}) (*ReferenceMapOfSets, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceMapOfSetsServiceAPIPrefix+"/bulk_load/"+name, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result ReferenceMapOfSets
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
