package qradar

import (
	"context"
	"fmt"
	"net/http"
)

// RuleWithDataService handles methods related to RuleWithData of the QRadar Undocumented API.
type RuleWithDataService service

const ruleWithDataAPIPrefix = "api/analytics/rules_with_data"

// RuleWithData represents QRadar's RuleWithData. Undocumented.
type RuleWithData struct {
	Rule
	CRE             *int    `json:"cre,omitempty"`
	TypeID          *int    `json:"type_id,omitempty"`
	IsBuildingBlock *bool   `json:"is_building_block,omitempty"`
	RuleXML         *string `json:"rule_xml,omitempty"`
}

// Get returns RuleWithData of the current QRadar installation. Undocumented API.
func (c *RuleWithDataService) Get(ctx context.Context, fields, filter string, from, to int) ([]RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleWithDataAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates RuleWithData in the current QRadar installation. Undocumented API.
func (c *RuleWithDataService) Create(ctx context.Context, fields string, data interface{}) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodPost, ruleWithDataAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns RuleWithData of the current QRadar installation by ID. Undocumented API.
func (c *RuleWithDataService) GetByID(ctx context.Context, fields string, id int) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleWithDataAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates RuleWithData of the current QRadar installation by ID. Undocumented API.
func (c *RuleWithDataService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodPost, ruleWithDataAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByName returns RuleWithData of the current QRadar installation by Name. Undocumented API.
func (c *RuleWithDataService) GetByName(ctx context.Context, fields string, name string) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleWithDataAPIPrefix, fields, fmt.Sprintf("name=\"%s\"", name), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []RuleWithData
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

// GetByUUID returns RuleWithData of the current QRadar installation by UUID. Undocumented API.
func (c *RuleWithDataService) GetByUUID(ctx context.Context, fields string, uuid string) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleWithDataAPIPrefix, fields, fmt.Sprintf("identifier=\"%s\"", uuid), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []RuleWithData
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
