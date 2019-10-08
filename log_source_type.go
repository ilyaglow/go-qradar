package qradar

import (
	"context"
)

// LogSourceTypeService handles methods related to Log Source Types of the QRadar API.
type LogSourceTypeService service

const (
	logSourceTypeAPIPrefix = "api/config/event_sources/log_source_management/log_source_types"
)

// LogSourceType represents QRadar's Log Source Type.
type LogSourceType struct {
	ID                   *int    `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Internal             *bool   `json:"internal,omitempty"`
	Custom               *bool   `json:"custom,omitempty"`
	DefaultProtocolID    *int    `json:"default_protocol_id,omitempty"`
	LogSourceExtensionID *int    `json:"log_source_extension_id,omitempty"`
	Version              *string `json:"version,omitempty"`
	SupportedLanguageIDs []int   `json:"supported_language_ids,omitempty"`

	ProtocolTypes []struct {
		ProtocolID *int  `json:"protocol_id,omitempty"`
		Documented *bool `json:"documented,omitempty"`
	} `json:"protocol_types,omitempty"`
}

// Get returns Log Source Types of the current QRadar installation.
func (c *LogSourceTypeService) Get(ctx context.Context, fields, filter string, from, to int) ([]LogSourceType, error) {
	req, err := c.client.requestHelp("GET", logSourceTypeAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []LogSourceType
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates Log Source Type in the current QRadar installation.
func (c *LogSourceTypeService) Create(ctx context.Context, fields string, data interface{}) (*LogSourceType, error) {
	req, err := c.client.requestHelp("POST", logSourceTypeAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result LogSourceType
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns Log Source Type of the current QRadar installation by ID.
func (c *LogSourceTypeService) GetByID(ctx context.Context, fields string, id int) (*LogSourceType, error) {
	req, err := c.client.requestHelp("GET", logSourceTypeAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result LogSourceType
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates Log Source Type in QRadar installation by ID.
func (c *LogSourceTypeService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*LogSourceType, error) {
	req, err := c.client.requestHelp("POST", logSourceTypeAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result LogSourceType
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Log Source Type by ID.
// TODO need to be tested
func (c *LogSourceTypeService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp("DELETE", logSourceTypeAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result DeleteTask
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
