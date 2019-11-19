package qradar

import (
	"context"
	"net/http"
)

// QIDService handles methods related to QIDs of the QRadar API.
type QIDService service

const (
	qidAPIPrefix = "api/data_classification/qid_records"
)

// QID represents QRadar's QID.
type QID struct {
	Severity         *int    `json:"severity,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	LogSourceTypeID  *int    `json:"log_source_type_id,omitempty"`
	ID               *int    `json:"id,omitempty"`
	LowLevelCategory *int    `json:"low_level_category_id,omitempty"`
	Qid              *int    `json:"qid"`
	UUID             *string `json:"uuid,omitempty"`
}

// Get returns QIDs of the current QRadar installation.
func (c *QIDService) Get(ctx context.Context, fields, filter string, from, to int) ([]QID, error) {
	req, err := c.client.requestHelp(http.MethodGet, qidAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []QID
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns QID of the current QRadar installation by ID.
func (c *QIDService) GetByID(ctx context.Context, fields string, id int) (*QID, error) {
	req, err := c.client.requestHelp(http.MethodGet, qidAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result QID
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates QID in QRadar installation.
func (c *QIDService) Create(ctx context.Context, fields string, data interface{}) (*QID, error) {
	req, err := c.client.requestHelp(http.MethodPost, qidAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result QID
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates QID record in QRadar installation bu ID.
func (c *QIDService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*QID, error) {
	req, err := c.client.requestHelp(http.MethodPost, qidAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result QID
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
