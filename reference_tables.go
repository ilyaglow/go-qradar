package qradar

import (
	"context"
	"net/http"
)

// ReferenceTableService handles methods related to Reference tables of the QRadar API.
type ReferenceTableService service

const (
	referenceTableServiceAPIPrefix = "api/reference_data/tables"
)

// ReferenceTable represents QRadar's Reference table.
type ReferenceTable struct {
	Name             *string `json:"name,omitempty"`
	CreationTime     *int    `json:"creation_time,omitempty"`
	ElementType      *string `json:"element_type,omitempty"`
	NumberOfElements *int    `json:"number_of_elements,omitempty"`
	TimeToLive       *string `json:"time_to_live,omitempty"`
	TimeoutType      *string `json:"timeout_type,omitempty"`
}

// Get returns Reference tables of the current QRadar installation.
func (c *ReferenceTableService) Get(ctx context.Context, fields, filter string, from, to int) ([]ReferenceTable, error) {
	req, err := c.client.requestHelp(http.MethodGet, referenceTableServiceAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []ReferenceTable
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates Reference table in QRadar installation.
func (c *ReferenceTableService) Create(ctx context.Context, fields string, data *ReferenceTable) (*ReferenceTable, error) {
	req, err := c.client.requestHelp(http.MethodPost, referenceTableServiceAPIPrefix, fields, "", 0, 0, nil, nil)
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

	var result ReferenceTable
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
