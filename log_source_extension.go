package qradar

import (
	"context"
	"fmt"
	"net/http"
)

// LogSourceExtensionService handles methods related to Log Source Extensions of the QRadar Undocumented API.
type LogSourceExtensionService service

const (
	logSourceExtensionAPIPrefix = "api/config/event_sources/log_source_extensions"
)

// LogSourceExtension represents QRadar's Log Source Extension.
type LogSourceExtension struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Enabled      *bool   `json:"enabled,omitempty"`
	UseCondition *int    `json:"use_condition,omitempty"`
	XML          *string `json:"xml,omitempty"`
}

// Get returns Log Source Extension of the current QRadar installation. Undocumented API.
func (c *LogSourceExtensionService) Get(ctx context.Context, fields, filter string, from, to int) ([]LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceExtensionAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// todo: add support of multipart uploading
/*
	// Create creates Log Source Extension in the current QRadar installation. Undocumented API.
	func (c *LogSourceExtensionService) Create(ctx context.Context, fields string, data interface{}) (*LogSourceExtension, error) {
		req, err := c.client.requestHelp(http.MethodPost, logSourceExtensionAPIPrefix, fields, "", 0, 0, nil, data)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Allow-Hidden", "true")
		req.Header.Set("Content-Type", "multipart/form-data")

		var result LogSourceExtension
		_, err = c.client.Do(ctx, req, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
*/

// GetByID returns Log Source Extension of the current QRadar installation by ID. Undocumented API.
func (c *LogSourceExtensionService) GetByID(ctx context.Context, fields string, id int) (*LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceExtensionAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// todo: add support of multipart uploading
/*
	// UpdateByID updates Log Source Extension of the current QRadar installation by ID. Undocumented API.
	func (c *LogSourceExtensionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*LogSourceExtension, error) {
		req, err := c.client.requestHelp(http.MethodPost, logSourceExtensionAPIPrefix, fields, "", 0, 0, &id, data)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Allow-Hidden", "true")

		var result LogSourceExtension
		_, err = c.client.Do(ctx, req, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
*/

// GetByName returns Log Source Extension of the current QRadar installation by Name. Undocumented API.
func (c *LogSourceExtensionService) GetByName(ctx context.Context, fields string, name string) (*LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceExtensionAPIPrefix, fields, fmt.Sprintf("name=\"%s\"", name), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	} else if len(result) > 1 {
		return nil, fmt.Errorf("found more log_source_extensions than expected - %d", len(result))
	}
	return &result[0], nil
}
