package qradar

import (
	"context"
	"fmt"
	"net/http"
)

// RuleService handles methods related to Rule of the QRadar API.
type RuleService service

const ruleAPIPrefix = "api/analytics/rules"

// Rule represents QRadar's Rule.
type Rule struct {
	ID                   *int    `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty"`
	Owner                *string `json:"owner,omitempty"`
	Origin               *string `json:"origin,omitempty"`
	BaseCapacity         *int    `json:"base_capacity,omitempty"`
	BaseHostID           *int    `json:"base_host_id,omitempty"`
	AverageCapacity      *int    `json:"average_capacity,omitempty"`
	CapacityTimestamp    *int    `json:"capacity_timestamp,omitempty"`
	Identifier           *string `json:"identifier,omitempty"`
	LinkedRuleIdentifier *string `json:"linked_rule_identifier,omitempty"`
	CreationDate         *int    `json:"creation_date,omitempty"`
	ModificationDate     *int    `json:"modification_date,omitempty"`
}

// Get returns Rules of the current QRadar installation.
func (c *RuleService) Get(ctx context.Context, fields, filter string, from, to int) ([]Rule, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []Rule
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns Rule of the current QRadar installation by ID.
func (c *RuleService) GetByID(ctx context.Context, fields string, id int) (*Rule, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result Rule
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates the rule owner or toggle the rule enabled/disabled by ID.
func (c *RuleService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*Rule, error) {
	req, err := c.client.requestHelp(http.MethodPost, ruleAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result Rule
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Rule by ID.
func (c *RuleService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, ruleAPIPrefix, fields, "", 0, 0, &id, nil)
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

// GetByName returns Rule of the current QRadar installation by Name.
func (c *RuleService) GetByName(ctx context.Context, fields string, name string) (*Rule, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleAPIPrefix, fields, fmt.Sprintf("name=\"%s\"", name), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []Rule
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

// GetByUUID returns Rule of the current QRadar installation by UUID.
func (c *RuleService) GetByUUID(ctx context.Context, fields string, uuid string) (*Rule, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleAPIPrefix, fields, fmt.Sprintf("identifier=\"%s\"", uuid), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []Rule
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
