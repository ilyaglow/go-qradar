package qradar

import (
	"context"
	"net/http"
)

// EventCollectorService handles methods related to Event Collector of the QRadar API.
type EventCollectorService service

const (
	eventCollectorAPIPrefix = "api/config/event_sources/event_collectors"
)

// EventCollector represents QRadar's Event Collector
type EventCollector struct {
	ID            *int    `json:"id,omitempty"`
	ComponentName *string `json:"component_name,omitempty"`
	Name          *string `json:"name,omitempty"`
	HostID        *int    `json:"host_id,omitempty"`
}

// Get returns DSMs of the current QRadar installation.
func (c *EventCollectorService) Get(ctx context.Context, fields, filter string, from, to int) ([]EventCollector, error) {
	req, err := c.client.requestHelp(http.MethodGet, eventCollectorAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []EventCollector
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
