package qradar

import "context"

// TenantService handles methods related to Tenants of the QRadar API.
type TenantService service

const (
	tenantAPIPrefix = "api/config/access/tenant_management/tenants"
)

// Tenant represents QRadar's Tenant.
type Tenant struct {
	ID             *int    `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Deleted        *bool   `json:"deleted,omitempty"`
	FlowRateLimit  *int    `json:"flow_rate_limit,omitempty"`
	EventRateLimit *int    `json:"event_rate_limit,omitempty"`
	Description    *string `json:"description,omitempty"`
}

// Get returns Tenants of the current QRadar installation.
func (c *TenantService) Get(ctx context.Context, fields, filter string, from, to int) ([]Tenant, error) {
	req, err := c.client.requestHelp("GET", tenantAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []Tenant
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates Tenant in QRadar installation.
func (c *TenantService) Create(ctx context.Context, fields string, data interface{}) (*Tenant, error) {
	req, err := c.client.requestHelp("POST", tenantAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result Tenant
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns Tenant of the current QRadar installation by ID.
func (c *TenantService) GetByID(ctx context.Context, fields string, id int) (*Tenant, error) {
	req, err := c.client.requestHelp("GET", tenantAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result Tenant
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates Tenant record in QRadar installation by ID.
func (c *TenantService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*Tenant, error) {
	req, err := c.client.requestHelp("POST", tenantAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result Tenant
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID deletes Tenant in QRadar installation by ID.
func (c *TenantService) DeleteByID(ctx context.Context, fields string, id int) (*Tenant, error) {
	req, err := c.client.requestHelp("DELETE", tenantAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result Tenant
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
