package qradar

import (
	"context"
	"net/http"
)

// PropertyExpressionService handles methods related to Property Expressions of the QRadar API.
type PropertyExpressionService service

// PropertyJSONExpressionService handles methods related to Property JSON Expressions of the QRadar API.
type PropertyJSONExpressionService service

// PropertyLEEFExpressionService handles methods related to Property LEEF Expressions of the QRadar API.
type PropertyLEEFExpressionService service

// PropertyCEFExpressionService handles methods related to Property CEF Expressions of the QRadar API.
type PropertyCEFExpressionService service

const (
	propertyExpressionAPIPrefix     = "api/config/event_sources/custom_properties/property_expressions"
	propertyJSONExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_json_expressions"
	propertyCefExpressionAPIPrefix  = "api/config/event_sources/custom_properties/property_cef_expressions"
	propertyLeefExpressionAPIPrefix = "api/config/event_sources/custom_properties/property_leef_expressions"
)

// PropertyExpression represents QRadar various property expressions which are regular expression, json, cef and leef.
// The structure for those would be the same with a distinction that regular expression would have field "Regex" and "CaptureGroup"
// whereas others structure have just "Expression" field instead.
type PropertyExpression struct {
	Identifier              *string `json:"identifier,omitempty"`
	LogSourceTypeID         *int    `json:"log_source_type_id,omitempty"`
	ModificationDate        *int    `json:"modification_date,omitempty"`
	QID                     *int    `json:"qid,omitempty"`
	LogSourceID             *int    `json:"log_source_id,omitempty"`
	Enabled                 *bool   `json:"enabled,omitempty"`
	Payload                 *string `json:"payload,omitempty"`
	RegexPropertyIdentifier *string `json:"regex_property_identifier,omitempty"`
	ID                      *int    `json:"id,omitempty"`
	CreationDate            *int    `json:"creation_date,omitempty"`
	Username                *string `json:"username,omitempty"`
	LowLevelCategoryID      *int    `json:"low_level_category_id,omitempty"`

	Regex        *string `json:"regex,omitempty"`
	CaptureGroup *int    `json:"capture_group,omitempty"`

	Expression *string `json:"expression,omitempty"`
}

// Get returns Property Expressions of the current QRadar installation.
func (c *PropertyExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyExpressionAPIPrefix, fields, filter, from, to, nil, nil)
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

// GetByID returns Property Expressions of the current QRadar installation by ID.
func (c *PropertyExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

// Create creates Property Expression in QRadar installation.
func (c *PropertyExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyExpressionAPIPrefix, fields, "", 0, 0, nil, data)
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

// UpdateByID updates Property Expression in QRadar installation by ID.
func (c *PropertyExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyExpressionAPIPrefix, fields, "", 0, 0, &id, data)
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

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property Expression by its id.
func (c *PropertyExpressionService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, propertyExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

//
// Property JSON Expression Block
//

// Get returns Property JSON Expressions of the current QRadar installation
func (c *PropertyJSONExpressionService) Get(ctx context.Context, fields, filter string, from, to int) ([]PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyJSONExpressionAPIPrefix, fields, filter, from, to, nil, nil)
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

// GetByID returns Property JSON Expression of the current QRadar installation by ID.
func (c *PropertyJSONExpressionService) GetByID(ctx context.Context, fields string, id int) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodGet, propertyJSONExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

// Create creates Property JSON Expression in QRadar installation>
func (c *PropertyJSONExpressionService) Create(ctx context.Context, fields string, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyJSONExpressionAPIPrefix, fields, "", 0, 0, nil, data)
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

// UpdateByID updates Property JSON Expression in QRadar installation by ID.
func (c *PropertyJSONExpressionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*PropertyExpression, error) {
	req, err := c.client.requestHelp(http.MethodPost, propertyJSONExpressionAPIPrefix, fields, "", 0, 0, &id, data)
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

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Property JSON Expression by ID.
func (c *PropertyJSONExpressionService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, propertyJSONExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

//
// Property LEEF Expression Block
//

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
func (c *PropertyLEEFExpressionService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, propertyLeefExpressionAPIPrefix, fields, "", 0, 0, &id, nil)
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

//
// Property CEF Expression Block
//

// Get returns Property CEF Expressions of the current QRadar installation
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
