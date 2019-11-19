package qradar

import (
	"context"
	"net/http"
)

// DomainService handles methods related to Domains of the QRadar API.
type DomainService service

const domainsAPIPrefix = "api/config/domain_management/domains"

// Domain represents QRadar's Domain.
type Domain struct {
	AssetScannerIds  []int `json:"asset_scanner_ids,omitempty"`
	CustomProperties []struct {
		CaptureResult *string `json:"capture_result,omitempty"`
		ID            *int    `json:"id,omitempty"`
	} `json:"custom_properties,omitempty"`
	Deleted           *bool   `json:"deleted,omitempty"`
	Description       *string `json:"description,omitempty"`
	EventCollectorIds []int   `json:"event_collector_ids,omitempty"`
	FlowCollectorIds  []int   `json:"flow_collector_ids,omitempty"`
	FlowSourceIds     []int   `json:"flow_source_ids,omitempty"`
	FlowVlanIds       []int   `json:"flow_vlan_ids,omitempty"`
	ID                *int    `json:"id,omitempty"`
	LogSourceGroupIds []int   `json:"log_source_group_ids,omitempty"`
	LogSourceIds      []int   `json:"log_source_ids,omitempty"`
	Name              *string `json:"name,omitempty"`
	QvmScannerIds     []int   `json:"qvm_scanner_ids,omitempty"`
	TenantID          *int    `json:"tenant_id,omitempty"`
}

// Get returns Domains of the current QRadar installation
func (c *DomainService) Get(ctx context.Context, fields, filter string, from, to int) ([]Domain, error) {
	req, err := c.client.requestHelp(http.MethodGet, domainsAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []Domain
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns Domain of the current QRadar installation by ID.
func (c *DomainService) GetByID(ctx context.Context, fields string, id int) (*Domain, error) {
	req, err := c.client.requestHelp(http.MethodGet, domainsAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result Domain
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates Domain in the current QRadar installation.
func (c *DomainService) Create(ctx context.Context, fields string, data interface{}) (*Domain, error) {
	req, err := c.client.requestHelp(http.MethodPost, domainsAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result Domain
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates Domain in QRadar installation by ID.
func (c *DomainService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*Domain, error) {
	req, err := c.client.requestHelp(http.MethodPost, domainsAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result Domain
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID deletes Domain in QRadar installation by ID.
func (c *DomainService) DeleteByID(ctx context.Context, fields string, id int) (*Domain, error) {
	req, err := c.client.requestHelp(http.MethodDelete, domainsAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result Domain
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
