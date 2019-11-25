package qradar

import (
	"context"
	"net/http"
)

// BuildingBlockService handles methods related to BuildingBlock of the QRadar API.
type BuildingBlockService service

const buildingBlockAPIPrefix = "api/analytics/building_blocks"

// BuildingBlock represents QRadar's BuildingBlock.
type BuildingBlock struct {
	ID                   *int    `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	BuildingBlockType    *string `json:"building_block_type,omitempty"`
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

// Get returns BuildingBlocks of the current QRadar installation
func (c *BuildingBlockService) Get(ctx context.Context, fields, filter string, from, to int) ([]BuildingBlock, error) {
	req, err := c.client.requestHelp(http.MethodGet, buildingBlockAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []BuildingBlock
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns BuildingBlock of the current QRadar installation by ID.
func (c *BuildingBlockService) GetByID(ctx context.Context, fields string, id int) (*BuildingBlock, error) {
	req, err := c.client.requestHelp(http.MethodGet, buildingBlockAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result BuildingBlock
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates only the BuildingBlock owner or enabled/disabled by ID.
func (c *BuildingBlockService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*BuildingBlock, error) {
	req, err := c.client.requestHelp(http.MethodPost, buildingBlockAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result BuildingBlock
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID creates A Delete Task in QRadar installation in order to safely delete Rule by ID.
func (c *BuildingBlockService) DeleteByID(ctx context.Context, fields string, id int) (*DeleteTask, error) {
	req, err := c.client.requestHelp(http.MethodDelete, buildingBlockAPIPrefix, fields, "", 0, 0, &id, nil)
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
