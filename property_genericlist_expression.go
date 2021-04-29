package qradar

import (
	"context"
	"net/http"
)

// PropertyGenericListExpressionService handles methods related to Property GenericList Expressions of the QRadar API.
type PropertyGenericListExpressionService service

const (
	propertyGenericListExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_genericlist_expressions"
)

// Get returns Property GenericList Expressions of the current QRadar installation
func (c *PropertyGenericListExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyGenericListExpressionAPIPrefix, fields, filter, from, to, nil, nil)
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

// GetByID returns Property GenericList Expression of the current QRadar installation by ID.
func (c *PropertyGenericListExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyGenericListExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

// Create creates Property GenericList Expression in QRadar installation.
func (c *PropertyGenericListExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyGenericListExpressionAPIPrefix, fields, "", 0, 0, nil, data)
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

// UpdateByID updates Property GenericList Expression in QRadar installation by ID.
func (c *PropertyGenericListExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyGenericListExpressionAPIPrefix, fields, "", 0, 0, &id, data)
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

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property GenericList Expression by ID.
func (c *PropertyGenericListExpressionService) DeleteByID(ctx context.Context, fields string, id int) error {
	req, err := c.client.requestHelp(http.MethodDelete, propertyGenericListExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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
