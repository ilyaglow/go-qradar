package qradar

import (
	"context"
	"net/http"
)

// PropertyLEEFExpressionService handles methods related to Property LEEF Expressions of the QRadar API.
type PropertyLEEFExpressionService service

const (
	propertyLeefExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_leef_expressions"
)

// Get returns Property LEEF Expressions of the current QRadar installation.
func (c *PropertyLEEFExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyLeefExpressionAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []PropertyExpression
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns Property LEEF Expression of the current QRadar installation by ID.
func (c *PropertyLEEFExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyLeefExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result PropertyExpression
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates Property LEEF Expression in QRadar installation.
func (c *PropertyLEEFExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyLeefExpressionAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result PropertyExpression
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates Property LEEF Expression in QRadar installation by ID.
func (c *PropertyLEEFExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyLeefExpressionAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result PropertyExpression
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property LEEF Expression by ID.
func (c *PropertyLEEFExpressionService) DeleteByID(ctx context.Context, fields string, id int) error {
	req, err := c.client.requestHelp(http.MethodDelete, propertyLeefExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "text/plain")

	_, err = c.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}
