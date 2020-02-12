package qradar

import (
	"context"
	"net/http"
)

// PropertyNVPExpressionService handles methods related to Property NVP Expressions of the QRadar API.
type PropertyNVPExpressionService service

const (
	propertyNvpExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_nvp_expressions"
)

// Get returns Property NVP Expressions of the current QRadar installation
func (c *PropertyNVPExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyNvpExpressionAPIPrefix, fields, filter, from, to, nil, nil)
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

// GetByID returns Property NVP Expression of the current QRadar installation by ID.
func (c *PropertyNVPExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyNvpExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

// Create creates Property NVP Expression in QRadar installation>
func (c *PropertyNVPExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyNvpExpressionAPIPrefix, fields, "", 0, 0, nil, data)
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

// UpdateByID updates Property NVP Expression in QRadar installation by ID.
func (c *PropertyNVPExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyNvpExpressionAPIPrefix, fields, "", 0, 0, &id, data)
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

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property NVP Expression by ID.
func (c *PropertyNVPExpressionService) DeleteByID(ctx context.Context, fields string, id int) error {
	req, err := c.client.requestHelp(http.MethodDelete, propertyNvpExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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
