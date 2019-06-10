package qradar

import (
	"context"
	"fmt"
)

const offensesAPIPrefix = "api/siem/offenses"

// SIEMService handles communication with siem path-related methods of the
// QRadar API.
type SIEMService service

// Offense represents QRadar's generated offense.
type Offense struct {
	UsernameCount int    `json:"username_count,omitempty"`
	Description   string `json:"description,omitempty"`
	Rules         []struct {
		ID   int    `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"rules,omitempty"`
	EventCount                 int      `json:"event_count,omitempty"`
	FlowCount                  int      `json:"flow_count,omitempty"`
	AssignedTo                 string   `json:"assigned_to,omitempty"`
	SecurityCategoryCount      int      `json:"security_category_count,omitempty"`
	FollowUp                   bool     `json:"follow_up,omitempty"`
	SourceAddressIds           []int    `json:"source_address_ids,omitempty"`
	SourceCount                int      `json:"source_count,omitempty"`
	Inactive                   bool     `json:"inactive,omitempty"`
	Protected                  bool     `json:"protected,omitempty"`
	CategoryCount              int      `json:"category_count,omitempty"`
	SourceNetwork              string   `json:"source_network,omitempty"`
	DestinationNetworks        []string `json:"destination_networks,omitempty"`
	ClosingUser                string   `json:"closing_user,omitempty"`
	CloseTime                  int      `json:"close_time,omitempty"`
	RemoteDestinationCount     int      `json:"remote_destination_count,omitempty"`
	StartTime                  int      `json:"start_time,omitempty"`
	LastUpdatedTime            int      `json:"last_updated_time,omitempty"`
	Credibility                int      `json:"credibility,omitempty"`
	Magnitude                  int      `json:"magnitude,omitempty"`
	ID                         int      `json:"id,omitempty"`
	Categories                 []string `json:"categories,omitempty"`
	Severity                   int      `json:"severity,omitempty"`
	PolicyCategoryCount        int      `json:"policy_category_count,omitempty"`
	DeviceCount                int      `json:"device_count,omitempty"`
	ClosingReasonID            int      `json:"closing_reason_id,omitempty"`
	OffenseType                int      `json:"offense_type,omitempty"`
	Relevance                  int      `json:"relevance,omitempty"`
	DomainID                   int      `json:"domain_id,omitempty"`
	OffenseSource              string   `json:"offense_source,omitempty"`
	LocalDestinationAddressIds []int    `json:"local_destination_address_ids,omitempty"`
	LocalDestinationCount      int      `json:"local_destination_count,omitempty"`
	Status                     string   `json:"status,omitempty"`
}

// Offenses of the QRadar that could be filtered or paged.
func (s *SIEMService) Offenses(ctx context.Context, fields string, filter string, from int, to int) ([]Offense, error) {
	req, err := s.client.NewRequest("GET", offensesAPIPrefix, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Range", fmt.Sprintf("items=%d-%d", from, to))

	q := req.URL.Query()
	q.Add("fields", fields)
	q.Add("filter", filter)
	req.URL.RawQuery = q.Encode()

	var ofs []Offense
	_, err = s.client.Do(ctx, req, &ofs)
	if err != nil {
		return nil, err
	}

	return ofs, nil
}
