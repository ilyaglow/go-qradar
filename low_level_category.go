package qradar

import (
	"context"
	"net/http"
)

// LowLevelCategoryService handles methods related to Low Level Categories of the QRadar API.
type LowLevelCategoryService service

const (
	lowLevelCategoryAPIPrefix = "api/data_classification/low_level_categories"
)

// LowLevelCategory represents QRadar's LowLevelCategory.
type LowLevelCategory struct {
	ID                  *int    `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	Description         *string `json:"description,omitempty"`
	Severity            *int    `json:"severity,omitempty"`
	HighLevelCategoryID *int    `json:"high_level_category_id,omitempty"`
}

// Get returns LowLevelCategories of the current QRadar installation.
func (c *LowLevelCategoryService) Get(ctx context.Context, fields, filter string, from, to int) ([]LowLevelCategory, error) {
	req, err := c.client.requestHelp(http.MethodGet, lowLevelCategoryAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []LowLevelCategory
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns LowLevelCategory of the current QRadar installation by ID.
func (c *LowLevelCategoryService) GetByID(ctx context.Context, fields string, id int) (*LowLevelCategory, error) {
	req, err := c.client.requestHelp(http.MethodGet, lowLevelCategoryAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result LowLevelCategory
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
