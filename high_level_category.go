package qradar

import (
	"context"
	"net/http"
)

// HighLevelCategoryService handles methods related to High Level Categories of the QRadar API.
type HighLevelCategoryService service

const (
	highLevelCategoryAPIPrefix = "api/data_classification/high_level_categories"
)

// HighLevelCategory represents QRadar's HighLevelCategory.
type HighLevelCategory struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Get returns HighLevelCategories of the current QRadar installation.
func (c *HighLevelCategoryService) Get(ctx context.Context, fields, filter string, from, to int) ([]HighLevelCategory, error) {
	req, err := c.client.requestHelp(http.MethodGet, highLevelCategoryAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []HighLevelCategory
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns HighLevelCategory of the current QRadar installation by ID.
func (c *HighLevelCategoryService) GetByID(ctx context.Context, fields string, id int) (*HighLevelCategory, error) {
	req, err := c.client.requestHelp(http.MethodGet, highLevelCategoryAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result HighLevelCategory
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
