package qradar

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

// LogSourceExtensionService handles methods related to Log Source Extensions of the QRadar Undocumented API.
type LogSourceExtensionService service

const (
	logSourceExtensionAPIPrefix = "api/config/event_sources/log_source_extensions"
)

// LogSourceExtension represents QRadar's Log Source Extension.
type LogSourceExtension struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Enabled      *bool   `json:"enabled,omitempty"`
	UseCondition *int    `json:"use_condition,omitempty"`
	XML          *string `json:"xml,omitempty"`
}

// Get returns Log Source Extension of the current QRadar installation. Undocumented API.
func (c *LogSourceExtensionService) Get(ctx context.Context, fields, filter string, from, to int) ([]LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceExtensionAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// todo: add support of multipart uploading

// Create creates Log Source Extension in the current QRadar installation. Undocumented API.
func (c *LogSourceExtensionService) Create(ctx context.Context, fields string, data interface{}) (*LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodPost, logSourceExtensionAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	defer w.Close()

	req.Header.Set("Allow-Hidden", "true")
	// req.Header.Set("Content-Type", w.FormDataContentType()) //
	req.Header.Set("Content-Type", "multipart/form-data")

	var result LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns Log Source Extension of the current QRadar installation by ID. Undocumented API.
func (c *LogSourceExtensionService) GetByID(ctx context.Context, fields string, id int) (*LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceExtensionAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// todo: add support of multipart uploading
/*
	// UpdateByID updates Log Source Extension of the current QRadar installation by ID. Undocumented API.
	func (c *LogSourceExtensionService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*LogSourceExtension, error) {
		req, err := c.client.requestHelp(http.MethodPost, logSourceExtensionAPIPrefix, fields, "", 0, 0, &id, data)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Allow-Hidden", "true")

		var result LogSourceExtension
		_, err = c.client.Do(ctx, req, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
*/

// GetByName returns Log Source Extension of the current QRadar installation by Name. Undocumented API.
func (c *LogSourceExtensionService) GetByName(ctx context.Context, fields string, name string) (*LogSourceExtension, error) {
	req, err := c.client.requestHelp(http.MethodGet, logSourceExtensionAPIPrefix, fields, fmt.Sprintf("name=\"%s\"", name), 0, 0, nil, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Allow-Hidden", "true")

	var result []LogSourceExtension
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	} else if len(result) > 1 {
		return nil, fmt.Errorf("found more log_source_extensions than expected - %d", len(result))
	}
	return &result[0], nil
}

func (c *Client) customRequest(method, urlStr, fields, filter string, id *int, body interface{}) (*http.Request, error) {
	if id != nil {
		urlStr = fmt.Sprintf("%s/%d", urlStr, *id)
	}

	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	//io.Closer
	w.CreateFormFile("file", b.String())
	w.Close()
	//var buf io.ReadWriter
	if body != nil {

	}

	req, err := http.NewRequest(method, u.String(), &b)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Accept", "application/json")

	if c.APIv == "" {
		req.Header.Set("Version", defaultAPIVersion)
	} else {
		req.Header.Set("Version", c.APIv)
	}

	//if buf != nil {
	//	req.Header.Set("Content-Type", "application/json")
	//}

	if c.SECKey != "" {
		req.Header.Set("SEC", c.SECKey)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	q := req.URL.Query()
	if fields != "" {
		q.Add("fields", fields)
	}
	if filter != "" {
		q.Add("filter", filter)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil

}
