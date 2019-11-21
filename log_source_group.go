package qradar

import (
	"context"
	"net/http"
)

// LogSourceGroupService handles methods related to Log Source Groups of the QRadar API.
type LogSourceGroupService service

const (
	logSourceGroupAPIPrefix = "api/config/event_sources/log_source_management/log_source_groups"
)

// LogSourceGroup represents QRadar's Log Source Group.
type LogSourceGroup struct {
	ID               *int    `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	ParentID         *int    `json:"parent_id,omitempty"`
	Owner            *string `json:"owner,omitempty"`
	ModificationDate *int    `json:"modification_date,omitempty"`
	Assignable       *bool   `json:"assignable,omitempty"`
	ChildGroups      []int   `json:"child_groups,omitempty"`
}

// Get returns Log Source Groups of the current QRadar installation.
func (c *LogSourceGroupService) Get(ctx context.Context, fields, filter string, from, to int) ([]LogSourceGroup, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceTypeAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []LogSourceGroup
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns Log Source Group of the current QRadar installation by ID.
func (c *LogSourceGroupService) GetByID(ctx context.Context, fields string, id int) (*LogSourceGroup, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceTypeAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result LogSourceGroup
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
