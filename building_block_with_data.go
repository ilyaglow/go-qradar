package qradar

import (
	"context"
	"net/http"
)

// BuildingBlockWithDataService handles methods related to BuildingBlock of the QRadar API.
type BuildingBlockWithDataService service

const buildingBlockWithDataAPIPrefix = "api/analytics/building_blocks_with_data"

// BuildingBlockWithData represents QRadar's BuildingBlock.
type BuildingBlockWithData struct {
	BuildingBlock
	RuleXML *string `json:"rule_xml,omitempty"`
}

// Get returns BuildingBlockWithData of the current QRadar installation. Undocumented.
func (c *BuildingBlockWithDataService) Get(ctx context.Context, fields, filter string, from, to int) ([]BuildingBlockWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, buildingBlockWithDataAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []BuildingBlockWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates BuildingBlockWithData in the current QRadar installation. Undocumented.
func (c *BuildingBlockWithDataService) Create(ctx context.Context, fields string, data interface{}) (*BuildingBlockWithData, error) {
	req, err := c.client.requestHelp(http.MethodPost, ruleWithDataAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result BuildingBlockWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns BuildingBlockWithData of the current QRadar installation by ID. Undocumented.
func (c *BuildingBlockWithDataService) GetByID(ctx context.Context, fields string, id int) (*BuildingBlockWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, buildingBlockWithDataAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result BuildingBlockWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates BuildingBlockWithData by ID. Undocumented.
func (c *BuildingBlockWithDataService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*BuildingBlockWithData, error) {
	req, err := c.client.requestHelp(http.MethodPost, buildingBlockWithDataAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result BuildingBlockWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
