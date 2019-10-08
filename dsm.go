package qradar

import (
	"context"
)

// DSMService handles methods related to DSMs of the QRadar API.
type DSMService service

const (
	dsmAPIPrefix = "api/data_classification/dsm_event_mappings"
)

// DSM represents QRadar's DSM
type DSM struct {
	ID                     *int    `json:"id,omitempty"`
	LogSourceTypeID        *int    `json:"log_source_type_id,omitempty"`
	LogSourceEventID       *string `json:"log_source_event_id,omitempty"`
	LogSourceEventCategory *string `json:"log_source_event_category,omitempty"`
	CustomEvent            *bool   `json:"custom_event,omitempty"`
	QIDRecordID            *int    `json:"qid_record_id,omitempty"`
	UUID                   *string `json:"uuid,omitempty"`
}

// Get returns DSMs of the current QRadar installation.
func (c *DSMService) Get(ctx context.Context, fields, filter string, from, to int) ([]DSM, error) {
	req, err := c.client.requestHelp("GET", dsmAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []DSM
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates DSM in the current QRadar installation.
func (c *DSMService) Create(ctx context.Context, fields string, data interface{}) (*DSM, error) {
	req, err := c.client.requestHelp("POST", dsmAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result DSM
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns DSM of the current QRadar installation by ID.
func (c *DSMService) GetByID(ctx context.Context, fields string, id int) (*DSM, error) {
	req, err := c.client.requestHelp("GET", dsmAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result DSM
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates DSM in QRadar installation by ID.
func (c *DSMService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*DSM, error) {
	req, err := c.client.requestHelp("POST", dsmAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result DSM
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
