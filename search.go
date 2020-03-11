package goodreads

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type searchResponse struct {
	Results []Work `xml:"search>results>work"`
}

// Search find any book by title, author or isbn
// for pagination 0 or 1 seems to be the same thing
func (c client) Search(ctx context.Context, searchQuery string, page int) ([]Work, error) {
	var response = searchResponse{
		Results: []Work{},
	}

	q := url.Values{}
	q.Set("q", searchQuery)
	q.Set("page", strconv.Itoa(page))

	err := c.Get(ctx, "/search/index", q, &response)

	if err != nil {
		return []Work{}, fmt.Errorf("'%s' search at page %d failed: %w", searchQuery, page, err)
	}

	return response.Results, nil
}
