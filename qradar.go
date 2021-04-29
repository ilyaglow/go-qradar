// Package qradar provides an API client for the QRadar API.
// See examples of the usage in the examples folder.
package qradar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion    = "1.3.2"
	defaultAPIVersion = "12.0"
	userAgent         = "go-qradar/" + libraryVersion

	// ErrUnauthorized assigned on 401 http error.
	ErrUnauthorized = "unathorized"
)

// JobStatus represents status of the job: search, etc.
type JobStatus string

const (
	// StatusWait wait
	StatusWait JobStatus = "WAIT"

	// StatusExecute executing
	StatusExecute JobStatus = "EXECUTE"

	// StatusSorting sorting
	StatusSorting JobStatus = "SORTING"

	// StatusCompleted completed
	StatusCompleted JobStatus = "COMPLETED"

	// StatusCanceled canceled
	StatusCanceled JobStatus = "CANCELED"

	// StatusError errored
	StatusError JobStatus = "ERROR"
)

// Client manages communication with the QRadar API.
type Client struct {
	Client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	SECKey    string
	APIv      string

	common service

	Ariel                 *ArielService
	BuildingBlock         *BuildingBlockService
	BuildingBlockWithData *BuildingBlockWithDataService
	EventCollector        *EventCollectorService
	Offense               *OffenseService
	OffenseType           *OffenseTypeService
	Domain                *DomainService
	DSM                   *DSMService
	QID                   *QIDService
	LowLevelCategory      *LowLevelCategoryService
	HighLevelCategory     *HighLevelCategoryService
	RegexProperty         *RegexPropertyService
	Tenant                *TenantService
	Rule                  *RuleService
	RuleWithData          *RuleWithDataService
	RuleGroup             *RuleGroupService
	NetworkHierarchy      *NetworkHierarchyService

	PropertyExpression            *PropertyExpressionService
	PropertyJSONExpression        *PropertyJSONExpressionService
	PropertyLEEFExpression        *PropertyLEEFExpressionService
	PropertyCEFExpression         *PropertyCEFExpressionService
	ProperetyNVPExpression        *PropertyNVPExpressionService
	PropertyGenericListExpression *PropertyGenericListExpressionService
	PropertyXMLExpression         *PropertyXMLExpressionService

	LogSourceExtension *LogSourceExtensionService
	LogSourceType      *LogSourceTypeService
	LogSourceGroup     *LogSourceGroupService
	LogSource          *LogSourceService

	ReferenceMapOfSets *ReferenceMapOfSetsService
	ReferenceMap       *ReferenceMapService
	ReferenceSet       *ReferenceSetService
	ReferenceTable     *ReferenceTableService
}

type service struct {
	client *Client
}

// NewClient returns a new QRadar API client.
func NewClient(baseurl string, opts ...func(*Client) error) (*Client, error) {
	u, err := url.Parse(baseurl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client:    http.DefaultClient,
		UserAgent: userAgent,
		BaseURL:   u,
		APIv:      defaultAPIVersion,
	}
	c.common.client = c
	c.Ariel = (*ArielService)(&c.common)
	c.BuildingBlock = (*BuildingBlockService)(&c.common)
	c.BuildingBlockWithData = (*BuildingBlockWithDataService)(&c.common)
	c.EventCollector = (*EventCollectorService)(&c.common)
	c.Offense = (*OffenseService)(&c.common)
	c.OffenseType = (*OffenseTypeService)(&c.common)
	c.Domain = (*DomainService)(&c.common)
	c.DSM = (*DSMService)(&c.common)
	c.QID = (*QIDService)(&c.common)
	c.RegexProperty = (*RegexPropertyService)(&c.common)
	c.Rule = (*RuleService)(&c.common)
	c.RuleWithData = (*RuleWithDataService)(&c.common)
	c.RuleGroup = (*RuleGroupService)(&c.common)
	c.PropertyExpression = (*PropertyExpressionService)(&c.common)
	c.PropertyJSONExpression = (*PropertyJSONExpressionService)(&c.common)
	c.PropertyGenericListExpression = (*PropertyGenericListExpressionService)(&c.common)
	c.PropertyLEEFExpression = (*PropertyLEEFExpressionService)(&c.common)
	c.PropertyCEFExpression = (*PropertyCEFExpressionService)(&c.common)
	c.ProperetyNVPExpression = (*PropertyNVPExpressionService)(&c.common)
	c.PropertyXMLExpression = (*PropertyXMLExpressionService)(&c.common)
	c.LogSourceExtension = (*LogSourceExtensionService)(&c.common)
	c.LogSourceType = (*LogSourceTypeService)(&c.common)
	c.LogSourceGroup = (*LogSourceGroupService)(&c.common)
	c.LogSource = (*LogSourceService)(&c.common)
	c.LowLevelCategory = (*LowLevelCategoryService)(&c.common)
	c.HighLevelCategory = (*HighLevelCategoryService)(&c.common)
	c.Tenant = (*TenantService)(&c.common)
	c.ReferenceMapOfSets = (*ReferenceMapOfSetsService)(&c.common)
	c.ReferenceMap = (*ReferenceMapService)(&c.common)
	c.ReferenceSet = (*ReferenceSetService)(&c.common)
	c.ReferenceTable = (*ReferenceTableService)(&c.common)
	c.NetworkHierarchy = (*NetworkHierarchyService)(&c.common)

	for _, f := range opts {
		err := f(c)
		if err != nil {
			return c, err
		}
	}

	return c, nil
}

