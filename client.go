package goodreads

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client is a public interface for client
type Client interface {
	Search(ctx context.Context, searchQuery string, page int) ([]Work, error)
	GetAllSeriesForWork(ctx context.Context, workID int) ([]Series, error)
	GetOneSeries(ctx context.Context, serieID int, page int) (SeriesWithWorks, error)
}

// client is holding everything to interact with goodreads API
type client struct {
	APIKey string
	format string
	domain string
	http   *http.Client
}

func (c *client) Get(ctx context.Context, endpoint string, query url.Values, response interface{}) error {
	query.Set("key", c.APIKey)
	query.Set("format", c.format)

	u, err := url.Parse(fmt.Sprintf("%s/%s", c.domain, endpoint))

	if err != nil {
		return fmt.Errorf("could not build search url: %w", err)
	}

	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)

	if err != nil {
		return fmt.Errorf("failed to build request for '%s': %w", u.Path, err)
	}

	resp, err := c.http.Do(req)

	if err != nil {
		return fmt.Errorf("request failed for '%s': %w", u.Path, err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("request failed for '%s': %w", u.Path, errors.New(resp.Status))
	}

	err = xml.NewDecoder(resp.Body).Decode(response)

	if err != nil {
		return fmt.Errorf("failed to decode response for '%s': %w", u.Path, err)
	}

	return nil
}

// NewClient creates a new goodreads api client with the given api key
func NewClient(apikey string) Client {
	return client{
		APIKey: apikey,
		format: "xml",
		domain: "https://www.goodreads.com",
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
