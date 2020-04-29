package qradar

import (
	"context"
	"net/http"
)

// NetworkHierarchyService handles methods related to Networkhierarchy of the QRadar API.
type NetworkHierarchyService service

const networkHierarchyAPIPrefix = "api/config/network_hierarchy/networks"

// NetworkHierarchy represents QRadar's generated NetworkHierarchy.
type NetworkHierarchy struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Cidr        *string `json:"cidr,omitempty"`
	ID          *int    `json:"id,omitempty"`
	DomainID    *int    `json:"domain_id,omitempty"`
	Group       *string `json:"group,omitempty"`
}

// Get returns Network Hierarchy of the current QRadar installation.
func (c *NetworkHierarchyService) Get(ctx context.Context, fields string) ([]NetworkHierarchy, error) {
	req, err := c.client.requestHelp(http.MethodGet, networkHierarchyAPIPrefix, fields, "", 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []NetworkHierarchy
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
