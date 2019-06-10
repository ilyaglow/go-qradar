package qradar

import (
	"context"
	"fmt"
)

const domainsAPIPrefix = configAPIPrefix + "/domain_management/domains"

// Domain represents QRadar domains.
type Domain struct {
	AssetScannerIds  []int `json:"asset_scanner_ids"`
	CustomProperties []struct {
		CaptureResult string `json:"capture_result"`
		ID            int    `json:"id"`
	} `json:"custom_properties"`
	Deleted           bool   `json:"deleted"`
	Description       string `json:"description"`
	EventCollectorIds []int  `json:"event_collector_ids"`
	FlowCollectorIds  []int  `json:"flow_collector_ids"`
	FlowSourceIds     []int  `json:"flow_source_ids"`
	FlowVlanIds       []int  `json:"flow_vlan_ids"`
	ID                int    `json:"id"`
	LogSourceGroupIds []int  `json:"log_source_group_ids"`
	LogSourceIds      []int  `json:"log_source_ids"`
	Name              string `json:"name"`
	QvmScannerIds     []int  `json:"qvm_scanner_ids"`
	TenantID          int    `json:"tenant_id"`
}

// Domains of the current QRadar installation.
func (c *ConfigService) Domains(ctx context.Context, fields, filter string, from, to int) ([]Domain, error) {
	req, err := c.client.NewRequest("GET", domainsAPIPrefix, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Range", fmt.Sprintf("items=%d-%d", from, to))

	q := req.URL.Query()
	if fields != "" {
		q.Add("fields", fields)
	}
	if filter != "" {
		q.Add("filter", filter)
	}
	req.URL.RawQuery = q.Encode()

	var doms []Domain
	_, err = c.client.Do(ctx, req, &doms)
	if err != nil {
		return nil, err
	}

	return doms, nil
}
