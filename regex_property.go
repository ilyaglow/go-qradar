package qradar

import (
	"context"
	"fmt"
	"net/http"
)

// RegexPropertyService handles methods related to Regex Properties of the QRadar API.
type RegexPropertyService service

const (
	regexPropertyAPIPrefix = "api/config/event_sources/custom_properties/regex_properties"
)

// RegexProperty represents QRadar's Regex Property which is a metadata of a Custom Property.
type RegexProperty struct {
	Identifier       *string `json:"identifier,omitempty"`
	ModificationDate *int    `json:"modification_date,omitempty"`
	DatetimeFormat   *string `json:"datetime_format,omitempty"`
	PropertyType     *string `json:"property_type,omitempty"`
	Name             *string `json:"name,omitempty"`
	AutoDiscovered   *bool   `json:"auto_discovered,omitempty"`
	Description      *string `json:"description,omitempty"`
	ID               *int    `json:"id,omitempty"`
	UseForRuleEngine *bool   `json:"use_for_rule_engine,omitempty"`
	CreationDate     *int    `json:"creation_date,omitempty"`
	Locale           *string `json:"locale,omitempty"`
	Username         *string `json:"username,omitempty"`
}

// DeleteTask represents structure of a Delete Task to ensure safe deletion.
type DeleteTask struct {
	ID        *int    `json:"id,omitempty"`
	Message   *string `json:"message,omitempty"`
	Status    *string `json:"status,omitempty"`
	Name      *string `json:"name,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Created   *int    `json:"created,omitempty"`
	Started   *int    `json:"started,omitempty"`
	Modified  *int    `json:"modified,omitempty"`
	Completed *int    `json:"completed,omitempty"`
}

// Get returns Regex Properties of the current QRadar installation.
func (c *RegexPropertyService) Get(ctx context.Context, fields, filter string, from, to int) ([]RegexProperty, error) {
	req, err := c.client.requestHelp(http.MethodGet, regexPropertyAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []RegexProperty
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns Regex Property of the current QRadar installation by ID.
func (c *RegexPropertyService) GetByID(ctx context.Context, fields string, id int) (*RegexProperty, error) {
	req, err := c.client.requestHelp(http.MethodGet, regexPropertyAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result RegexProperty
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates Regex Property in QRadar installation.
func (c *RegexPropertyService) Create(ctx context.Context, fields string, data interface{}) (*RegexProperty, error) {
	req, err := c.client.requestHelp(http.MethodPost, regexPropertyAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result RegexProperty
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates Regex Property in QRadar installation by ID.
func (c *RegexPropertyService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*RegexProperty, error) {
	req, err := c.client.requestHelp(http.MethodPost, regexPropertyAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result RegexProperty
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Regex Property by ID.
func (c *RegexPropertyService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, regexPropertyAPIPrefix, fields, "", 0, 0, &id, nil)
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

// GetByName returns Regex Property of the current QRadar installation by Name.
func (c *RegexPropertyService) GetByName(ctx context.Context, fields string, name string) (*RegexProperty, error) {
	req, err := c.client.requestHelp(http.MethodGet, regexPropertyAPIPrefix, fields, fmt.Sprintf("name=\"%s\"", name), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []RegexProperty
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("found more rules than expected - %d", len(result))
	}
	return &result[0], nil
}

// GetByUUID returns Regex Property of the current QRadar installation by UUID.
func (c *RegexPropertyService) GetByUUID(ctx context.Context, fields string, uuid string) (*RegexProperty, error) {
	req, err := c.client.requestHelp(http.MethodGet, regexPropertyAPIPrefix, fields, fmt.Sprintf("identifier=\"%s\"", uuid), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []RegexProperty
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("found more rules than expected - %d", len(result))
	}
	return &result[0], nil
}
