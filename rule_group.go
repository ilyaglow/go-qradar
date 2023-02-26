package qradar

import (
	"context"
	"net/http"
)

// RuleGroupService handles methods related to Rule Groups of the QRadar API.
type RuleGroupService service

const (
	ruleGroupAPIPrefix = "api/analytics/rule_groups"
)

// RuleGroup represents QRadar's Rule Group.

type RuleGroup struct {
	Owner        *string  `json:"owner"`
	ModifiedTime *int     `json:"modified_time"`
	Level        *int     `json:"level"`
	Name         *string  `json:"name"`
	Description  *string  `json:"description"`
	ChildGroups  []int    `json:"child_groups"`
	ID           *int     `json:"id"`
	ChildItems   []string `json:"child_items"`
	Type         *string  `json:"type"`
	ParentID     *int     `json:"parent_id"`
}

// Get returns Rule Groups of the current QRadar installation.
func (c *RuleGroupService) Get(ctx context.Context, fields, filter string, from, to int) ([]RuleGroup, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleGroupAPIPrefix, fields, filter, from, to, nil, nil)

	if err != nil {
		return nil, err
	}
	var result []RuleGroup
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns Rule Group of the current QRadar installation by ID.
func (c *RuleGroupService) GetByID(ctx context.Context, fields string, id int) (*RuleGroup, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleGroupAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result RuleGroup
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
