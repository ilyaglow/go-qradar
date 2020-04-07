package qradar

import (
	"context"
	"net/http"
)

// ReferenceMapService handles methods related to Reference Map of the QRadar API.
type ReferenceMapService service

const (
	referenceMapServiceAPIPrefix = "api/reference_data/maps"
)

// ReferenceMap represents QRadar's Reference Map.
type ReferenceMap struct {
	Name             *string `json:"name,omitempty"`
	CreationTime     *int    `json:"creation_time,omitempty"`
	ElementType      *string `json:"element_type,omitempty"`
	KeyLabel         *string `json:"key_label,omitempty"`
	NumberOfElements *int    `json:"number_of_elements,omitempty"`
	TimeToLive       *string `json:"time_to_live,omitempty"`
	TimeoutType      *string `json:"timeout_type,omitempty"`
	ValueLabel       *string `json:"value_label,omitempty"`

	Data map[string]ReferenceSetData `json:"data,omitempty"`
}

// Get returns Reference maps of the current QRadar installation.
func (c *ReferenceMapService) Get(ctx context.Context, fields, filter string, from, to int) ([]ReferenceMap, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceMapServiceAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []ReferenceMap
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates Reference map in QRadar installation.
func (c *ReferenceMapService) Create(ctx context.Context, fields string, data *ReferenceMap) (*ReferenceMap, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceMapServiceAPIPrefix, fields, "", 0, 0, nil, nil)
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

	var result ReferenceMap
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWithData returns Reference Map with data of the current QRadar installation.
func (c *ReferenceMapService) GetWithData(ctx context.Context, fields, filter, name string, from, to int) (*ReferenceMap, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceMapServiceAPIPrefix+"/"+name, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result ReferenceMap
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkLoad uploads many values in QRadar's Reference Map
func (c *ReferenceMapService) BulkLoad(ctx context.Context, fields, name string, data interface{}) (*ReferenceMap, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceMapServiceAPIPrefix+"/bulk_load/"+name, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result ReferenceMap
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