// SetHTTPClient sets an HTTP client.
func SetHTTPClient(httpClient *http.Client) func(*Client) error {
	return func(c *Client) error {
		c.Client = httpClient
		return nil
	}
}

// SetSECKey sets a key to auth on the QRadar API
func SetSECKey(key string) func(*Client) error {
	return func(c *Client) error {
		c.SECKey = key
		return nil
	}
}

// SetAPIversion sets a version of QRadar API
func SetAPIversion(api string) func(*Client) error {
	return func(c *Client) error {
		c.APIv = api
		return nil
	}
}

func (c *Client) requestHelp(method, urlStr, fields, filter string, from, to int, id *int, body interface{}) (*http.Request, error) {
	if id != nil {
		urlStr = fmt.Sprintf("%s/%d", urlStr, *id)
	}
	req, err := c.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	if from == 0 && to != 0 {
		req.Header.Add("Range", fmt.Sprintf("items=%d-%d", from, to))
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

// NewRequest constructs and new request to send.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}

		// QRadar reference_data API errors if json has a trailing new line
		if strings.HasPrefix(urlStr, "api/reference_data") && strings.Contains(urlStr, "/bulk_load/") {
			bs, err := ioutil.ReadAll(buf)
			if err != nil {
				return nil, err
			}
			buf = bytes.NewBuffer(bs[:len(bs)-1])
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if c.APIv == "" {
		req.Header.Set("Version", defaultAPIVersion)
	} else {
		req.Header.Set("Version", c.APIv)
	}

	if buf != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.SECKey != "" {
		req.Header.Set("SEC", c.SECKey)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.Client.Do(req)

	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return resp, err
}

// CheckResponse checks the API response for errors.
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		return nil
	case http.StatusUnauthorized:
		return fmt.Errorf("%s %d: %s", r.Request.URL.Path, r.StatusCode, ErrUnauthorized)
	case http.StatusNotFound, http.StatusConflict, http.StatusUnprocessableEntity, http.StatusBadRequest,
		http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusForbidden:
		var v ErrorMessage
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			return fmt.Errorf("%s %d: %s", r.Request.URL.Path, r.StatusCode, err.Error())
		}
		v.resp = r
		return &v
	default:
		return fmt.Errorf("%s %d: unknown error", r.Request.URL.Path, r.StatusCode)
	}
}

// ErrorMessage represents generic error message by the QRadar API.
type ErrorMessage struct {
	resp        *http.Response
	Code        json.Number `json:"code,omitempty"`
	Contexts    []string    `json:"contexts,omitempty"`
	Message     string      `json:"message,omitempty"`
	Description string      `json:"description,omitempty"`
	Severity    string      `json:"severity,omitempty"`
	Details     struct {
		Reason      string `json:"reason,omitempty"`
		Code        int    `json:"code,omitempty"`
		StartIndex  int    `json:"start_index,omitempty"`
		LineNumber  int    `json:"line_number,omitempty"`
		QueryString string `json:"query_string,omitempty"`
		TokenText   string `json:"token_text,omitempty"`
	} `json:"details,omitempty"`
}

// Error satisfies the error interface.
func (e *ErrorMessage) Error() string {
	return fmt.Sprintf(
		"%s %d: %s [%s]",
		e.resp.Request.URL.Path, e.resp.StatusCode, e.Message, string(e.Code),
	)
}
