package qradar

import (
	"context"
	"net/http"
)

// LogSourceService handles methods related to Log Sources of the QRadar API.
type LogSourceService service

const (
	logSourceAPIPrefix = "api/config/event_sources/log_source_management/log_sources"
)

// LogSource represents QRadar's Log Source Type.
type LogSource struct {
	SendingIP           *string `json:"sending_ip,omitempty"`
	Internal            *bool   `json:"internal,omitempty"`
	LegacyBulkGroupName *string `json:"legacy_bulk_group_name,omitempty"`
	ProtocolParameters  []struct {
		Name  *string `json:"name,omitempty"`
		ID    *int    `json:"id,omitempty"`
		Value *string `json:"value,omitempty"`
	} `json:"protocol_parameters,omitempty"`
	Description                      *string `json:"description,omitempty"`
	CoalesceEvents                   *bool   `json:"coalesce_events,omitempty"`
	Enabled                          *bool   `json:"enabled,omitempty"`
	GroupIDs                         []int   `json:"group_ids,omitempty"`
	AverageEps                       *int    `json:"average_eps,omitempty"`
	Credibility                      *int    `json:"credibility,omitempty"`
	ID                               *int    `json:"id,omitempty"`
	StoreEventPayload                *bool   `json:"store_event_payload,omitempty"`
	TargetEventCollectorID           *int    `json:"target_event_collector_id,omitempty"`
	ProtocolTypeID                   *int    `json:"protocol_type_id,omitempty"`
	LanguageID                       *int    `json:"language_id,omitempty"`
	CreationDate                     *int    `json:"creation_date,omitempty"`
	LogSourceExtensionID             *int    `json:"log_source_extension_id,omitempty"`
	WincollectExternalDestinationIDs []int   `json:"wincollect_external_destination_ids,omitempty"`
	Name                             *string `json:"name,omitempty"`
	AutoDiscovered                   *bool   `json:"auto_discovered,omitempty"`
	ModifiedDate                     *int    `json:"modified_date,omitempty"`
	TypeID                           *int    `json:"type_id,omitempty"`
	LastEventTime                    *int    `json:"last_event_time,omitempty"`
	RequiresDeploy                   *bool   `json:"requires_deploy,omitempty"`
	Gateway                          *bool   `json:"gateway,omitempty"`
	WincollectInternalDestinationID  *int    `json:"wincollect_internal_destination_id,omitempty"`
	Status                           struct {
		LastUpdated *int `json:"last_updated,omitempty"`
		Messages    []struct {
			Severity  *string `json:"severity,omitempty"`
			Text      *string `json:"text,omitempty"`
			Timestamp *int    `json:"timestamp,omitempty"`
		} `json:"messages,omitempty"`
		Status *string `json:"status,omitempty"`
	} `json:"status,omitempty"`
}

// Get returns Log Sources of the current QRadar installation.
func (c *LogSourceService) Get(ctx context.Context, fields, filter string, from, to int) ([]LogSource, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []LogSource
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
