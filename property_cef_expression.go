package qradar

import (
	"context"
	"net/http"
)

// PropertyCEFExpressionService handles methods related to Property CEF Expressions of the QRadar API.
type PropertyCEFExpressionService service

const (
	propertyCefExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_cef_expressions"
)

// Get returns Property CEF Expressions of the current QRadar installation.
func (c *PropertyCEFExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyCefExpressionAPIPrefix, fields, filter, from, to, nil, nil)
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

// GetByID returns Property CEF Expression of the current QRadar installation by ID.
func (c *PropertyCEFExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyCefExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

// Create creates Property CEF Expression in QRadar installation.
func (c *PropertyCEFExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyCefExpressionAPIPrefix, fields, "", 0, 0, nil, data)
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

// UpdateByID updates Property CEF Expression in QRadar installation by ID.
func (c *PropertyCEFExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyCefExpressionAPIPrefix, fields, "", 0, 0, &id, data)
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

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property CEF Expression by ID.
func (c *PropertyCEFExpressionService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, propertyCefExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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
