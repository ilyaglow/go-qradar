// Package qradar provides an API client for the QRadar API.
// See examples of the usage in the examples folder.
package qradar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.1.0"
	apiVersion     = "10.0"
	userAgent      = "go-qradar/" + libraryVersion
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

	common service

	Ariel *ArielService
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
	}
	c.common.client = c
	c.Ariel = (*ArielService)(&c.common)

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
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Version", apiVersion)

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
