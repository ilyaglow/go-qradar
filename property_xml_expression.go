package qradar

import (
	"context"
	"net/http"
)

// PropertyXMLExpressionService handles methods related to Property XML Expressions of the QRadar API.
type PropertyXMLExpressionService service

const (
	propertyXMLExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_xml_expressions"
)

// Get returns Property XML Expressions of the current QRadar installation
func (c *PropertyXMLExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyXMLExpressionAPIPrefix, fields, filter, from, to, nil, nil)
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

// GetByID returns Property XML Expression of the current QRadar installation by ID.
func (c *PropertyXMLExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyXMLExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

// Create creates Property XML Expression in QRadar installation.
func (c *PropertyXMLExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyXMLExpressionAPIPrefix, fields, "", 0, 0, nil, data)
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

// UpdateByID updates Property XML Expression in QRadar installation by ID.
func (c *PropertyXMLExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyXMLExpressionAPIPrefix, fields, "", 0, 0, &id, data)
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

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property XML Expression by ID.
func (c *PropertyXMLExpressionService) DeleteByID(ctx context.Context, fields string, id int) error {
	req, err := c.client.requestHelp(http.MethodDelete, propertyXMLExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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
