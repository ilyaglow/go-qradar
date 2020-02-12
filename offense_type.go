package qradar

import (
	"context"
	"net/http"
)

// OffenseTypeService handles methods related to OffenseTypes of the QRadar API.
type OffenseTypeService service

const offenseTypeAPIPrefix = "api/siem/offense_types"

// OffenseType represents QRadar's generated OffenseType.
type OffenseType struct {
	ID           *int    `json:"id,omitempty"`
	PropertyName *string `json:"property_name,omitempty"`
	Name         *string `json:"name,omitempty"`
	DatabaseType *string `json:"database_type,omitempty"`
	Custom       *bool   `json:"custom,omitempty"`
}

// Get returns OffenseTypes of the current QRadar installation.
func (c *OffenseTypeService) Get(ctx context.Context, fields, filter string, from, to int) ([]OffenseType, error) {
	req, err := c.client.requestHelp(http.MethodGet, offenseTypeAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []OffenseType
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns OffenseType of the current QRadar installation by ID.
func (c *OffenseTypeService) GetByID(ctx context.Context, fields string, id int) (*OffenseType, error) {
	req, err := c.client.requestHelp(http.MethodGet, offenseTypeAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result OffenseType
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
